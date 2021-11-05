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

func main() {
	fmt.Println("Starting client.....")
	fmt.Print("Please enter your username: ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')

	if err != nil {
		log.Printf("Failed to read from console :: %v", err)
	}
	username = strings.Trim(username, "\r\n")

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

	//go routine for sending
	go func() {
		fmt.Println("Now you can chat!")
		fmt.Println("Say something to get connected")
		fmt.Println("")
		for {
			reader := bufio.NewReader(os.Stdin)
			message, err := reader.ReadString('\n')

			if err != nil {
				log.Printf("Failed to read from console :: %v", err)
			}
			message = strings.Trim(message, "\r\n")

			sendMessage := &pb.Message{
				Name:    username,
				Message: message,
			}
			stream.Send(sendMessage)
		}
	}()

	//for loop to receive message
	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error when receiving response: %v", err)
		}

		fmt.Printf("%v ---> %v\n", res.Name, res.Message)
	}
}
