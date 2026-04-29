package main

import (
	"fmt"
	"gin-contact-list/config"
	"gin-contact-list/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	PORT := os.Getenv("PORT")

	config.ConnectDB()
	app := gin.Default()

	routes.RegisterRoutes(app)

	app.Run(fmt.Sprintf(":%v", PORT))
}
