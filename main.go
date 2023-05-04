package main

import (
	"coba-db-mssql/adapter"
	"coba-db-mssql/handlers"
	"coba-db-mssql/repositories"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	conn := adapter.ConnectDB()

	userRepository := repositories.NewUserRepository(conn)
	userHandler := handlers.NewUserHandler(userRepository)
	mode := os.Getenv("GIN_MODE")

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.GET("/all-users", userHandler.GetAll)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("Defaulting to port %s", port)
	}

	log.Fatal(router.Run(":" + port))
}
