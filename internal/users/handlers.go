package users

import (
	"net/http"
	"temp/internal/models"

	"temp/internal/auth"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (uh *UserHandler) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	newUser := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashedPassword),
		Age:      user.Age,
	}
	uh.DB.Create(&newUser)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (uh *UserHandler) LoginUser(c *gin.Context) {
	var userRequest models.Login
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	var storedPassword string
	result := uh.DB.Where("username = ?", userRequest.Username).First(&user)
	storedPassword = user.Password
	if result == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(userRequest.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	token, err := auth.GenerateUserJWT(userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "accessToken": token})
}
