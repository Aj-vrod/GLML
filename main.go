package main

import (
	"fmt"

	"github.com/GLML/pkg/server"
)

func main() {
	fmt.Println("Starting server...")
	fmt.Println("Listening on port :8080")

	server.StartServer()
}
