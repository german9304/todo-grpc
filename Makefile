run:
	protoc -I . ./api/todo.proto --go_out=plugins=grpc:. 

build:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/todo.proto
