package dto

import "time"

// UserResponse is a struct that represents the user response model
type UserResponse struct {
	ID      uint   `json:"id"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

// LoginResponse is a struct that represents the login request model
type LoginResponse struct {
	ID      uint   `json:"id"`
	OTP     string `json:"otp"`
	Message string `json:"message"`
}

// GetUserResponse is response DTO for GetUserDetails
type GetUserResponse struct {
	ID          uint      `json:"id"`
	PhoneNumber string    `json:"phone_number"`
	DateOfBirth time.Time `json:"date_of_birth"`
}
