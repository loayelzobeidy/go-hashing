package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type jsonObject struct {
	ID       int64
	Username string
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
