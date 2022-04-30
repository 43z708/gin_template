package service

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
    "github.com/43z708/gin_template/model"
    "log"
)

var db *gorm.DB

func init()  {
    DsName := "default:secret@tcp(mariadb:3306)/default?charset=utf8mb4&parseTime=True&loc=Local"
	err := errors.New("")
    db, err = gorm.Open(mysql.Open(DsName), &gorm.Config{})
    if err != nil && err.Error() != ""{
        log.Fatal(err.Error())
    }
	db.AutoMigrate(&model.Book{})
    fmt.Println("init data base ok")
}