package handlers

import (
	"log"
	"otp-authentication-system/repository"

	"gorm.io/gorm"
)

// Handler struct holds the database connection
type Handler struct {
	UserOps repository.UserRepository
	OTPOps  repository.OTPRepository
}

// NewHandler initializes the handler with the database connection
func NewHandler(db *gorm.DB) *Handler {
	userOps := repository.NewUserRepository(db)
	log.Printf("User repository created successfully\n")
	otpOps := repository.NewOTPRepository(db)
	log.Printf("OTP repository created successfully\n")
	return &Handler{
		UserOps: userOps,
		OTPOps:  otpOps,
	}
}
