package encrypt

import (
	"net/http"
	"temp/internal/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	encrypted := r.Group("/encrypted")
	encrypted.Use(auth.AuthMiddleware())
	{
		encrypted.GET("/resource", func(c *gin.Context) {
			claims, _ := c.Get("claims")
			c.JSON(http.StatusOK, gin.H{"message": "Protected resource accessed", "claims": claims})
		})
		encrypted.POST("/encrypt", EncryptHandler)
		encrypted.POST("/decrypt", DecryptHandler)
		encrypted.POST("/hash", HashingHandler)
	}

}
