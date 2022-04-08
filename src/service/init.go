package service

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
    "github.com/43z708/gin_template/model"
    "log"
)

var DbEngine *xorm.Engine

func init()  {
    driverName := "mysql"
    DsName := "default:secret@tcp(mariadb:3306)/default"
    err := errors.New("")
    DbEngine, err = xorm.NewEngine(driverName,DsName)
    if err != nil && err.Error() != ""{
        log.Fatal(err.Error())
    }
    DbEngine.ShowSQL(true)
    DbEngine.SetMaxOpenConns(2)
    DbEngine.Sync2(new(model.Book))
    fmt.Println("init data base ok")
}