package config

import (
	"os"

	"github.com/AntonioHenriqueGF/apigo/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		// Create the database file
		logger.Info("Database file does not exist. Creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Error creating sqlite database directory: %s", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Error creating sqlite database: %s", err)
			return nil, err
		}
		file.Close()
	}

	// Create DB and Connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Error connecting to sqlite database: %s", err)
		return nil, err
	}

	// Migrate the schema of the database. This will create the ORM schemas if they don't exist.
	err = db.AutoMigrate(&schemas.Note{})

	if err != nil {
		logger.Errorf("Error migrating sqlite database: %s", err)
		return nil, err
	}

	// Return the database connection
	return db, nil
}
