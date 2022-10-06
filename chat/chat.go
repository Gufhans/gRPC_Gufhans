package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedSendMessageServer
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Hello From the Server"}, nil
}

func (s *Server) SayFuckJohan(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received new message: %s", message.Body)
	return &Message{Body: "Received a very important message"}, nil
}
