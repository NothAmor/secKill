package common

import (
	"secKill/app/config"

	"github.com/charmbracelet/log"
	"github.com/redis/rueidis/rueidiscompat"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Logger *log.Logger
	DB     *gorm.DB
	Redis  rueidiscompat.Cmdable
)
