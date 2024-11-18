package model

import "github.com/sachin-gautam/gin-api/database"

func (entry *Entry) Save() (*Entry, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}
