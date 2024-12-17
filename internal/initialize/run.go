package initialize

import (
	"fmt"

	"github.com/vucong2018/study-go/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()

	fmt.Println("Loading configuration mysql", global.Config.Mysql)
	InitLogger()
	global.Logger.Info("Config Log HAHA!!", zap.String("Ok", "success"))
	InitMySql()
	InitRedis()

	r := InitRouter()
	r.Run(":6969")
}
