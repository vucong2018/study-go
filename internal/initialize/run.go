package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitMySql()
	InitRedis()
	InitRouter()
}