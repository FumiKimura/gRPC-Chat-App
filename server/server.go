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
	clients map[string]pb.ChatService_ChatServer
}

var s = &server{
	clients: make(map[string]pb.ChatService_ChatServer),
}

func main() {
	fmt.Println("Started Listening to server...")
	PORT := 8080 //Default port number
	Listen, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))

	if err != nil {
		log.Fatalf("Unable to establish connection to")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, &server{
		clients: make(map[string]pb.ChatService_ChatServer),
	})

	if err := grpcServer.Serve(Listen); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}

func (s *server) Chat(stream pb.ChatService_ChatServer) error {

	for {
		req, err := stream.Recv()
		log.Printf("From: %v, Message: %v", req.Name, req.Message)

		_, ok := s.clients[req.Name]
		if ok == false {
			s.clients[req.Name] = stream
		}
		defer delete(s.clients, req.Name)

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v", err)
		}

		for name, client := range s.clients {
			responseMessage := &pb.Message{
				Name:    req.Name,
				Message: req.Message,
			}

			var res error
			if name != req.Name {
				res = client.Send(responseMessage)
			}

			if res != nil {
				log.Fatalf("Error when sending response from server %v", res)
			}
		}
	}
}
