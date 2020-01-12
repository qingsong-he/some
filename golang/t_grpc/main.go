package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	. "github.com/qingsong-he/ce"
	"github.com/qingsong-he/testcode/golang/t_grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
	"google.golang.org/grpc/peer"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	Print(os.Args[0])
}

type msg struct{}

func (*msg) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	if p, ok := peer.FromContext(ctx); ok {
		Print(p.Addr.String())
	}

	h := sha1.Sum([]byte(in.Name))
	return &pb.HelloReply{Message: hex.EncodeToString(h[:])}, nil
}

var serverName = "srv1"

var s = func(serverAddr string) {
	defer func() {
		recover()
	}()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 2 * time.Second,
	})
	CheckError(err)
	defer cli.Close()

	r := &etcdnaming.GRPCResolver{Client: cli}

	leaseId, err := r.Client.Grant(context.TODO(), 4)
	CheckError(err)
	keepAliveChan, err := r.Client.KeepAlive(context.TODO(), leaseId.ID)
	CheckError(err)
	err = r.Update(context.TODO(), serverName, naming.Update{Op: naming.Add, Addr: serverAddr}, clientv3.WithLease(leaseId.ID))
	CheckError(err)

	go func() {
		for {
			select {
			case _, ok := <-keepAliveChan:
				func() {
					defer func() {
						recover()
					}()
					if !ok {
						leaseId, err = r.Client.Grant(r.Client.Ctx(), 4)
						CheckError(err)
						keepAliveChan, err = r.Client.KeepAlive(r.Client.Ctx(), leaseId.ID)
						CheckError(err)
						err = r.Update(context.TODO(), serverName, naming.Update{Op: naming.Add, Addr: serverAddr}, clientv3.WithLease(leaseId.ID))
						CheckError(err)
					}
				}()
			}
		}
	}()

	lis, err := net.Listen("tcp", serverAddr)
	CheckError(err)

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &msg{})

	err = grpcServer.Serve(lis)
	CheckError(err)
}

var c = func() {
	defer func() {
		recover()
	}()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 2 * time.Second,
	})
	CheckError(err)
	defer cli.Close()

	r := &etcdnaming.GRPCResolver{Client: cli}
	conn, err := grpc.Dial(serverName, grpc.WithInsecure(), grpc.WithBalancer(grpc.RoundRobin(r)), grpc.WithTimeout(2*time.Second))
	CheckError(err)
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	for {
		func() {
			defer func() {
				recover()
			}()

			time.Sleep(1 * time.Second)
			resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: time.Now().String()})
			CheckError(err)
			Print(resp.Message)
		}()
	}
}

var listenSignal = func() {
	mainByExitAlarm := make(chan os.Signal, 1)
	signal.Notify(mainByExitAlarm, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)

forLableByNotify:
	for {
		s := <-mainByExitAlarm
		Print(s)
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
