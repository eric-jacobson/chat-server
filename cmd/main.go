package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/eric-jacobson/chat-server/internal/db"
	"github.com/eric-jacobson/chat-server/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type AppConfig struct {
	DB          *db.Queries
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

	var appConfig = AppConfig{DB: db.New(conn)}

	userHandler := users.NewUserHandler(appConfig.DB)
	appConfig.UserHandler = *userHandler

	router := initRoutes(&appConfig)
	router.Run("localhost:" + port)
}

func initRoutes(appConfig *AppConfig) *gin.Engine {
	router := gin.Default()

	router.POST("/users", appConfig.UserHandler.HandleCreateUser)
	router.GET("/users/:name", appConfig.UserHandler.HandleGetUserByName)
	router.DELETE("/users/:name", appConfig.UserHandler.HandleDeleteUserByName)

	return router
}
