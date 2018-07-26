package main

import (
	"io"
	"log"
	"os"

	pb "Customer"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials"
)

var (
	address, isTls string
)

const (
	TRUE = "1"
)

// createCustomer calls the RPC method CreateCustomer of CustomerServer
func createCustomer(client pb.CustomerClient, customer *pb.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added with id: %d", resp.Id)
	}
}

// getCustomers calls the RPC method GetCustomers of CustomerServer
func getCustomers(client pb.CustomerClient, filter *pb.CustomerFilter) {
	// calling the streaming API
	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get customers: %v", err)
	}
	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
		}
		log.Printf("Customer: %v", customer)
	}
}

func main() {
	address = os.Args[1]
	isTls = os.Args[2]
	var lis *grpc.ClientConn
	var err error
	if isTls == TRUE {
		log.Println("******加密*****")
		creds, err := credentials.NewClientTLSFromFile("/Users/ant_oliu/go/1.8/src/LypTest/server.pem",
			"localhost")
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
			return
		}
		log.Println(address)
		lis, err = grpc.Dial(address, grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Fatalln(err)
			return
		}
	} else {
		lis, err = grpc.Dial(address, grpc.WithInsecure()) //连接grpc服务器
		if err != nil {
			log.Fatalf("did not connect: %v", err)
			return
		}
	}

	defer lis.Close()                   //工作做完之后要记得关闭conn
	client := pb.NewCustomerClient(lis) //创建一个新的customerClient
	customer := &pb.CustomerRequest{//构造CustomerRequest结构体
		Id: 01,
		Name: "Lyp",
		Email: "yupeng.liu02@ele.me",
		Phone: "182-9218-0367",
		Addresses: []*pb.CustomerRequest_Address{
			{
				Street:            "lalala",
				City:              "Shanghai",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: false,
			},
			{
				Street:            "sisisi",
				City:              "Kochi",
				State:             "KL",
				Zip:               "68356",
				IsShippingAddress: true,
			},
		},
	}

	log.Println("****create*****")
	createCustomer(client, customer) //创建新的customer

	customer = &pb.CustomerRequest{
		Id:    02,
		Name:  "yyy",
		Email: "yyy@xyz.com",
		Phone: "732-757-2924",
		Addresses: []*pb.CustomerRequest_Address{
			{
				Street:            "lalala",
				City:              "Shenzhen",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: true,
			},
		},
	}

	createCustomer(client, customer)          //创建新的customer
	filter := &pb.CustomerFilter{Keyword: ""} //过滤调没有意义的关键字
	getCustomers(client, filter)
}