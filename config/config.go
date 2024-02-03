package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	db     *sql.DB
	logger *Logger
)

func Init() error {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		logger.Errorf("Error loading .env file: %v", err)
	}
	logger.Infof("Loaded .env file")

	// Initialize Database Connection
	db, err = InitializeMySQL()

	db.Stats()

	if err != nil {
		return fmt.Errorf("Error initializing database: %v", err)
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetLogger(p string) *Logger {
	// Initialized Logger
	logger = NewLogger(p)
	return logger
}
