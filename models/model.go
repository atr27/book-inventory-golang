package models

type Books struct {
	ID          int    `json:"id" form:"id" gorm:"primary_key"`
	Title       string `json:"title" form:"title" binding:"required"`
	Author      string `json:"author" form:"author" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Stock       int    `json:"stock" form:"stock" binding:"required"`
}

type Login struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

const (
	USER = "admin"
	PASSWORD = "admin123"
	SECRET = "secret"
)