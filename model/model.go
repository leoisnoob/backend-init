package model

import "time"

type User struct {
	ID        int
	UserID    string `gorm:"uniqueIndex;size:20"`
	Pass      string
	IsRoot    int `gorm:"size:1"`
	LastLogin *time.Time
	CreatedAt time.Time
}
