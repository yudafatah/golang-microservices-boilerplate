package main

import (
	"github.com/yudafatah/golang-microservices-boilerplate/tree/main/Identity/internal/grpc/domain"
	"github.com/yudafatah/golang-microservices-boilerplate/tree/main/Identity/internal/grpc/service"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := "localhost:7000"

	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())

	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := service.NewRepositoryServiceClient(conn)

	for i := range [10]int{} {
		repositoryModel := domain.Repository{
			Id:        int64(i),
			IsPrivate: true,
			Name:      string("Grpc-Demo"),
			UserId:    1245,
		}

		if responseMessage, e := client.Add(context.Background(), &repositoryModel); e != nil {
			panic(fmt.Sprintf("Was not able to insert Record %v", e))
		} else {
			fmt.Println("Record Inserted..")
			fmt.Println(responseMessage)
			fmt.Println("=============================")
		}
	}
}