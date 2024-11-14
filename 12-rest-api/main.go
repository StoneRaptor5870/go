package main

import (
	"fmt"
	"log"
	"os"
	"rest_api/db"
	"rest_api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	port := os.Getenv("PORT")
	server.Run(fmt.Sprintf(`:%s`, port))
}
