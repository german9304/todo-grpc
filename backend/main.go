package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/grpc/todo/api"
	"google.golang.org/grpc"
)

type service struct {
	api.UnimplementedTodoServiceServer
}

var todos []*api.Todo = []*api.Todo{
	&api.Todo{Id: "1223", Content: "this is content", Completed: false},
	&api.Todo{Id: "4569", Content: "this is new content", Completed: false},
}

func (s *service) Todos(r *api.TodosRequest, stream api.TodoService_TodosServer) error {
	for _, v := range todos {
		log.Printf("here requesting todo %v \n", v)
		stream.Send(&api.TodosResponse{Todo: v})
	}
	return nil

}

func (s *service) Todo(ctx context.Context, r *api.TodoRequest) (*api.TodoResponse, error) {
	find := func(id string) *api.Todo {
		for i := 0; i < len(todos); i++ {
			if todos[i].GetId() == id {
				return todos[i]
			}
		}
		return nil
	}

	todo := find(r.GetId())
	return &api.TodoResponse{Todo: todo}, nil
}

func main() {
	const PORT = 8080
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", PORT))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on localhost:%d", PORT)
	grpcServer := grpc.NewServer()
	api.RegisterTodoServiceServer(grpcServer, &service{})
	grpcServer.Serve(lis)
}
