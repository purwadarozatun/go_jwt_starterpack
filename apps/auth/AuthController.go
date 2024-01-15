package controller

import (
	user_models "company/myproject/apps/users"
	"company/myproject/databases"
	"company/myproject/deps"
	"company/myproject/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var loginRequest dto.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// deps.DB.Create(&user_models.User{Name: "Javan", Email: "purwa232@javan.co.id"})
	var person user_models.User

	ok, user, err := deps.AuthenticateLdapUser(loginRequest)
	if err != nil {
		c.JSON(500, gin.H{"message": "ERROR LDAP", "error": err})
		return
	}
	if !ok {
		c.JSON(401, gin.H{"message": "User not found"})
		return
	}

	if err := databases.DB.Where(user_models.User{Email: user["mail"]}).First(&person).Error; err != nil {
		fmt.Println("Record not found", "creating new user")
		databases.DB.Create(&user_models.User{Name: user["givenName"], Email: user["mail"]})
	}
	token, err := deps.CreateToken(user["mail"], user["givenName"])
	if (err) == nil {
		c.JSON(200, gin.H{"message": "SUCCESS", "token": token})
		return
	} else {
		c.JSON(500, gin.H{"message": "TOKEN_GENERATE_ERROR ", "token": token})
		return
	}
}

func Profile(c *gin.Context) {

	email, err := deps.ExtractTokenID(c)
	if err != nil {
		c.JSON(200, gin.H{"message": "TOKEN_NOT_VALID"})
	}

	fmt.Println(email)

	var person user_models.User

	if err := databases.DB.Where(user_models.User{Email: email}).First(&person).Error; err != nil {
		c.JSON(200, gin.H{"message": "USER_NOT_FOUND"})
		return
	}

	c.JSON(200, gin.H{"message": "SUCCESS", "data": person})
	return
}
