package models

import "github.com/gin-gonic/gin"

type Inventory struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Status uint   `json:"status"`
	UserID uint   `json:"uid"`
	User   User
}

type User struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Status     uint   `json:"status"`
	Ip         string `json:"ip"`
	Inventorys []Inventory
}

func WebFile(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
}
