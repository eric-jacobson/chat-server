package users

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/eric-jacobson/chat-server/internal/auth"
	"github.com/eric-jacobson/chat-server/internal/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	DB *db.Queries
}

func NewUserHandler(db *db.Queries) *UserHandler {
	return &UserHandler{DB: db}
}

func (u *UserHandler) HandleRegister(c *gin.Context) {
	var newUser UserReq

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("error parsing  request body: %s", err))
	}
	log.Println("Create user request for:", newUser.UserName)

	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("error creating new user: %s", err))
	}

	user, err := u.DB.CreateUser(c, db.CreateUserParams{UserName: newUser.UserName, PasswordHash: hashedPassword})
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("error creating new user: %s", err))
	}

	c.JSON(http.StatusCreated, UserResp{UserName: user.UserName, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt})
}

func (u *UserHandler) HandleLogin(c *gin.Context) {
	var user UserReq
	if err := c.BindJSON(&user); err != nil {
		log.Printf("error parsing request body: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error parsing request body",
		})
	}

	storedUser, err := u.DB.GetUser(c, user.UserName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "invalid user name or password",
		})
		return
	}
	if storedUser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "invalid user name or password",
		})
		return
	}

	if err := checkHash(user.Password, storedUser.PasswordHash); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "invalid user name or password",
		})
		return
	}

	token, err := auth.GenerateToken(storedUser.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error generating token",
		})
		return
	}

	log.Printf("user %s logged in at %s\n", storedUser.UserName, time.Now().UTC())
	c.JSON(http.StatusOK, LoginResp{AccessToken: token})
}

func (u *UserHandler) HandleGetUserByName(c *gin.Context) {
	user, err := u.DB.GetUser(c, c.Param("name"))
	if err != nil {
		log.Printf("did not find user %v: %v", c.Param("name"), err)
		c.JSON(http.StatusNotFound, fmt.Sprintf("did not find user"))
	}

	c.JSON(http.StatusOK, UserResp{UserName: user.UserName, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt})
}

func (u *UserHandler) HandleDeleteUserByName(c *gin.Context) {
	if err := u.DB.DeleteUser(c, c.Param("name")); err != nil {
		log.Printf("did not find user %v: %v", c.Param("name"), err)
		c.JSON(http.StatusNotFound, fmt.Sprintf("did not find user"))
	}

	c.Status(http.StatusOK)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password %v", err.Error())
	}
	return string(hash), nil
}

func checkHash(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
