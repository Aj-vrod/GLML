package main

import (
	"fmt"

	"github.com/GLML/pkg/db"
	"github.com/GLML/pkg/server"
)

func main() {
	fmt.Println("Connecting to database...")
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()
	fmt.Println("Successfully connected to database!")

	fmt.Println("Starting server...")
	fmt.Println("Listening on port :8080")
	server.StartServer()
}
