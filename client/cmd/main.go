package main

import (
	"log"
	"os"

	"github.com/grpc/todo/client/app"
)

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
