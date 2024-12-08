package repository

import (
	"otp-authentication-system/models"

	"gorm.io/gorm"
)

// UserRepository is an interface that defines the operations that can be performed on a user
type UserRepository interface {
	Create(u models.User) (uint, error)
	GetUserByID(id uint) (models.User, error)
	GetUserByPhoneNumber(phoneNumber string) (models.User, error)
}

// UserOperations is a struct that defines the operations that can be performed on a user,
// such as creating, updating, deleting, and retrieving user records from the database.
type userOperations struct {
	db *gorm.DB
}

// NewUserRepository creates and returns a new instance of UserOperations
func NewUserRepository(db *gorm.DB) UserRepository {
	return userOperations{db: db}
}

// Create creates a new user in the database and returns the ID of the created entry.
// It takes a models.User object as input and returns the ID of the created user and an error, if any.
func (r userOperations) Create(u models.User) (uint, error) {
	if err := r.db.Create(&u).Error; err != nil {
		return 0, err
	}
	return u.ID, nil
}

// GetUserByPhoneNumber retrieves a user from the database based on the phone number provided.
func (r userOperations) GetUserByID(id uint) (models.User, error) {
	var user models.User
	if err := r.db.Table("users").Where("id = ?", id).Scan(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// GetUserByPhoneNumber retrieves a user from the database based on the phone number provided.
func (r userOperations) GetUserByPhoneNumber(phoneNumber string) (models.User, error) {
	var user models.User
	if err := r.db.Table("users").Where("phone_number = ?", phoneNumber).Scan(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
