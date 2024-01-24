package main

import (
	"context"
	"google.golang.org/grpc"
	hello "grpcTest/api"
	"net"
)

type myServer struct {
	hello.UnimplementedGreeterServer
}

func (m *myServer) SayHello(a context.Context, b *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: b.Name}, nil
}

func (m *myServer) SayHelloAgain(a context.Context, b *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: "hello, hello"}, nil
}

func main() {
	listet, _ := net.Listen("tcp", ":8080")
	serverRegister := grpc.NewServer()
	myServer := &myServer{}
	hello.RegisterGreeterServer(serverRegister, myServer)
	err := serverRegister.Serve(listet)
	if err != nil {
		panic(err)
	}
}
