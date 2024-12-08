package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"otp-authentication-system/utils"
)

// Default values for connection
const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbName   = "db"
	sslMode  = "disable"
)

func init() {
	utils.LoadEnvVariables()
}

func getPort(key string, fallback int) (int, error) {
	if key == "PORT" {
		value, ok := os.LookupEnv(key)
		if ok {
			if port, err := strconv.Atoi(value); err == nil {
				return port, nil
			}
		}
		return fallback, nil
	}
	return 0, fmt.Errorf("invalid key %s", key)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getDatabaseURL() string {
	host := getEnv("HOST", host)
	port, err := getPort("PORT", port)
	if err != nil {
		log.Fatal("Failed to get port:", err)
	}
	user := getEnv("USERNAME", user)
	password := getEnv("PASSWORD", password)
	dbname := getEnv("DB_NAME", dbName)
	sslmode := getEnv("SSL_MODE", sslMode)

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
}

// ConnectToDatabase connects to the database
func ConnectToDatabase() *gorm.DB {
	dbURL := getDatabaseURL()
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalf("Could not connect to database: %s\n", err.Error())
	}

	return db
}
