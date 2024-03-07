package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/chenhangch/awesomego/gRPCdemo/gen"
	"github.com/fatih/color"
	"google.golang.org/grpc"
)

type HelloServiceImpl struct {
	gen.UnimplementedHelloServiceServer
}

func (h *HelloServiceImpl) Hello(ctx context.Context, world *gen.World) (*gen.World, error) {
	w := &gen.World{Value: "Hello :: " + world.GetValue() + " " + time.Now().String()}
	return w, nil
}

func (h *HelloServiceImpl) mustEmbedUnimplementedHelloServiceServer() {
	fmt.Println(color.GreenString("hello :: world"))
}

func (h *HelloServiceImpl) Channel(stream gen.HelloService_ChannelServer) error {
	for true {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &gen.World{Value: "Hello :: " + args.GetValue()}
		err = stream.Send(reply)

		if err != nil {
			return err
		}
	}
	return nil
}

var _ gen.HelloServiceServer = (*HelloServiceImpl)(nil)

func main() {
	server := grpc.NewServer()
	gen.RegisterHelloServiceServer(server, new(HelloServiceImpl))

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		return
	}
	err = server.Serve(listen)
	if err != nil {
		return
	}
}
