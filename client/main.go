package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/grpc/todo/api"
	"github.com/grpc/todo/client/todo"
	"google.golang.org/grpc"
)

func todos(client api.TodoServiceClient) error {
	response, err := client.Todos(context.Background(), &api.TodosRequest{})
	if err != nil {
		return err
	}

	for {
		todor, err := response.Recv()
		// if error is EOF
		if errors.Is(err, io.EOF) {
			log.Println("Reached EOF")
			break
		}
		log.Println(todor.GetTodo().GetId())
	}
	return nil
}

func main() {
	const PORT = 8080
	// calls server
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", PORT), grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewTodoServiceClient(conn)
	err = todos(client)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(todo.Greet())
}
