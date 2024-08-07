package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupLogger() {
	logFileLocation, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}

func InitZapLogger() *zap.Logger {
	// 指定日志写入文件位置
	writeSyncer := getLogWriter()
	// 日志编码器
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// 添加一个option，作用是将函数信息记录到日志中
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return logger
}

func getEncoder() zapcore.Encoder {
	// zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()) // 普通Encoder
	// zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()) // JSON Encoder\
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	// file, _ := os.Create("./zaptest.log")
	file, _ := os.OpenFile("./zaplog.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0744)
	t := time.Now().Format("2006-01-01 15:05")
	file.WriteString(fmt.Sprintf("**%s**\n", t))
	// 将日志同时输出到文件和标准输出
	writers := io.MultiWriter(file, os.Stdout)
	return zapcore.AddSync(writers)

}
