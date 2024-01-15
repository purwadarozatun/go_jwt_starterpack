package databases

import (
	"company/myproject/deps"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDB(config *deps.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("? Connected Successfully to the Database")
}

func GetDatabaseFields(out interface{}) *schema.Schema {

	s, err2 := schema.Parse(&out, &sync.Map{}, DB.NamingStrategy)
	if err2 != nil {
		panic("failed to parse schema")
	}
	return s

}
func GetFieldNames(out interface{}) []string {

	s := GetDatabaseFields(out)
	var fields []string
	for _, field := range s.Fields {
		fields = append(fields, field.DBName)
	}
	return fields

}
