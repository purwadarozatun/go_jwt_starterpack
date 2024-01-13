package main

import (
	"log"
	"net/http"

	auth "{project_package}/controller"
	"{project_package}/databases"
	"{project_package}/deps"
	"{project_package}/docs"
	middlewares "{project_package}/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config, err := deps.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	databases.ConnectDB(&config)
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/auth/v1")
	{
		v1.POST("/login", auth.Login)
		protected := v1.Group("/user")
		protected.Use(middlewares.JwtAuthMiddleware())
		protected.GET("/profile", auth.Profile)
		// v1.POST("/read", readEndpoint)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
