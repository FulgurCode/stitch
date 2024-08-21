package main

import (
	"os"

	"github.com/VAISHAKH-GK/cloth-shop-website/server"
	"github.com/joho/godotenv"
)

func main() {
	var port string
	godotenv.Load(".env")

	if port = os.Getenv("PORT"); port == "" {
		port = "8000"
	}
	
	server.Run(port)
}
