package controller

import (
	user_models "company/myproject/apps/users"
	"company/myproject/databases"
	"company/myproject/deps"
	"company/myproject/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {

	config, _ := deps.LoadConfig(".")
	var loginRequest dto.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// deps.DB.Create(&user_models.User{Name: "Javan", Email: "purwa232@javan.co.id"})
	var person user_models.User

	if config.AUTH_USE_LDAP == "false" {

		if err := databases.DB.Where(user_models.User{Email: loginRequest.Username}).First(&person).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "User not found"})
			return
		}
		errorCompare := bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(loginRequest.Password))
		if errorCompare != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Password not match"})
			return
		}

		email := person.Email
		name := person.Name
		givenToken, err := deps.CreateToken(email, name)

		if (err) == nil {
			c.JSON(200, gin.H{"message": "SUCCESS", "token": givenToken})
			return
		}

	} else {
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

		email := user["mail"]
		name := user["givenName"]
		givenToken, err := deps.CreateToken(email, name)

		if (err) == nil {
			c.JSON(200, gin.H{"message": "SUCCESS", "token": givenToken})
			return
		}

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
