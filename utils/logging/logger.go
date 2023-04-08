package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Initialize() *zap.Logger {
	var s = string(os.PathSeparator)
	path, _ := os.Getwd()
	if _, err := os.Stat(path + s + "logs"); os.IsNotExist(err) {
		os.Mkdir(path+s+"logs", os.ModePerm)
	}
	if _, err := os.Stat(path + s + "logs" + s + "logs.json"); os.IsNotExist(err) {
		os.Create(path + s + "logs" + s + "logs.json")
	}
	logFile, err := os.OpenFile(path+s+"logs"+s+"logs.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}
