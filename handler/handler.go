package handler

import (
	"github.com/AntonioHenriqueGF/apigo/config"
	"github.com/AntonioHenriqueGF/apigo/repositories"
)

var (
	logger *config.Logger
	repo   *repositories.Container
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	repo = config.GetRepository()
}
