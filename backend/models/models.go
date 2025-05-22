package models

import (
	"time"
)

type Entry struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    string `gorm:"index;not null"` // ID пользователя из SuperTokens
	Title     string `gorm:"size:255"`       // Заголовок записи
	Content   string `gorm:"type:text"`      // Текст записи
	CreatedAt time.Time
	UpdatedAt time.Time
}
