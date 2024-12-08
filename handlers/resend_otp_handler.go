package handlers

import (
	"log"
	"net/http"
	"otp-authentication-system/dto"
	"otp-authentication-system/utils"

	"github.com/gin-gonic/gin"
)

// ResendOTP is a handler function that resends the OTP to the user
func (h *Handler) ResendOTP(c *gin.Context) {
	resendOTPReq := dto.ResendOTPRequest{}
	if err := c.BindJSON(&resendOTPReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	phoneNum := resendOTPReq.PhoneNumber

	// Check if user exists in Database
	user, err := h.UserOps.GetUserByPhoneNumber(phoneNum)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := h.OTPOps.DeleteOTPByID(user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Could not delete OTP from database",
			"message": err.Error(),
		})
		return
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
		Message: "New OTP sent successfully",
	})

}
