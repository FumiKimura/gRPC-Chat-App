package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	pb "github.com/FumiKimura/ccp2-project-polygottal/proto"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func main() {
	fmt.Println("Started Listening to server...")
	PORT := 8080 //Default port number
	Listen, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))

	if err != nil {
		log.Fatalf("Unable to establish connection to")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(Listen); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}

func (*server) Chat(stream pb.ChatService_ChatServer) error {

	for {
		req, err := stream.Recv()

		fmt.Println(req)

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v", err)
		}

		//First build Echo server
		res := stream.Send(&pb.Message{
			Name:    req.Name,
			Message: req.Message,
		})

		if res != nil {
			log.Fatalf("Error when sending response from server %v", res)
		}
	}

}
