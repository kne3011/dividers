package main

import (
	"com.grpc.tleu/greet/greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	doManyTimesFromServer(c)
}

func doManyTimesFromServer(c greetpb.GreetServiceClient) {
	ctx := context.Background()
	fmt.Print("Enter the number:")
	var num int64
	fmt.Scan(&num)
	req := &greetpb.GreetManyTimesRequest{Greeting: &greetpb.Greeting{
		Number : num,
	}}
	stream, err := c.GreetManyTimes(ctx, req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC %v", err)
	}
	defer stream.CloseSend()

LOOP:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break LOOP
			}
			log.Fatalf("error while reciving from GreetManyTimes RPC %v", err)
		}
		log.Printf("response from GreetManyTimes:%v \n", res.GetResult())
	}

}
