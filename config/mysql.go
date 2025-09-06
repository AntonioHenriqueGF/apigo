package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// InitializeMySQL creates a new MySQL database connection and returns it.
func InitializeMySQL() (*sql.DB, error) {
	logger := GetLogger("mysql")

	// Get environment variables
	user := GetEnv("MYSQL_USER")
	password := GetEnv("MYSQL_ROOT_PASSWORD")
	host := GetEnv("MYSQL_HOST")
	port := GetEnv("MYSQL_PORT")
	database := GetEnv("MYSQL_DATABASE")

	// Create the connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true&parseTime=true", user, password, host, port, database)

	// Connect to the database
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		logger.Errorf("Error connecting to MySQL database: %s", err)
		return nil, err
	}
	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		logger.Errorf("Error pinging MySQL database: %s", err)
		return nil, err
	}
	logger.Info("Successfully connected to MySQL database")

	return db, nil
}
