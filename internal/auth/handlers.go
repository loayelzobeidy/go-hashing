package auth

import (
	"net/http"
	"strings"
	"temp/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var accessTokenSecret = []byte("your-access-token-secret")
var refreshTokenSecret = []byte("your-refresh-token-secret")

type AuthHandler struct {
	DB *gorm.DB
}
type jsonObject struct {
	ID       int64
	Username string
}

type RefreshClaims struct {
	UUID     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}

func GenerateTokin(c *gin.Context) {

	tempObject := jsonObject{ID: 1, Username: "testuser"}
	token, err := GenerateJWT(tempObject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ah *AuthHandler) RefreshHandler(c *gin.Context) {
	refreshTokenString := c.GetHeader("Authorization")
	if refreshTokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	println("refresh token", refreshTokenString)
	tokenString := strings.Replace(refreshTokenString, "Bearer ", "", 1)

	claims, err := VerifyRefreshToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var user models.User

	ah.DB.Where("username = ?", claims.Username).First(&user)
	accessToken, _, err := GenerateTokens(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
