package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(localhost:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database Connection Error", err)
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	router.GET("/books", bookHandler.GetBooks)
	router.GET("/books/:id", bookHandler.GetBook)
	router.POST("/books", bookHandler.PostBooksHandler)
	router.PUT("/books/:id", bookHandler.UpdateBooksHandler)
	router.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	// // API Versioning
	// v1 := router.Group("/v1")
	// v1.GET("/address", handler.AddressHandler)

	router.Run(":8083")
}
