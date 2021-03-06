package Utils

import (
	"github.com/jinzhu/gorm"
	"log"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetConnection obtiene una conexión a la base de datos
func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	return db
}