package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	pb "github.com/FumiKimura/ccp2-project-polygottal/proto"
	"google.golang.org/grpc"
)

var waitChannel = make(chan struct{})

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

func sendMessage(stream pb.ChatService_ChatClient) {
	fmt.Println("Starting client.....")
	fmt.Print("Please enter your username: ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')

	if err != nil {
		log.Printf("Failed to read from console :: %v", err)
	}
	username = strings.Trim(username, "\r\n")

	fmt.Println("===============================")
	fmt.Println("Say something to get connected")
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

	PORT := 8080
	URL := "localhost:"

	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(URL+strconv.Itoa(PORT), opts)

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
	sendMessage(stream)
}
