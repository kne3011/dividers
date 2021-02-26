package main

import (
	"com.grpc.tleu/greet/greetpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

type Server struct {
	greetpb.UnimplementedGreetServiceServer
}


func (s *Server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {

	number := int(req.GetGreeting().GetNumber())
	for number > 1 {
		for i := 2; number >= i;{
			if number % i == 0{
				number = number / i
				res := &greetpb.GreetManyTimesResponse{Result: fmt.Sprintf(strconv.Itoa(i))}
				if err := stream.Send(res); err != nil {
					log.Fatalf("error while sending greet many times responses: %v", err.Error())
				}
				time.Sleep(time.Second)
				i = number
			}
			i = i + 1
		}
	}
	return nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
