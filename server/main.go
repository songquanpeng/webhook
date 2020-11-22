package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	if os.Getenv("MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	setRouter(server)
	var port = "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		if os.Getenv("PORT") != "" {
			port = os.Getenv("PORT")
		}
	}
	fmt.Println("Server listen on port: " + port)
	log.Fatal(server.Run(":" + port))
}
