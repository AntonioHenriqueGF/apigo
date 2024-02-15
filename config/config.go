package config

import (
	"database/sql"
	"os"

	"github.com/AntonioHenriqueGF/apigo/repositories"
	"github.com/joho/godotenv"
)

var (
	db     *sql.DB
	logger *Logger
	repo   *repositories.Container
)

func Init() error {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		logger.Errorf("Error loading .env file: %v", err)
	}
	logger.Infof("Loaded .env file")

	// Initialize Database Repositories

	repo = repositories.New(repositories.Options{
		WriterSqlx: GetWriterSqlx(),
		ReaderSqlx: GetReaderSqlx(),
	})

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

func GetRepository() *repositories.Container {
	return repo
}
