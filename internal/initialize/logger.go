package initialize

import (
	"github.com/vucong2018/study-go/global"
	"github.com/vucong2018/study-go/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
