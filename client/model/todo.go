package model

import (
	"context"

	"github.com/grpc/todo/api"
)

// Todo type represents the model
type todo struct {
	client api.TodoServiceClient
	ctx    context.Context
}

// New creates a new todo instance
func New(ctx context.Context, client api.TodoServiceClient) todo {
	return todo{
		ctx:    ctx,
		client: client,
	}
}

// All returns all todos
func (t *todo) All() (api.TodoService_TodosClient, error) {
	response, err := t.client.Todos(t.ctx, &api.TodosRequest{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

// One returns one todo by id
func (t *todo) One(id string) (*api.Todo, error) {
	response, err := t.client.Todo(t.ctx, &api.TodoRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return response.GetTodo(), nil
}
