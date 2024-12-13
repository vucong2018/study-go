package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/vucong2018/study-go/global"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	//
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Failed to read configuration %v", err)
	}
	//
	fmt.Println("Server Port::", viper.GetInt("server.port"))

	// config struct

	// var config Config
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode config %v", err)
	}
}
