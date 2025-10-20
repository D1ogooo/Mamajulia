package services

import (
	"mamajulia/src/database"
	"mamajulia/src/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignupService(c *gin.Context) error {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	return database.DB.CreateUser(&user).Error
}

func SigninService() {}
