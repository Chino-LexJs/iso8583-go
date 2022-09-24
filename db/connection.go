package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DSN = "user:user@tcp(127.0.0.1:3308)/pideakygo?charset=utf8mb4&parseTime=True&loc=Local"
var DB *gorm.DB

func DBconnection() {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Println("EROR 01: Error al intentar conectarse a DBF")
		log.Fatal(err)
	} else {
		log.Println("DB connected")
	}
}
