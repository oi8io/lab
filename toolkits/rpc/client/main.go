package main

import (
	"fmt"
	"io"
	"log"
	"oi.io/toolkits/rpc/pb"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	// proto.Marshal()
	creds, err := credentials.NewClientTLSFromFile("../creds/cert.pem", "")
	if err != nil {
		log.Fatal("New credentials error", err)
	}
	options := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	conn, err := grpc.Dial("localhost:8001", options...)
	if err != nil {
		log.Fatal("Dial error", err)
	}
	client := pb.NewEmployeeServiceClient(conn)
	SaveAll(client)
	GetEmployee(client)
	//SendFile(client)
}

func SaveAll(client pb.EmployeeServiceClient) {
	var emps = []pb.Employee{
		{
			Id:        5,
			No:        1988,
			FirstName: "Jao",
			LastName:  "Lou",
			MonthSalary: &pb.MonthSalary{
				Basic: 4000,
				Bonus: 110,
			},
			Status: pb.EmployeeStatus_RETIRED,
			LastModified: &timestamppb.Timestamp{
				Seconds: time.Now().Unix(),
			},
		},
		{
			Id:        5,
			No:        1988,
			FirstName: "Jao",
			LastName:  "Lou",
			MonthSalary: &pb.MonthSalary{
				Basic: 4000,
				Bonus: 110,
			},
			Status: pb.EmployeeStatus_RETIRED,
			LastModified: &timestamppb.Timestamp{
				Seconds: time.Now().Unix(),
			},
		},
	}

	all, err := client.SaveAll(context.Background())
	if err != nil {
		context.Background()
	}

	finishChan := make(chan struct{})
	go func() {
		for {
			recv, err := all.Recv()
			if err == io.EOF {
				finishChan <- struct{}{}
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(recv.Employee)
		}
	}()
	fmt.Println("finish")

	for _, row := range emps {
		err := all.Send(&pb.EmployeeRequest{Employee: &row})
		if err != nil {
			break
		}
	}
	all.CloseSend()
	<-finishChan
}

func GetEmployee(client pb.EmployeeServiceClient) {
	employee, err := client.GetByNo(context.Background(), &pb.GetByNoRequest{No: 1989})
	if err != nil {
		log.Fatal("GetByNo error", err)
	}
	fmt.Println("GetByNo .... ", employee)

	all, err := client.GetAll(context.Background(), &pb.GetAllRequest{})
	for {
		recv, err := all.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println(recv)
	}
}

func SendFile(client pb.EmployeeServiceClient) {
	file, err := os.Open("/Users/ocean/pdca_script.sql")
	md := metadata.New(map[string]string{"no": "2004"})
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, md)
	photo, err := client.AddPhoto(ctx)
	if err != nil {
		log.Fatal("AddPhoto error", err)
	}
	for {
		chunk := make([]byte, 128*1024)
		chunkSize, err := file.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if chunkSize < len(chunk) {
			chunk = chunk[:chunkSize]
			break
		}
		err = photo.Send(&pb.AddPhotoRequest{Data: chunk})
		if err != nil {
			log.Fatal(err)
		}
	}
}
