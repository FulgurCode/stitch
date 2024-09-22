package main

import (
	"os"

	"github.com/FulgurCode/stitch/pkg/mysql"
	"github.com/FulgurCode/stitch/server"
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
