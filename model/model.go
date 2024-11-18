package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `json:"username" gorm:"unique;not null"`
	Password string  `json:"password" gorm:"not null"`
	Entries  []Entry `json:"entries" gorm:"foreignKey:UserID"`
}

type Entry struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null"`
	Email  string `json:"email" gorm:"not null"`
	Age    int    `json:"age" gorm:"not null"`
	UserID uint   `json:"userId" gorm:"not null"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
}

type AuthenticationInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
