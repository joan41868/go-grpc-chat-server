package main

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"grpc-chat/proto"
	"log"
	"net"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	// server should keep the other clients alive and make them wait for nem messages for them
	lis, _ := net.Listen("tcp", ":"+os.Getenv("PORT"))

	baseServer := grpc.NewServer()
	srv := proto.NewChatServer()
	proto.RegisterChatServiceServer(baseServer, srv)
	baseServer.Serve(lis)
}
