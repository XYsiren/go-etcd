package main

import (
	"flag"
	"fmt"
	"go-etcd/echo"
	"go-etcd/echo-server/server"
	"go-etcd/etcd"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	echo.RegisterEchoServer(s, &server.EchoService{})
	etcd.CusServiceRegister("echo-service", fmt.Sprintf(":%d", *port))
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
