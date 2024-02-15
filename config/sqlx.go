package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Get environment variables

func GetReaderSqlx() *sqlx.DB {

	// Get environment variables
	user := GetEnv("DB_USER")
	password := GetEnv("DB_PASSWORD")
	host := GetEnv("DB_HOST")
	database := GetEnv("DB_DATABASE")

	// Create the connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true&parseTime=true", user, password, host, database)

	reader := sqlx.MustConnect("mysql", connectionString)

	return reader
}

func GetWriterSqlx() *sqlx.DB {

	// Get environment variables
	user := GetEnv("DB_USER")
	password := GetEnv("DB_PASSWORD")
	host := GetEnv("DB_HOST")
	database := GetEnv("DB_DATABASE")

	// Create the connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true&parseTime=true", user, password, host, database)

	writer := sqlx.MustConnect("mysql", connectionString)

	return writer
}
