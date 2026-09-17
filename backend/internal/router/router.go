package router

import (
	"knowledge/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Knowledge Platform API"})
	})

	//认证接口
	authGroup := r.Group("/api/auth")
	authGroup.GET("/login",controller.Login)
	authGroup.POST("/register",controller.Register)
}
