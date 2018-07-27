package main

import (
	"log"
	"net"
	"strings"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "GoTest/customer"
)

var (
	address, isTls string
)

const (
	TRUE = "1"
)

// server is used to implement customer.CustomerServer.
type server struct {
	savedCustomers []*pb.CustomerRequest
}

// CreateCustomer creates a new Customer
func (s *server) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, in)
	return &pb.CustomerResponse{Id: in.Id, Success: true}, nil
}

// GetCustomers returns all customers by given filter
func (s *server) GetCustomers(filter *pb.CustomerFilter, stream pb.Customer_GetCustomersServer) error {
	for _, customer := range s.savedCustomers {
		if filter.Keyword != "" {
			if !strings.Contains(customer.Name, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(customer); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	//./RpcServer + port + isTls
	address = os.Args[1]
	isTls = os.Args[2]
	lis, err := net.Listen("tcp", address)
	log.Println(">>> server is starting in" + address + " >>>")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if isTls == TRUE {
		log.Println("********加密**********")
		//创建Tls服务
		creds, err := credentials.NewServerTLSFromFile("/Users/ant_oliu/go/1.8/src/GoTest/server.pem",
			"/Users/ant_oliu/go/1.8/src/GoTest/server.key")
		s := grpc.NewServer(grpc.Creds(creds))
		pb.RegisterCustomerServer(s, &server{})
		if err != nil {
			log.Println("Failed to generate credentials: ", err)
			return
		}

		s.Serve(lis)
		log.Println("服务器已经创建好")
	} else {
		// Creates a new gRPC server
		s := grpc.NewServer()
		pb.RegisterCustomerServer(s, &server{})
		s.Serve(lis)
		log.Println("服务器已经创建好")
	}
}
