package main

import (
	"log"
	"otp-authentication-system/database"
	"otp-authentication-system/handlers"
	"otp-authentication-system/router"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	db := database.ConnectToDatabase()

	log.Printf("Database initialized, and connection established")

	router := router.SetupRouter()
	log.Printf("Router setup completed")

	handler := handlers.NewHandler(db)
	log.Printf("New handler created")

	router.GET("/user", handler.GetUserDetails)
	router.POST("/register", handler.RegisterUser)
	router.POST("/login", handler.LoginUser)
	router.POST("/verify-otp", handler.VerifyOTP)
	router.POST("/login/resend-otp", handler.ResendOTP)

	log.Printf("Routes registered")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
