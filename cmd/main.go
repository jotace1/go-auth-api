package main

import (
	"github.com/joho/godotenv"
	"github.com/jotace1/simple-authentication/internal/http/server"
)

func main() {
	_ = godotenv.Load(".env")
	server := server.NewServer()

	server.Run()
}
