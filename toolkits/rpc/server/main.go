package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"net"
	"oi.io/toolkits/rpc/pb"
	"os"
	"sync"
	"time"
)

const port = "localhost:8001"

var employees = []pb.Employee{
	{
		Id:        1,
		No:        1989,
		FirstName: "Any",
		LastName:  "Chu",
		MonthSalary: &pb.MonthSalary{
			Basic: 5000,
			Bonus: 125.5,
		},
		Status: pb.EmployeeStatus_NORMAL,
		LastModified: &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
		},
	}, {
		Id:        2,
		No:        1993,
		FirstName: "Ann",
		LastName:  "Chu",
		MonthSalary: &pb.MonthSalary{
			Basic: 5000,
			Bonus: 165.5,
		},
		Status: pb.EmployeeStatus_NORMAL,
		LastModified: &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
		},
	},
}

var lock sync.Mutex

func main() {
	conn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Listen to port 8001 error", err)
	}
	creds, err := credentials.NewServerTLSFromFile("../creds/cert.pem", "../creds/key.pem")
	if err != nil {
		log.Fatal("create credentials error", err)
	}
	options := []grpc.ServerOption{grpc.Creds(creds)}
	server := grpc.NewServer(options...)
	pb.RegisterEmployeeServiceServer(server, new(EmployeeServerService))
	log.Println("Let's start with grpc started")
	err = server.Serve(conn)
	if err != nil {
		log.Fatal("Server error ", err)
	}
}

type EmployeeServerService struct{}

func (e EmployeeServerService) GetByNo(ctx context.Context, request *pb.GetByNoRequest) (*pb.EmployeeResponse, error) {
	fmt.Println("request by No", request.No)
	for _, row := range employees {
		if request.No == row.No {
			return &pb.EmployeeResponse{Employee: &row}, nil
		}
	}
	return nil, errors.New("Employee not found ")
}

func (e EmployeeServerService) GetAll(request *pb.GetAllRequest, server pb.EmployeeService_GetAllServer) error {
	for _, row := range employees {
		_ = server.Send(&pb.EmployeeResponse{Employee: &row})
		time.Sleep(time.Second * 3)
	}
	return nil
}

func (e EmployeeServerService) AddPhoto(server pb.EmployeeService_AddPhotoServer) error {
	md, ok := metadata.FromIncomingContext(server.Context())
	file, err := os.OpenFile("this.sql", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0775)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		//fmt.Println(md)
		fmt.Printf("Employee: %s\n ", md["no"][0])
	}
	for {
		recv, err := server.Recv()
		if err == io.EOF {
			return server.SendAndClose(&pb.AddPhotoResponse{IsOK: true})
		}
		if recv != nil {
			write, err := file.Write(recv.Data)
			fmt.Printf("write %d, %s", write, err)
		}
	}
}

func (e EmployeeServerService) Save(ctx context.Context, request *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	panic("implement me")
}

func (e EmployeeServerService) SaveAll(server pb.EmployeeService_SaveAllServer) error {
	for {
		recv, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		employees = append(employees, *recv.Employee)
		fmt.Println(recv.Employee)
		err = server.Send(&pb.EmployeeResponse{Employee: recv.Employee})
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, em := range employees {
		fmt.Println(em)
	}
	return nil
}
