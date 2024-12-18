package global

import (
	"github.com/vucong2018/study-go/pkg/logger"
	"github.com/vucong2018/study-go/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)

/*
Config
Redis
Mysql
...
*/
