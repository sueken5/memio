package main

import (
	"fmt"
	"github.com/sueken5/memio/server"
	"github.com/sueken5/memio/server/apis/v1"
	"google.golang.org/grpc"
	"net"
)

func main() {
	app := server.NewServer()

	s := grpc.NewServer()
	apis.RegisterMemioServer(s, app)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9090))
	if err != nil {
		panic(err)
	}

	fmt.Println("start server")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
