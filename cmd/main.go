package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eric-jacobson/chat-server/internal/db"
	"github.com/eric-jacobson/chat-server/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

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

	router := InitRoutes(&apiConfig{DB: db.New(conn)})
	router.Run("localhost:" + port)
}

func InitRoutes(apiConfig *apiConfig) *gin.Engine {
	router := gin.Default()

	router.POST("/users", apiConfig.handleCreateUser)
	router.GET("/users/:name", apiConfig.handleGetUserByName)

	return router
}

type apiConfig struct {
	DB *db.Queries
}

func (apiConfig *apiConfig) handleCreateUser(c *gin.Context) {
	var newUser users.CreateUserReq

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error parsing create new user request: %s", err))
	}
	log.Println("Create user request for:", newUser.UserName)

	user, err := apiConfig.DB.CreateUser(c, newUser.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error creating new user: %s", err))
	}

	c.JSON(http.StatusCreated, user)
}

func (apiConfig *apiConfig) handleGetUserByName(c *gin.Context) {
	user, err := apiConfig.DB.GetUser(c, c.Param("name"))
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Did not find user: %s", err))
	}

	c.JSON(http.StatusOK, user)
}
