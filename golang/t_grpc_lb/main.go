package main

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"github.com/qingsong-he/ce"
	"github.com/qingsong-he/some/golang/t_grpc_lb/consistentlb"
	"github.com/qingsong-he/some/golang/t_grpc_lb/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
	"net"
	"os"
	"os/signal"
	"runtime/debug"
	"strconv"
	"strings"
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
	return &pb.HelloReply{Message: serverAddrByGlobal}, nil
}

var serverName = "srv1"
var serverAddrByGlobal string

var s = func(serverAddr string) {
	serverAddrByGlobal = serverAddr
	defer myRecover()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 2 * time.Second,
	})
	ce.CheckError(err)
	defer cli.Close()

	r := &etcdnaming.GRPCResolver{Client: cli}

	leaseId, err := r.Client.Grant(context.TODO(), 4)
	ce.CheckError(err)
	keepAliveChan, err := r.Client.KeepAlive(context.TODO(), leaseId.ID)
	ce.CheckError(err)
	port, err := strconv.Atoi(strings.Split(serverAddr, ":")[1])
	ce.CheckError(err)
	err = r.Update(context.TODO(), serverName, naming.Update{Op: naming.Add, Addr: serverAddr, Metadata: port}, clientv3.WithLease(leaseId.ID))
	ce.CheckError(err)

	go func() {
		for {
			select {
			case _, ok := <-keepAliveChan:
				func() {
					defer myRecover()
					if !ok {
						leaseId, err = r.Client.Grant(r.Client.Ctx(), 4)
						ce.CheckError(err)
						keepAliveChan, err = r.Client.KeepAlive(r.Client.Ctx(), leaseId.ID)
						ce.CheckError(err)
						err = r.Update(context.TODO(), serverName, naming.Update{Op: naming.Add, Addr: serverAddr}, clientv3.WithLease(leaseId.ID))
						ce.CheckError(err)
					}
				}()
			}
		}
	}()

	lis, err := net.Listen("tcp", serverAddr)
	ce.CheckError(err)

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
		grpc.WithBalancer(consistentlb.ConsistentLB(r)),
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

	for {
		func() {
			defer myRecover()

			var name string
			if time.Now().UnixNano()%2 == 0 {
				name = "1"
			}
			resp, err := c.SayHello(context.TODO(), &pb.HelloRequest{Name: name})
			ce.CheckError(err)
			ce.Print(name, "->", resp.Message)

		}()
		time.Sleep(1 * time.Second)
	}
}

var myRecover = func() {
	if errByRecov := recover(); errByRecov != nil {
		if _, ok := ce.IsFromCe(errByRecov); !ok {
			ce.Print("panic with:", errByRecov, string(debug.Stack()))
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

	if tp == "s1" {
		go s("localhost:3001")
	} else if tp == "s2" {
		go s("localhost:3002")
	} else if tp == "c" {
		go c()
	} else {
		return
	}

	listenSignal()
}
