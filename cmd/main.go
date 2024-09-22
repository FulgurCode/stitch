package main

import (
	"os"

	"github.com/VAISHAKH-GK/benster-website/pkg/mysql"
	"github.com/VAISHAKH-GK/benster-website/server"
	"github.com/joho/godotenv"
)

func main() {
	var port string
	godotenv.Load(".env")

	if port = os.Getenv("PORT"); port == "" {
		port = "8000"
	}

	mysql.Connect()

	server.Run(port)
}
