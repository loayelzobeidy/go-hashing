package auth

import "github.com/gin-gonic/gin"

func (ah *AuthHandler) SetupRoutes(r *gin.Engine) {
	users := r.Group("/auth")

	users.GET("/refresh", ah.RefreshHandler)

}
