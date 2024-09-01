package global

import (
	"golang-BE/pkg/logger"
	"golang-BE/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)
