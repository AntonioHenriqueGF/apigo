package config

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetReaderSqlx() *sqlx.DB {
	user := GetEnv("DB_USER")
	password := GetEnv("DB_PASSWORD")
	host := GetEnv("DB_HOST")
	database := GetEnv("DB_DATABASE")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=false&parseTime=true", user, password, host, database)

	var db *sqlx.DB
	var err error
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		db, err = sqlx.Connect("mysql", connectionString)
		if err == nil {
			return db
		}
		logger.Warningf("Tentativa %d: Falha ao conectar ao banco de dados: %v", i+1, err)
		// Exponential backoff: 1s, 2s, 4s, ...
		wait := 1 << i
		if wait > 30 {
			wait = 30
		}
		time.Sleep(time.Duration(wait) * time.Second)
	}
	logger.debug.Fatalf("Não foi possível conectar ao banco de dados após %d tentativas: %v", maxRetries, err)
	return nil
}

func GetWriterSqlx() *sqlx.DB {
	user := GetEnv("DB_USER")
	password := GetEnv("DB_PASSWORD")
	host := GetEnv("DB_HOST")
	database := GetEnv("DB_DATABASE")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=false&parseTime=true", user, password, host, database)
	logger.Debug(connectionString)

	var db *sqlx.DB
	var err error
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		db, err = sqlx.Connect("mysql", connectionString)
		if err == nil {
			return db
		}
		logger.Warningf("Tentativa %d: Falha ao conectar ao banco de dados: %v", i+1, err)
		wait := 1 << i
		if wait > 30 {
			wait = 30
		}
		time.Sleep(time.Duration(wait) * time.Second)
	}
	logger.debug.Fatalf("Não foi possível conectar ao banco de dados após %d tentativas: %v", maxRetries, err)
	return nil
}
