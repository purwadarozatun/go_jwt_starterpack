package user

import (
	"company/myproject/databases"

	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context, users []User) (databases.Data, error) {

	data, err := databases.Paginate(databases.DB, c, &users)

	return data, err
}
func GetDetail(c *gin.Context, user User) (User, error) {

	data := User{}
	err := databases.DB.Model(&user).Where("id = ?", c.Param("id")).First(&data).Error

	return data, err
}

func DoStore(c *gin.Context, user map[string]interface{}) (User, error) {

	data := User{}
	result := databases.DB.Model(&User{}).Create(&user)
	err := result.Error
	if err != nil {
		return data, err
	}

	errSelect := databases.DB.Model(&User{}).Where("id = ?", user["id"]).First(&data).Error
	if errSelect != nil {
		return data, errSelect
	}
	return data, err
}

func DoUpdate(c *gin.Context, id any, updated map[string]interface{}) (User, error) {

	data := User{}
	err := databases.DB.Model(&User{}).Where("id = ?", id).First(&data).Error
	if err != nil {
		return data, err
	}
	errUpdate := databases.DB.Model(&data).Updates(updated).Error
	if errUpdate != nil {
		return data, errUpdate
	}

	errSelect := databases.DB.Model(&User{}).Where("id = ?", id).First(&data).Error
	if errSelect != nil {
		return data, errSelect
	}
	return data, err
}

func DoDelete(c *gin.Context, id any) error {

	data := User{}
	err := databases.DB.Model(&User{}).Where("id = ?", id).First(&data).Error
	if err != nil {
		return err
	}
	errDelete := databases.DB.Delete(&data).Error
	if errDelete != nil {
		return errDelete
	}
	return err
}
