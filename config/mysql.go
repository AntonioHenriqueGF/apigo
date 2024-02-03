package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// InitializeMySQL creates a new MySQL database connection and returns it.
func InitializeMySQL() (*sql.DB, error) {
	logger := GetLogger("sqlite")

	// Get environment variables
	user := GetEnv("DB_USER")
	password := GetEnv("DB_PASSWORD")
	host := GetEnv("DB_HOST")
	database := GetEnv("DB_DATABASE")

	// Create the connection string
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	// Connect to the database
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		logger.Errorf("Error connecting to mysql database: %s", err)
		return nil, err
	}
	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully connected to MySQL database")

	return db, nil
}
