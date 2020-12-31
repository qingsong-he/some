package consistentlb

import (
	"context"
	"github.com/qingsong-he/ce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/naming"
	"google.golang.org/grpc/status"
	"io"
	"strconv"
	"sync"
)

var (
	noAddrErr = status.Errorf(codes.Unavailable, "there is no address available")
)

func ConsistentLB(r naming.Resolver, serverCount int64) grpc.Balancer {
	return &consistentLB{
		serverCount: serverCount,
		r:           r,
	}
}

type addrInfo struct {
	addr      grpc.Address
	connected bool
}

type consistentLB struct {
	serverCount int64
	r           naming.Resolver
	w           naming.Watcher
	addrs       []*addrInfo // all the addresses the client should potentially connect
	mu          sync.Mutex
	addrCh      chan []grpc.Address // the channel to notify gRPC internals the list of addresses the client should connect to.
	done        bool                // The Balancer is closed.
}

func (rr *consistentLB) watchAddrUpdates() error {
	updates, err := rr.w.Next()
	if err != nil {
		ce.Print(err.Error())
		return err
	}
	rr.mu.Lock()
	defer rr.mu.Unlock()
	for _, update := range updates {
		addr := grpc.Address{
			Addr:     update.Addr,
			Metadata: update.Metadata,
		}
		switch update.Op {
		case naming.Add:
			var exist bool
			for _, v := range rr.addrs {
				if addr == v.addr {
					exist = true
					ce.Print("grpc: The name resolver wanted to add an existing address: ", addr)
					break
				}
			}
			if exist {
				continue
			}
			rr.addrs = append(rr.addrs, &addrInfo{addr: addr})
		case naming.Delete:
			for i, v := range rr.addrs {
				if addr == v.addr {
					copy(rr.addrs[i:], rr.addrs[i+1:])
					rr.addrs = rr.addrs[:len(rr.addrs)-1]
					break
				}
			}
		default:
			ce.Print("Unknown update.Op ", update.Op)
		}
	}

	open := make([]grpc.Address, len(rr.addrs))
	for i, v := range rr.addrs {
		open[i] = v.addr
	}
	if rr.done {
		return grpc.ErrClientConnClosing
	}
	select {
	case <-rr.addrCh:
	default:
	}
	rr.addrCh <- open
	return nil
}

func (rr *consistentLB) Start(target string, config grpc.BalancerConfig) error {
	rr.mu.Lock()
	defer rr.mu.Unlock()
	if rr.done {
		return grpc.ErrClientConnClosing
	}
	if rr.r == nil {
		return io.EOF
	}
	w, err := rr.r.Resolve(target)
	if err != nil {
		return err
	}
	rr.w = w
	rr.addrCh = make(chan []grpc.Address, 1)
	go func() {
		for {
			if err := rr.watchAddrUpdates(); err != nil {
				return
			}
		}
	}()
	return nil
}

func (rr *consistentLB) Up(addr grpc.Address) func(error) {
	rr.mu.Lock()
	defer rr.mu.Unlock()
	var cnt int
	for _, a := range rr.addrs {
		if a.addr == addr {
			if a.connected {
				return nil
			}
			a.connected = true
		}
		if a.connected {
			cnt++
		}
	}
	return func(err error) {
		rr.down(addr, err)
	}
}

func (rr *consistentLB) down(addr grpc.Address, err error) {
	rr.mu.Lock()
	defer rr.mu.Unlock()
	for _, a := range rr.addrs {
		if addr == a.addr {
			a.connected = false
			break
		}
	}
}

func (rr *consistentLB) Get(ctx context.Context, opts grpc.BalancerGetOptions) (addr grpc.Address, put func(), err error) {
	rr.mu.Lock()
	defer rr.mu.Unlock()

	if rr.done {
		err = grpc.ErrClientConnClosing
		return
	}

	for _, v := range rr.addrs {
		if v.connected && v.addr.Metadata != nil {
			name := ctx.Value("name").(string)
			nameByInt, err1 := strconv.Atoi(name)
			if err1 != nil {
				err = err1
				return
			}
			if int64(nameByInt)%rr.serverCount == int64(v.addr.Metadata.(float64)) {
				addr = v.addr
				return
			}
		}

	}

	err = noAddrErr
	return
}

func (rr *consistentLB) Notify() <-chan []grpc.Address {
	return rr.addrCh
}

func (rr *consistentLB) Close() error {
	rr.mu.Lock()
	defer rr.mu.Unlock()
	if rr.done {
		return io.EOF
	}
	rr.done = true
	if rr.w != nil {
		rr.w.Close()
	}
	if rr.addrCh != nil {
		close(rr.addrCh)
	}
	return nil
}
