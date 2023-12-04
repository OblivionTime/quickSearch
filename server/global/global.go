package global

import (
	"quickSearch/config"

	"github.com/jinzhu/gorm"
)

var (
	DB     *gorm.DB
	CONFIG *config.Server
)
