package main

import (
	"miniProject/app"
	"miniProject/auth"
	"miniProject/db"
	"miniProject/middleware"

	"github.com/gin-gonic/gin"
)

// var router *gin.Engine

func main() {
	conn := db.InitDB()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	handler := app.New(conn)

	//Home
	router.GET("/", auth.HomeHandler)
	//AllBooks
	router.GET("/books", middleware.AuthValidation, handler.GetBooks)

	//Get Book By Id
	router.GET("/book/:id", middleware.AuthValidation, handler.GetBookById)

	//Add Book
	router.GET("/addBook", middleware.AuthValidation, handler.AddBook)
	router.POST("/book", middleware.AuthValidation, handler.PostBook)

	//Update Book
	router.GET("/updateBook/:id", middleware.AuthValidation, handler.UpdateBook)
	router.POST("/updateBook/:id", middleware.AuthValidation, handler.PutBook)
	
	//Delete Book
	router.POST("/deleteBook/:id", middleware.AuthValidation, handler.DeleteBook)

	//Auth
	router.GET("/login", auth.LoginGetHandler)
	router.POST("/login", auth.LoginPostHandler)

	router.Run()
}