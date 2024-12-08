package handlers

import (
	"net/http"
	"otp-authentication-system/dto"
	"otp-authentication-system/models"
	"otp-authentication-system/utils"

	"github.com/gin-gonic/gin"
)

// RegisterUser registers a new user
func (h *Handler) RegisterUser(c *gin.Context) {
	var userRequest dto.RegisterRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input",
			"message": err.Error()})
		return
	}

	// Convert userRequest to user
	user := models.User{
		Name:              userRequest.Name,
		PhoneNumber:       userRequest.PhoneNumber,
		DateOfBirth:       userRequest.DateOfBirth,
		DeviceFingerprint: utils.GenerateFingerprint(userRequest.DeviceModel + userRequest.DeviceManufacturer + userRequest.DeviceCPUArchitecture),
	}

	// Add user to the database
	id, err := h.UserOps.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, dto.UserResponse{
		ID:      id,
		Phone:   user.PhoneNumber,
		Message: "User created successfully",
	})
}
