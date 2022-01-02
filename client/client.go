package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"os"

	pb "github.com/FumiKimura/ccp2-project-polygottal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var waitChannel = make(chan struct{})

func argsCheck() bool {
	if len(os.Args) != 3 {
		fmt.Println("You need: 1st Arg = URL, 2nd Arg = Username")
		return false
	}
	return true
}

func receiveMessage(stream pb.ChatService_ChatClient) {

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			close(waitChannel)
			break
		}

		if err != nil {
			log.Fatalf("Error when receiving response: %v", err)
		}

		fmt.Printf("%v ---> %v\n", res.Name, res.Message)
	}

}

func sendMessage(stream pb.ChatService_ChatClient, username string) {
	fmt.Println("===============================")
	fmt.Println("WELCOME TO GRPC CHAT SERVER!!!")
	fmt.Println("===============================")
	fmt.Println("Say something to get connected")
	fmt.Println("===============================")
	fmt.Println("Enter !exit to exit from chat")
	fmt.Println("===============================")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()

		if message == "!exit" {
			err := stream.CloseSend()
			if err != nil {
				log.Fatalf("Failed to exit properly", err)
			}
			break
		}

		err := stream.Send(&pb.Message{
			Name:    username,
			Message: message,
		})

		if err != nil {
			log.Fatal("Failed to send message to server", err)
		}
	}
	<-waitChannel
	fmt.Println("========See you later!!========")
}

func main() {

	ok := argsCheck()
	if !ok {
		return
	}

	URL := os.Args[1]
	tlsCredentials := credentials.NewTLS(&tls.Config{})
	conn, err := grpc.Dial(URL, grpc.WithTransportCredentials(tlsCredentials), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Unable to establish connection %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("Unable to create cilent object %v", err)
	}
	defer stream.CloseSend()

	//go routine for receiving message from server
	go receiveMessage(stream)

	//sending message from client
	sendMessage(stream, os.Args[2])
}
