package user

import (
	"company/myproject/deps"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	data, err := GetList(c, []User{})
	deps.SendData(c, data, err)
	return
}
func Detail(c *gin.Context) {
	data, err := GetDetail(c, User{})
	deps.SendData(c, data, err)
	return
}

func Store(c *gin.Context) {
	jsonData := make(map[string]interface{}, 0)
	c.BindJSON(&jsonData)
	data, err := DoStore(c, jsonData)
	deps.SendData(c, data, err)
	return
}
func Update(c *gin.Context) {
	jsonData := make(map[string]interface{}, 0)
	c.BindJSON(&jsonData)
	data, err := DoUpdate(c, c.Param("id"), jsonData)
	deps.SendData(c, data, err)
	return
}
func Delete(c *gin.Context) {
	err := DoDelete(c, c.Param("id"))
	deps.SendStatus(c, err)
	return
}
