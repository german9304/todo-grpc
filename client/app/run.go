package app

import (
	"context"
	"fmt"
	"log"

	"github.com/grpc/todo/api"
	"github.com/grpc/todo/client/model"
	"google.golang.org/grpc"
)

const port = 8080

// Run runs the logic for the app and takes an array of strings
func Run(args []string) error {
	// calls server
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	client := api.NewTodoServiceClient(conn)

	ctx := context.Background()
	t := model.New(ctx, client)
	todos, err := t.All()
	if err != nil {
		log.Fatal(err)
		return err
	}
	r, _ := todos.Recv()
	log.Printf("%v \n", r.GetTodo())

	todo, err := t.One("1223")
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println(todo)

	return nil
}
