package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	. "github.com/qingsong-he/ce"
	"github.com/qingsong-he/some/golang/t_grpc/pb"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/naming"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	Print(os.Args[0])
}

func InitTracer(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: "localhost:6831",
			LogSpans:           true,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	CheckError(err)
	return tracer, closer
}

func client1(ctx context.Context) {
	span, ctxByThis := opentracing.StartSpanFromContext(ctx, "client1")
	defer span.Finish()
	time.Sleep(1 * time.Second)
	span.LogKV("c1", "c1")
	client1Sub1(ctxByThis)
}

func client1Sub1(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "client1Sub1")
	defer span.Finish()
	time.Sleep(1 * time.Second)
	span.LogKV("c1Sub1", "c1Sub1")
}

func server1(ctx context.Context) {
	span, ctxByThis := opentracing.StartSpanFromContext(ctx, "server1")
	defer span.Finish()
	time.Sleep(1 * time.Second)
	span.LogKV("s1", "s1")
	server1Sub1(ctxByThis)
}

func server1Sub1(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "server1Sub1")
	defer span.Finish()
	time.Sleep(1 * time.Second)
	span.LogKV("s1Sub1", "s1Sub1")
}

type msg struct{}

func (*msg) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	if p, ok := peer.FromContext(ctx); ok {
		Print(p.Addr.String())
	}

	server1(ctx)
	h := sha1.Sum([]byte(in.Name))
	return &pb.HelloReply{Message: hex.EncodeToString(h[:])}, nil
}

var serverName = "srv1"

var s = func(serverAddr string) {
	defer func() {
		recover()
	}()

	tracer, closer := InitTracer(serverName)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

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

	// Define customfunc to handle panic
	customFunc := func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(customFunc),
	}

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(opts...),
			grpc_opentracing.StreamServerInterceptor(grpc_opentracing.WithTracer(tracer)),
		),
	)
	pb.RegisterGreeterServer(grpcServer, &msg{})

	err = grpcServer.Serve(lis)
	CheckError(err)
}

var c = func() {
	defer func() {
		recover()
	}()

	tracer, closer := InitTracer(serverName + "ByClient")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 2 * time.Second,
	})
	CheckError(err)
	defer cli.Close()

	r := &etcdnaming.GRPCResolver{Client: cli}
	conn, err := grpc.Dial(
		serverName,
		grpc.WithInsecure(),
		grpc.WithBalancer(grpc.RoundRobin(r)),
		grpc.WithTimeout(2*time.Second),
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(tracer),
			),
		),
	)
	CheckError(err)
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	for {
		func() {
			defer func() {
				recover()
			}()

			span := tracer.StartSpan("callSayHello")
			span.LogKV("root", "root")
			span.SetTag("rootTag", "rootTag")
			defer span.Finish()
			ctx := opentracing.ContextWithSpan(context.Background(), span)

			client1(ctx)

			resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: time.Now().String()})
			CheckError(err)
			Print(resp.Message)

		}()
		time.Sleep(30 * time.Second)
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
