package routers

import (
	"gin-demo/controllers"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	authGroup := r.Group("/user")
	authGroup.Use()
	{
		authGroup.GET("/", controllers.GetUser)
	}
}
