package main

import (
	"log"

	srv "github.com/echodiv/test_todo_rest/internal/app/tasks_srv"
)

func main() {
	server := srv.NewServer()
	if err := server.StartServer(); err != nil {
		log.Fatal(err.Error())
	}

}
