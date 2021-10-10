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
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	lis, _ := net.Listen("tcp", ":"+port)

	baseServer := grpc.NewServer()
	srv := proto.NewChatServer()
	proto.RegisterChatServiceServer(baseServer, srv)
	log.Println("gRPC server listening on :" + port)
	baseServer.Serve(lis)
}
