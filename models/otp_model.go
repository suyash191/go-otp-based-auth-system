package models

import (
	"time"

	"gorm.io/gorm"
)

// OTPSession is a struct that represents the OTP session model
type OTPSession struct {
	ID             uint      `gorm:"primaryKey"`
	PhoneNumber    string    `gorm:"size:15;index;not null"`
	OTP            string    `gorm:"size:6;not null" validate:"len=6"`
	CreatedAt      time.Time `gorm:"not null"`
	ExpirationTime time.Time `gorm:"not null"`
	Used           bool      `gorm:"not null"`
}

// BeforeCreate is a GORM hook that sets the ExpirationTime to 3 minutes ahead of CreatedAt
func (o *OTPSession) BeforeCreate(tx *gorm.DB) (err error) {
	if o.CreatedAt.IsZero() {
		o.CreatedAt = time.Now()
	}
	o.ExpirationTime = o.CreatedAt.Add(3 * time.Minute)
	return nil
}

// OTP fetched from the database
type OTP struct {
	ID             uint
	OTP            string
	ExpirationTime time.Time
	Used           bool
}
