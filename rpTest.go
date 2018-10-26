package main

import (
	"net"
	"log"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"

	"golang.org/x/net/context"
)
const (

	//port = ":50051"

)
func main() {

	lis, err := net.Listen("tcp", port)

	if err != nil {

		log.Fatalf("failed to listen: %v", err)

	}

	log.Println(">>> server is starting in 127.0.0.1 and port " + port + " >>>")

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &grpc.Server{})



	creds, err := credentials.NewServerTLSFromFile("D:/BaiduYunDownload/server1.pem", "D:/BaiduYunDownload/server1.key")
	TransportCredentials
	if err != nil {

		log.Println("Failed to generate credentials: ", err)

	}

	s.Serve(creds.(lis))

}