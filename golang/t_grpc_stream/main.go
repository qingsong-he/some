package main

import (
	"context"
	. "github.com/qingsong-he/ce"
	"github.com/qingsong-he/some/golang/t_grpc_stream/pb"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	Print(os.Args[0])
}

type msg struct{}

// client stream
func (*msg) SayHello(s pb.Greeter_SayHelloServer) error {
	var names []string
	for {
		in, err := s.Recv()
		if err == io.EOF {
			s.SendAndClose(&pb.HelloReply{Message: "Hello " + strings.Join(names, ",")})
			return nil
		}
		if err != nil {
			Printf("failed to recv: %v", err)
			return err
		}
		names = append(names, in.Name)
	}
	return nil
}

// server stream
func (*msg) SayHello1(in *pb.HelloRequest, s pb.Greeter_SayHello1Server) error {
	for i := 0; i < 3; i++ {
		err := s.Send(&pb.HelloReply{Message: in.Name + " " + strconv.Itoa(i)})
		if err != nil {
			Print(err)
			return err
		}
	}
	return nil
}

// client stream, server stream
func (*msg) SayHello2(s pb.Greeter_SayHello2Server) error {
	n := 0
	for {
		r, err := s.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			Print(err)
			return err
		}
		err = s.Send(&pb.HelloReply{Message: r.Name + " " + strconv.Itoa(n)})
		if err != nil {
			Print(err)
			return err
		}
		n++
	}
	return nil
}

func main() {
	go func() {
		defer func() {
			recover()
		}()

		lis, err := net.Listen("tcp", ":9999")
		CheckError(err)

		grpcServer := grpc.NewServer()
		pb.RegisterGreeterServer(grpcServer, &msg{})

		err = grpcServer.Serve(lis)
		CheckError(err)
	}()
	time.Sleep(2 * time.Second)

	conn, err := grpc.Dial(":9999", grpc.WithInsecure())
	CheckError(err)
	defer conn.Close()

	// client stream
	c := pb.NewGreeterClient(conn)
	s, err := c.SayHello(context.Background())
	CheckError(err)
	for i := 0; i < 3; i++ {
		err := s.Send(&pb.HelloRequest{Name: strconv.Itoa(i)})
		CheckError(err)
	}
	resp, err := s.CloseAndRecv()
	CheckError(err)
	Print(resp.Message)

	// server stream
	s1, err := c.SayHello1(context.Background(), &pb.HelloRequest{Name: "client"})
	CheckError(err)
	for {
		resp, err := s1.Recv()
		if err == io.EOF {
			break
		}
		CheckError(err)
		Print(resp.Message)
	}

	// client stream, server stream
	s2, err := c.SayHello2(context.Background())
	CheckError(err)
	for i := 0; i < 3; i++ {
		err := s2.Send(&pb.HelloRequest{Name: "client"})
		CheckError(err)
		resp, err := s2.Recv()
		if err == io.EOF {
			break
		}
		CheckError(err)
		Print(resp.Message)
	}
	s2.CloseSend()
}
