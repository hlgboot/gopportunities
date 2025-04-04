package handler

import (
	"github.com/hlgboot/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("hanlder")
	db = config.GetSQLite()
}
