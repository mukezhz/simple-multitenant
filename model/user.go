package model

import (
	"html"
	"strings"

	"github.com/sachin-gautam/gin-api/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Password = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}
