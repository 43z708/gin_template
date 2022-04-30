package main

import (
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "github.com/43z708/gin_template/controller"
    "github.com/43z708/gin_template/middleware"
)

func main(){
    engine := gin.Default()
    // ミドルウェア
    engine.Use(middleware.RecordUaAndTime)
    // CRUD 書籍
    bookEngine := engine.Group("/book")
    {
        v1 := bookEngine.Group("/v1")
        {
            v1.POST("/", controller.BookAdd)
            v1.GET("/", controller.BookList)
            v1.GET("/:id", controller.BookGet)
            v1.PUT("/:id", controller.BookUpdate)
            v1.DELETE("/:id", controller.BookDelete)
        }
    }
    engine.Run(":3000")
}