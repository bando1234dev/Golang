package initialize

import (
	"golang-BE/global"
	"golang-BE/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}