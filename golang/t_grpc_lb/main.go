package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"github.com/qingsong-he/ce"
	"github.com/qingsong-he/some/golang/t_grpc_lb/consistentlb"
	"github.com/qingsong-he/some/golang/t_grpc_lb/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"runtime/debug"
	"strconv"
	"syscall"
	"time"
)

func init() {
	ce.Print(os.Args[0])
}

type getName interface {
	GetName() string
}

type msg struct{}

func (*msg) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: serverAddr}, nil
}

var serverName = "srv1"
var serverAddr string
var serverCount int64 = 3

func regSelf(cli *clientv3.Client, serverName, serverAddr string, serverCount int64) (ok bool) {
	defer myRecover(func() {
		ok = false
	})

	m := make([]int64, serverCount)

	sess, err := concurrency.NewSession(cli)
	ce.CheckError(err)
	defer sess.Close()

	ctx := context.Background()

	elec := concurrency.NewElection(sess, "commonElectionPfx_"+serverName)
	err = elec.Campaign(ctx, serverAddr)
	ce.CheckError(err)
	defer elec.Resign(ctx)

	keyByServerName, err := cli.Get(ctx, serverName+"/", clientv3.WithPrefix())
	ce.CheckError(err)

	if keyByServerName.Count == 0 {
		_regSelf(cli, serverName, serverAddr, 0)
	} else {
		for _, v := range keyByServerName.Kvs {
			var up naming.Update
			err := json.Unmarshal(v.Value, &up)
			ce.CheckError(err)
			m[int64(up.Metadata.(float64))] = 1
		}
		var isFull bool = true
		for index, v := range m {
			if v == 0 {
				isFull = false
				_regSelf(cli, serverName, serverAddr, index)
				break
			}
		}
		if isFull {
			panic(fmt.Errorf("no more index resource: %d", serverCount))
		}
	}
	return true
}

func _regSelf(cli *clientv3.Client, serverName, serverAddr string, index int) {
	r := &etcdnaming.GRPCResolver{Client: cli}

	leaseId, err := r.Client.Grant(context.TODO(), 4)
	ce.CheckError(err)

	keepAliveChan, err := r.Client.KeepAlive(context.TODO(), leaseId.ID)
	ce.CheckError(err)

	err = r.Update(context.TODO(), serverName, naming.Update{Op: naming.Add, Addr: serverAddr, Metadata: index}, clientv3.WithLease(leaseId.ID))
	ce.CheckError(err)

	go func() {
		for {
			select {
			case _, ok := <-keepAliveChan:
				func() {
					defer myRecover()
					if !ok {
						ctx := context.Background()
						leaseId, err = r.Client.Grant(ctx, 4)
						ce.CheckError(err)
						keepAliveChan, err = r.Client.KeepAlive(ctx, leaseId.ID)
						ce.CheckError(err)
						err = r.Update(ctx, serverName, naming.Update{Op: naming.Add, Addr: serverAddr, Metadata: index}, clientv3.WithLease(leaseId.ID))
						ce.CheckError(err)
					}
				}()
			}
		}
	}()
}

var s = func(addr string) {
	defer myRecover()

	// random port
	lis, err := net.Listen("tcp", addr)
	ce.CheckError(err)

	serverAddr = lis.Addr().(*net.TCPAddr).IP.String() + ":" + strconv.Itoa(lis.Addr().(*net.TCPAddr).Port)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 2 * time.Second,
	})
	ce.CheckError(err)
	defer cli.Close()

	for {
		if !regSelf(cli, serverName, serverAddr, serverCount) {
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &msg{})

	err = grpcServer.Serve(lis)
	ce.CheckError(err)
}

var c = func() {
	defer myRecover()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 2 * time.Second,
	})
	ce.CheckError(err)
	defer cli.Close()

	r := &etcdnaming.GRPCResolver{Client: cli}
	conn, err := grpc.Dial(
		serverName,
		grpc.WithInsecure(),
		grpc.WithBalancer(consistentlb.ConsistentLB(r, serverCount)),
		grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			// get name by req
			var name string
			switch v := req.(type) {
			case getName:
				name = v.GetName()
			}
			err := invoker(context.WithValue(ctx, "name", name), method, req, reply, cc, opts...)
			return err
		}),
		grpc.WithTimeout(2*time.Second),
	)
	ce.CheckError(err)
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	rand.Seed(time.Now().UnixNano())

	for {
		func() {
			defer myRecover()
			name := strconv.Itoa(int(time.Now().UnixNano()))
			resp, err := c.SayHello(context.TODO(), &pb.HelloRequest{Name: name})
			ce.CheckError(err)
			ce.Print(name, "->", resp.Message)

		}()
		time.Sleep(1 * time.Second)
	}
}

var myRecover = func(deferList ...func()) {
	if errByRecov := recover(); errByRecov != nil {
		if _, ok := ce.IsFromCe(errByRecov); !ok {
			ce.Print("panic with:", errByRecov, string(debug.Stack()))
		}
		for _, v := range deferList {
			v()
		}
	}
}

var listenSignal = func() {
	mainByExitAlarm := make(chan os.Signal, 1)
	signal.Notify(mainByExitAlarm, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)

forLableByNotify:
	for {
		s := <-mainByExitAlarm
		ce.Print(s)
		switch s {
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
			break forLableByNotify

		case syscall.SIGHUP:
		default:
			break forLableByNotify
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	tp := os.Args[1]

	if tp == "s" {
		go s("localhost:0")
	} else if tp == "c" {
		go c()
	} else {
		return
	}

	listenSignal()
}
