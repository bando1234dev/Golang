package initialize

import (
	"fmt"
	"golang-BE/global"
)

func Run(){
	//Load configuration
	LoadConfig()
	fmt.Print("Loading configuration mysql", global.Config.Mysql.Username)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}