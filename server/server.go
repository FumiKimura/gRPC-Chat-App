package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/FumiKimura/ccp2-project-polygottal/proto"
	"github.com/joho/godotenv"
)

func envLoad() {
	err := godotenv.Load("./../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

type server struct {
	pb.UnimplementedChatServiceServer
	clients map[string]pb.ChatService_ChatServer
}

var s = &server{
	clients: make(map[string]pb.ChatService_ChatServer),
}

func main() {
	envLoad()

	PORT := os.Getenv("PORT")
	Listen, err := net.Listen("tcp", ":"+PORT)
	fmt.Println("Started Listening to port:", PORT)

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
		fmt.Println(req)

		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading the stream %v", err)
		}
		_, ok := s.clients[req.Name]
		if ok == false {
			s.clients[req.Name] = stream
		}
		defer delete(s.clients, req.Name)

		for name, client := range s.clients {

			responseMessage := &pb.Message{
				Name:    req.Name,
				Message: req.Message,
			}

			var res error
			if name != req.Name {
				res = client.Send(responseMessage)
			}

			if req.Message == "!exit" {
				delete(s.clients, req.Name)
			}

			if res != nil {
				log.Fatalf("Error when sending response from server %v", res)
			}

		}
	}
}
