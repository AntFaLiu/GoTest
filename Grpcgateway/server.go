package Grpcgateway

import (
	"context"
	pb "github.com/rongfengliang/restyapp/service/pb"
	"google.golang.org/grpc"
	"net"
	"log"
)

type server struct{}

func (s *server) Echo(ctx context.Context, in *pb.EchoMessage) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Message: in.Message,
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatal("some wrong")
	}
	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("some wrong")
	}
}
