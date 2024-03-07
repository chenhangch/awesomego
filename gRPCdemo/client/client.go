package main

import (
	"context"
	"fmt"
	"github.com/chenhangch/awesomego/gRPCdemo/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()

	client := gen.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &gen.World{Value: "client"})
	if err != nil {
		return
	}
	fmt.Println(reply.GetValue())

	stream, err := client.Channel(context.Background())
	if err != nil {
		return
	}
	go func() {
		for true {
			if err := stream.Send(&gen.World{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for true {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
