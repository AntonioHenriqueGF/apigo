package handler

import (
	"database/sql"

	"github.com/AntonioHenriqueGF/apigo/config"
)

var (
	Logger *config.Logger
	Db     *sql.DB
)

func InitializeHandler() {
	Logger = config.GetLogger("handler")
	Db = config.GetDB()
}
