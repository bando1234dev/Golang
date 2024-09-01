package initialize

import (
	"fmt"
	"golang-BE/global"

	"go.uber.org/zap"
)

func Run(){
	//Load configuration
	LoadConfig()
	fmt.Print("Loading configuration mysql", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info(("Config log ok 1? "), zap.String("ok","success"))
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}