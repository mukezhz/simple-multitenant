package model

import (
	"errors"

	"github.com/sachin-gautam/gin-api/database"
	"gorm.io/gorm"
)

func (entry *Entry) Save() (*Entry, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}

func FindEntryByID(userID uint, entryID string) (Entry, error) {
	var entry Entry
	err := database.Database.Where("id = ? AND user_id = ?", entryID, userID).First(&entry).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return Entry{}, errors.New("entry not found")
		}
		return Entry{}, err
	}

	return entry, nil
}
