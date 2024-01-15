package controller

import (
	middlewares "company/myproject/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.RouterGroup, parentUrl string) {

	r.POST("/login", Login)

	protected := r.Group("/user")
	{
		protected.Use(middlewares.JwtAuthMiddleware())
		protected.GET("/profile", Profile)
	}

}
