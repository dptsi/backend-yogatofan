package handlers

import (
	"coba-db-mssql/repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserHandler struct {
	userRepository repositories.UserRepositoryInterface
}

func NewUserHandler(userRepository *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		userRepository: userRepository,
	}
}

func (u UserHandler) GetAll(c *gin.Context) {
	userSlice, err := u.userRepository.AllUser()

	if err != nil {
		log.Fatal("AllUsers failed:", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"result": userSlice,
		"count":  len(userSlice),
	})
}
