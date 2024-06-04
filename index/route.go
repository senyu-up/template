package index

import (
	"github.com/gin-gonic/gin"
	"template/internal/service"
)

func RegisterRouter(app *gin.Engine) {
	userGroup := app.Group("/user")
	{
		userGroup.POST("/login", service.UserController.UserLogin)
		userGroup.POST("/search", service.UserController.UserSearch)
	}
}
