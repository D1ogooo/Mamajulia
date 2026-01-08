package services

import (
	"errors"
	"mamajulia/internal/database"
	"mamajulia/internal/models"
	"mamajulia/pkg/jwt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignupService(c *gin.Context) error {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return err
	}

	var existingUser models.User
	if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("email já cadastrado")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	if user.Role == "" {
		user.Role = "user"
	}

	return database.DB.Create(&user).Error
}

func SigninService(c *gin.Context) (string, error) {
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		return "", err
	}

	var user models.User
	if err := database.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		return "", errors.New("email ou senha incorretos")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return "", errors.New("email ou senha incorretos")
	}

	token, err := jwt.GenerateJWT(user.Email, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
