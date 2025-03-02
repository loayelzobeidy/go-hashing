package users

import (
	"github.com/gin-gonic/gin"
)

func (uh *UserHandler) SetupRoutes(r *gin.Engine) {
	users := r.Group("/users")

	users.POST("/login", uh.LoginUser)
	users.POST("/register", uh.RegisterUser)

}
