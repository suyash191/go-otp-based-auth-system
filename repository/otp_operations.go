package repository

import (
	"otp-authentication-system/models"

	"gorm.io/gorm"
)

// OTPRepository is an interface that defines the operations that can be performed on an OTP
type OTPRepository interface {
	CreateOTPinDB(phoneNumber string, otp string) (uint, error)
	FetchOTPFromDB(phoneNumber string) (models.OTP, error)
	UpdateOTPUsed(id uint) error
	DeleteOTPByID(id uint) error
	DeleteOTPByPhoneNum(phoneNumber string) error
}

type otpOperations struct {
	DB *gorm.DB
}

// NewOTPRepository creates and returns a new instance of OTPRepository
func NewOTPRepository(db *gorm.DB) OTPRepository {
	return otpOperations{DB: db}
}

func (o otpOperations) CreateOTPinDB(phoneNumber string, otp string) (uint, error) {
	otpSession := models.OTPSession{
		PhoneNumber: phoneNumber,
		OTP:         otp,
		Used:        false,
	}

	if err := o.DB.Table("otp_sessions").Create(&otpSession).Error; err != nil {
		return 0, err
	}

	return otpSession.ID, nil
}

func (o otpOperations) FetchOTPFromDB(phoneNumber string) (models.OTP, error) {
	var otp models.OTP
	// Get columns OTP and ExpirationTime from the OTPSession table where PhoneNumber is equal to the phoneNumber
	if err := o.DB.Table("otp_sessions").Select("id, otp, expiration_time, used").Where("phone_number = ?", phoneNumber).Scan(&otp).Error; err != nil {
		return models.OTP{}, err
	}
	return otp, nil
}

// UpdateOTPUsed updates the Used column to true for the given OTP ID
func (o otpOperations) UpdateOTPUsed(id uint) error {
	// Update the Used column in the OTPSession table where the ID is equal to the id
	if err := o.DB.Table("otp_sessions").Where("id = ?", id).Update("used", true).Error; err != nil {
		return err
	}
	return nil
}

// DeleteOTPById deletes the OTP record from the OTPSession table where the ID is equal to the id
func (o otpOperations) DeleteOTPByID(id uint) error {
	// Delete the record from the OTPSession table where the PhoneNumber is equal to the phoneNumber
	if err := o.DB.Table("otp_sessions").Where("id = ?", id).Delete(&models.OTPSession{}).Error; err != nil {
		return err
	}
	return nil
}

func (o otpOperations) DeleteOTPByPhoneNum(phoneNumber string) error {
	// Delete the record from the OTPSession table where the PhoneNumber is equal to the phoneNumber
	if err := o.DB.Table("otp_sessions").Where("phone_number = ?", phoneNumber).Delete(&models.OTPSession{}).Error; err != nil {
		return err
	}
	return nil
}
