package handlers

import (
	"net/http"
	"otp-authentication-system/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserDetails is a handler function that returns the details of a user
func (h *Handler) GetUserDetails(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	phoneNumber := c.Query("phone_number")

	if err != nil && phoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id or phone_number is required"})
		return
	}
	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if id > 0 {
		user, err := h.UserOps.GetUserByID(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userGet := dto.GetUserResponse{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			DateOfBirth: user.DateOfBirth,
		}

		c.JSON(http.StatusOK, userGet)
		return
	}

	if phoneNumber != "" {
		user, err := h.UserOps.GetUserByPhoneNumber(phoneNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userGet := dto.GetUserResponse{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			DateOfBirth: user.DateOfBirth,
		}
		c.JSON(http.StatusOK, userGet)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid phone_number"})
}
