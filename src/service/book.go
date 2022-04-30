package service

import (
    "github.com/43z708/gin_template/model"
	"fmt"
)

type BookService struct {}

func (BookService) SetBook(book *model.Book) error {
    result := db.Create(&book)
    if result.Error!= nil{
		return  result.Error
    }
    return nil
}


func (BookService) GetBookList() []model.Book {
    tests := make([]model.Book, 0)
    result := db.Limit(10).Find(&tests)
    if result.Error != nil {
        panic(result.Error)
    }
    return tests
}


func (BookService) GetBook(id int) model.Book {
    var book model.Book
    result := db.First(&book, id)
    if result.Error != nil {
        panic(result.Error)
    }
    return book
}

func (BookService) UpdateBook(book *model.Book) error {
    db.Save(&book)
    return nil
}

func (BookService) DeleteBook(id int) error {
    book := make([]model.Book, 0)
    result := db.First(&book, id)
	if result.Error != nil {
        panic(result.Error)
    }
    db.Delete(&book)
    return nil
}