package initialize

import (
	"fmt"
	"time"

	"github.com/vucong2018/study-go/global"
	"github.com/vucong2018/study-go/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}
func InitMySql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	checkErrorPanic(err, "InitMysql initialization Error")
	global.Logger.Info("Connect MySQL Successfully")
	global.Mdb = db

	// set Pool
	SetPool()
	MigrateTables()

}

func SetPool() {
	m := global.Config.Mysql

	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("mysql error: %s::", err)
	}

	sqlDb.SetConnMaxLifetime(time.Duration(m.MaxIdleConns) * time.Second)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
}

func MigrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		fmt.Printf("Migrating table errors %s::", err)
	}
}
