package app

import (
	"fmt"
	"miniProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{
		DB: db,
	}
}

func (h *Handler) GetBooks(c *gin.Context)  {

	var books []models.Books
	
	h.DB.Find(&books)
	
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home",
		"payload": books,
		"auth": c.Query("auth"),
	})
}

func (h *Handler) GetBookById(c *gin.Context) {
	bookID := c.Param("id")
	var books models.Books

	if h.DB.Find(&books, "id = ?", bookID).RecordNotFound() {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.HTML(http.StatusOK, "books.html", gin.H{
		"title": books.Title,
		"payload": books,
		"auth": c.Query("auth"),
	})
}

func (h *Handler) AddBook(c *gin.Context) {
	c.HTML(http.StatusOK, "formBook.html", gin.H{
		"title": "Add Books",
		"auth": c.Query("auth"),
	})
}

func (h *Handler) PostBook(c *gin.Context) {
	var book models.Books
	c.Bind(&book)
	h.DB.Create(&book)
	c.Redirect(http.StatusMovedPermanently, "/books?auth="+c.PostForm("auth"))
}

func (h *Handler) UpdateBook(c *gin.Context) {
	var book models.Books
	bookID := c.Param("id")
	if h.DB.Find(&book, "id = ?", bookID).RecordNotFound(){
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": fmt.Sprintf("Book with ID %s not found", bookID),
		})
	}

	c.HTML(http.StatusOK, "formBook.html", gin.H{
		"title": "Update Books",
		"payload": book,
		"auth": c.Query("auth"),
	})
}

func (h *Handler) PutBook(c *gin.Context) {
	var books models.Books

	bookId := c.Param("id")
	if h.DB.Find(&books, "id= ?", bookId).RecordNotFound() {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error":"Book not found",
		})
	}

	//Update Data
	var reqBook = books
	c.Bind(&reqBook)

	h.DB.Model(&books).Where("id = ?", bookId).Update(&reqBook)

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/book/%s?auth=%s", bookId, c.PostForm("auth")))
}

func (h *Handler) DeleteBook(c *gin.Context) {
	var book models.Books
	bookID := c.Param("id")
	h.DB.Delete(&book, "id = ?", bookID)
	c.Redirect(http.StatusMovedPermanently, "/books?auth="+c.PostForm("auth"))
}

