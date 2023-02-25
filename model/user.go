package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Avatar   string `json:"avatar"`
}

type Login struct {
	gorm.Model

	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
