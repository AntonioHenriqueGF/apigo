package handler

import (
	"database/sql"

	"github.com/AntonioHenriqueGF/apigo/config"
)

var (
	logger *config.Logger
	db     *sql.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetDB()
}
