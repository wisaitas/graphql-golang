package main

import (
	"log"

	"github.com/wisaitas/graphql-golang/internal/app/app"
)

func main() {
	// สร้าง GraphQL server
	server, err := app.NewServer()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// เริ่มต้น server ที่ port 8080
	if err := server.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
