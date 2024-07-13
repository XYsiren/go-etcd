package main

import (
	"context"
	"flag"
	"fmt"
	"go-etcd/echo"
	"go-etcd/etcd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var (
// addr = flag.String("addr", "localhost:50051", "")
)

func main() {
	flag.Parse()
	etcd.CusLoadService("echo-service")
	addr := etcd.CusServiceDiscovery("echo-service")
	fmt.Println(addr)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := echo.NewEchoClient(conn)
	CallUnaryEcho(c)
}

func CallUnaryEcho(c echo.EchoClient) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	in := echo.EchoMessage{
		Message: "client say Hello World",
	}
	res, err := c.UnaryEcho(ctx, &in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Client recv", res.Message)
}
