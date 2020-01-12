package server

import (
	"context"
	"fmt"
	"github.com/sueken5/memio/server/apis/v1"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Echo(ctx context.Context, req *apis.EchoRequest) (*apis.EchoResponse, error) {
	fmt.Println("request come")
	return &apis.EchoResponse{Text: fmt.Sprintf("hello! %s", req.Name)}, nil
}
