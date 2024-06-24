package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/eric-jacobson/chat-server/internal/auth"
	"github.com/eric-jacobson/chat-server/internal/db"
	"github.com/eric-jacobson/chat-server/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type AppConfig struct {
	UserHandler users.UserHandler
}

func main() {
	godotenv.Load()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatal("there is no SERVER_PORT environment variable defined")
	}

	dbUrl := os.Getenv("DB_URL")
	if port == "" {
		log.Fatal("there is no DB_URL environment variable defined")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("could not connect to database at", dbUrl)
	}

	userHandler := users.NewUserHandler(db.New(conn))
	var appConfig = AppConfig{UserHandler: *userHandler}

	router := initRoutes(&appConfig)
	router.Run("localhost:" + port)
}

func initRoutes(appConfig *AppConfig) *gin.Engine {
	router := gin.Default()

	router.POST("/user/register", appConfig.UserHandler.HandleRegister)
	router.POST("/user/login", appConfig.UserHandler.HandleLogin)
	router.GET("/user/:name", auth.Authenticate, appConfig.UserHandler.HandleGetUserByName)
	router.DELETE("/user/:name", auth.Authenticate, appConfig.UserHandler.HandleDeleteUserByName)

	return router
}
