package models

import (
	"time"
)

// User is a struct that represents the user model
type User struct {
	ID                uint      `gorm:"primaryKey"`
	Name              string    `gorm:"size:255;not null"`
	PhoneNumber       string    `gorm:"size:15;unique;not null"`
	DateOfBirth       time.Time `gorm:"not null"`
	DeviceFingerprint string    `gorm:"size:255;not null"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
}
