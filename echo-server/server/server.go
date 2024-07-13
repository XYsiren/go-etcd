package server

import (
	"context"
	"fmt"
	"go-etcd/echo"
)

type EchoService struct {
	echo.UnimplementedEchoServer
}

func (EchoService) UnaryEcho(ctx context.Context, in *echo.EchoMessage) (*echo.EchoMessage, error) {
	fmt.Println("server recv", in.Message)
	return &echo.EchoMessage{
		Message: "server send" + "hello client",
	}, nil
}
