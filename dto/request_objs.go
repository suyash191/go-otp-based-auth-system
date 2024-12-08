package dto

import "time"

// RegisterRequest is a struct that represents the user request model
type RegisterRequest struct {
	Name                  string    `json:"name" binding:"required"`
	DateOfBirth           time.Time `json:"date_of_birth" binding:"required"`
	PhoneNumber           string    `json:"phone_number" binding:"required"`
	DeviceModel           string    `json:"device_model" binding:"required"`
	DeviceManufacturer    string    `json:"device_manufacturer" binding:"required"`
	DeviceCPUArchitecture string    `json:"device_cpu_architecture" binding:"required"`
}

// LoginRequest is a struct that represents the login request model
type LoginRequest struct {
	PhoneNumber           string `json:"phone_number" binding:"required"`
	DeviceModel           string `json:"device_model" binding:"required"`
	DeviceManufacturer    string `json:"device_manufacturer" binding:"required"`
	DeviceCPUArchitecture string `json:"device_cpu_architecture" binding:"required"`
}

// ResendOTPRequest shows request object for ResendOTP
type ResendOTPRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

// VerifyRequest is a struct that represents the request body for verifying OTP
type VerifyRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	OTP         string `json:"otp" binding:"required"`
}
