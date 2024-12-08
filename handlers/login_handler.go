package handlers

import (
	"log"
	"net/http"
	"otp-authentication-system/dto"
	"otp-authentication-system/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// LoginUser logs in a user
func (h *Handler) LoginUser(c *gin.Context) {
	loginRequest := dto.LoginRequest{}

	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input", "message": err.Error()})
		return
	}

	// Validate user from database
	user, err := h.UserOps.GetUserByPhoneNumber(loginRequest.PhoneNumber)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	log.Printf("User found in Database\n")

	if deviceFingerPrint := utils.GenerateFingerprint(loginRequest.DeviceModel + loginRequest.DeviceManufacturer + loginRequest.DeviceCPUArchitecture); deviceFingerPrint != user.DeviceFingerprint {
		c.JSON(401, gin.H{
			"error":       "Device fingerprint mismatch",
			"fromDB":      user.DeviceFingerprint,
			"fromRequest": deviceFingerPrint,
		})
		return
	}

	log.Printf("Device fingerprints matched\n")

	// check if OTP already exists for user's phone number
	if otp, err := h.OTPOps.FetchOTPFromDB(user.PhoneNumber); err == nil {
		// OTP exists. Check if it's expired
		if otp.ExpirationTime.Before(time.Now()) {
			// TODO: Update new OTP and send back in response
			log.Println("Deleting expired OTP from database")
			// OTP is expired. Delete it from database
			if errDel := h.OTPOps.DeleteOTPByPhoneNum(user.PhoneNumber); errDel != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete expired OTP from database"})
				return
			}
		} else {
			c.JSON(http.StatusConflict, gin.H{"error": "An OTP, which is yet to expire already exists for this phone number"})
			return
		}
	}

	otp := utils.GenerateOTP()
	log.Printf("OTP generated: %s\n", otp)

	// store OTP in database
	id, err := h.OTPOps.CreateOTPinDB(user.PhoneNumber, otp)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to store OTP in database"})
		return
	}

	c.JSON(200, dto.LoginResponse{
		ID:      id,
		OTP:     otp,
		Message: "OTP created and added to DB successfully",
	})
}
