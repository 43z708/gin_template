package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/43z708/gin_template/model"
    "github.com/43z708/gin_template/service"
    "strconv"
	"fmt"
)

func BookAdd(c *gin.Context) {
	book := model.Book{}
	err := c.Bind(&book)
	if err != nil{
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService :=service.BookService{}
	err = bookService.SetBook(&book)
	if err != nil{
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func BookList(c *gin.Context){
	bookService :=service.BookService{}
	BookLists := bookService.GetBookList()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data": BookLists,
	})
}


func BookGet(c *gin.Context){
	// :idから値を取り出す
	id := c.Params.ByName("id")
	intId, _ := strconv.ParseInt(id, 10, 0)
	bookService :=service.BookService{}
	Book := bookService.GetBook(int(intId))
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data": Book,
	})
}

func BookUpdate(c *gin.Context){
	// :idから値を取り出す
	id := c.Params.ByName("id")
	intId, err := strconv.ParseUint(id, 10, 0)
	bookService :=service.BookService{}
	book := bookService.GetBook(int(intId))
	// update
	err = c.Bind(&book)
	fmt.Println(book)
	if err != nil{
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	err = bookService.UpdateBook(&book)
	if err != nil{
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func BookDelete(c *gin.Context){
	// :idから値を取り出す
	id := c.Params.ByName("id")
	intId, err := strconv.ParseInt(id, 10, 0)
	if err != nil{
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService :=service.BookService{}
	err = bookService.DeleteBook(int(intId))
	if err != nil{
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}