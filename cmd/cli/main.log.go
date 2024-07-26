package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main (){
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name:%s, age:%d", "TipsGo",40)

	// //logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "TipsGo"), zap.Int("age",40))
	
	// logger := zap.NewExample()
	// logger.Info("Hello")

	// // Dev
	// logger,_ = zap.NewDevelopment()
	// logger.Info("Hello Dev")

	// // Production

	// logger,_ = zap.NewProduction()
	// logger.Info("Hello Prod")

	// 3. Customize
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log:: ", zap.Int("Line", 1))
	logger.Error("Error log:: ", zap.Int("Line", 2))

}


func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts - Time
	encodeConfig.TimeKey = "time"
	// from info -> INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// caller cli/main.log.go:23
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

// Write

func getWriterSync() zapcore.WriteSyncer{
	file,_ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}