package main

import (
	"bufio"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/Gufhans/gRPC_Gufhans/chat"
)

func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	Scanner.Scan()
	messagetosend := Scanner.Text()

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer conn.Close()

	c := chat.NewSendMessageClient(conn)

	message := chat.Message{
		Body: messagetosend,
	}

	response, err := c.SayFuckJohan(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)

}
