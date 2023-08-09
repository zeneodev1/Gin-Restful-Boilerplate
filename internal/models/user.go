package models

import "time"

type User struct {
	ID        int `gorm:"primaryKey"`
	FullName  string
	Email     string `gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
