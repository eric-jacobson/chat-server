package users

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eric-jacobson/chat-server/internal/db"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	DB *db.Queries
}

func NewUserHandler(db *db.Queries) *UserHandler {
	return &UserHandler{DB: db}
}

func (u *UserHandler) HandleCreateUser(c *gin.Context) {
	var newUser CreateUserReq

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error parsing create new user request: %s", err))
	}
	log.Println("Create user request for:", newUser.UserName)

	user, err := u.DB.CreateUser(c, newUser.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error creating new user: %s", err))
	}

	c.JSON(http.StatusCreated, user)
}

func (u *UserHandler) HandleGetUserByName(c *gin.Context) {
	user, err := u.DB.GetUser(c, c.Param("name"))
	if err != nil {
		log.Printf("Did not find user %v: %v", c.Param("name"), err)
		c.JSON(http.StatusNotFound, fmt.Sprintf("Did not find user"))
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) HandleDeleteUserByName(c *gin.Context) {
	if err := u.DB.DeleteUser(c, c.Param("name")); err != nil {
		log.Printf("Did not find user %v: %v", c.Param("name"), err)
		c.JSON(http.StatusNotFound, fmt.Sprintf("Did not find user"))
	}

	c.Status(http.StatusOK)
}
