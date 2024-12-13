package initialize

import (
	"fmt"

	"github.com/vucong2018/study-go/global"
)

func Run() {
	LoadConfig()

	fmt.Println("Loading configuration mysql", global.Config.Mysql)
	InitLogger()
	InitMySql()
	InitRedis()

	r := InitRouter()
	r.Run(":6969")
}
