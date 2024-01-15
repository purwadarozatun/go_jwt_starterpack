package main

import (
	"log"
	"net/http"

	auth "company/myproject/apps/auth"
	user "company/myproject/apps/users"
	"company/myproject/databases"
	"company/myproject/deps"
	"company/myproject/docs"

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

	auth.RegisterRoute(v1, "/")

	crud := v1.Group("/crud")
	user.RegisterRoute(crud, "/user")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
