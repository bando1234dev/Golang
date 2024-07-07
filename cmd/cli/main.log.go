package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name: %s, age:%d", "TipsGo", 40)
	// // fmt.Printf()
	// // logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name","jqk"),zap.Int("age", 40))

	// logger := zap.NewExample()
	// logger.Info("Hello")

	// //Dev
	// logger,_ = zap.NewDevelopment()
	// logger.Info("HEllo new Development")

	// //Prod
	// logger,_ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	// 3.
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder,sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	//ts -> time
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// at line?
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt",os.O_CREATE | os.O_WRONLY,os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole,syncFile)
}