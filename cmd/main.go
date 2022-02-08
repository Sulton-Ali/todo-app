package main

import (
	"log"

	"github.com/Sulton-Ali/todo-app"
	"github.com/Sulton-Ali/todo-app/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occuring while running http server: %s", err.Error())
	}
}
