package handlers

import (
	"net/http"
	"otp-authentication-system/dto"
	"time"

	"github.com/gin-gonic/gin"
)

// VerifyOTP verifies OTP
func (h *Handler) VerifyOTP(c *gin.Context) {
	var req dto.VerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input body", "message": err.Error()})
		return
	}

	// Fetch OTP for requested phone number
	otp, err := h.OTPOps.FetchOTPFromDB(req.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch OTP", "message": err.Error()})
		return
	}

	// check if otp is expired
	if otp.ExpirationTime.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "otp expired"})
		return
	}

	// check if otp is already used
	if otp.Used {
		c.JSON(http.StatusBadRequest, gin.H{"error": "otp already used"})
		return
	}

	// check if otp matches the one provided by the user
	if otp.OTP != req.OTP {
		c.JSON(http.StatusBadRequest, gin.H{"error": "otp mismatch"})
		return
	}

	// mark otp as used
	if err := h.OTPOps.UpdateOTPUsed(otp.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update OTP to true", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "otp verified"})
}
