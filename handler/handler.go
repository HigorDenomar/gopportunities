package handler

import (
	"github.com/higordenomar/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init() {
	logger = config.GetLogger("handler")

	db = config.GetSQLite()
}
