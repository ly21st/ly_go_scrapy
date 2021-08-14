package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

type VarArgFunc func(args ...interface{})
type TemplateVarArgFunc func(template string, args ...interface{})

var Info VarArgFunc
var Infof TemplateVarArgFunc

var Debug VarArgFunc
var Debugf TemplateVarArgFunc

var Warn VarArgFunc
var Warnf TemplateVarArgFunc

var Error VarArgFunc
var Errorf TemplateVarArgFunc

var DPanic VarArgFunc
var DPanicf TemplateVarArgFunc

var Panic VarArgFunc
var Panicf TemplateVarArgFunc

var Fatal VarArgFunc
var Fatalf TemplateVarArgFunc

var FileAndLine = true
var FileName = "app.log"
var MaxSize = 100
var MaxBackups = 30
var MaxAge = 0
var Level = "info"


// 初始化
func init()  {
	writeSyncer := getLogWriter(FileName, MaxSize, MaxBackups, MaxAge)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, writeSyncer)

	encoder := getEncoder()
	var levelEnabler = new(zapcore.Level)
	err := levelEnabler.UnmarshalText([]byte(Level))
	if err != nil {
		log.Panic()
		os.Exit(1)
	}

	fileCore := zapcore.NewCore(encoder, writeSyncer, levelEnabler)
	stdoutCore := zapcore.NewCore(encoder, os.Stdout, levelEnabler)
	core := zapcore.NewTee(fileCore, stdoutCore)
	// logger := zap.New(core, zap.AddCaller())
	var logger *zap.Logger
	if FileAndLine {
		logger = zap.New(core, zap.AddCaller())
	} else {
		logger = zap.New(core, zap.AddCallerSkip(1))
	}

	// core := zapcore.NewCore(encoder, writeSyncer, levelEnabler)
	// logger := zap.New(core, zap.AddCaller())

	sugarLogger = logger.Sugar()

	Info = sugarLogger.Info
	Infof = sugarLogger.Infof

	Debug = sugarLogger.Debug
	Debugf = sugarLogger.Debugf

	Warn = sugarLogger.Warn
	Warnf = sugarLogger.Warnf

	Error = sugarLogger.Error
	Errorf = sugarLogger.Errorf

	DPanic = sugarLogger.DPanic
	DPanicf = sugarLogger.DPanicf

	Panic = sugarLogger.Panic
	Panicf = sugarLogger.Panicf

	Fatal = sugarLogger.Fatal
	Fatalf = sugarLogger.Fatalf
}

// 编码器（如何写入日志）
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// return zapcore.NewJSONEncoder(encoderConfig)
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 日志将写到哪里去
func getLogWriter(filename string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}


func srcCodeMsg(skip int) string {
	if skip == 0 {
		skip = 1
	}
	funcName, file, line, _ := runtime.Caller(skip)
	fullFuncName := runtime.FuncForPC(funcName).Name()
	arr := strings.Split(fullFuncName, "/")

	baseName := filepath.Base(file)
	dir := filepath.Dir(file)
	parentDir := filepath.Base(dir)
	var showFileName = ""
	if parentDir != "" {
		showFileName = parentDir + "/" + baseName
	} else {
		showFileName = baseName
	}
	msg := fmt.Sprintf("%s %s:%d ", arr[len(arr)-1], showFileName, line)
	return msg
}

// func Debug(args ...interface{}) {
// 	sugarLogger.Debug(args...)
// }

// func Debugf(template string, args ...interface{}) {
// 	sugarLogger.Debugf(template, args...)
// }

// func Info(args ...interface{}) {
// 	sugarLogger.Info(args...)
// }

// func Infof(template string, args ...interface{}) {
// 	sugarLogger.Infof(template, args...)
// }

// func Warn(args ...interface{}) {
// 	sugarLogger.Warn(args...)
// }

// func Warnf(template string, args ...interface{}) {
// 	sugarLogger.Warnf(template, args...)
// }

// func Error(args ...interface{}) {
// 	sugarLogger.Error(args...)
// }

// func Errorf(template string, args ...interface{}) {
// 	msg := srcCodeMsg(2)
// 	sugarLogger.Errorf(msg+template, args...)
// }

// func DPanic(args ...interface{}) {
// 	sugarLogger.DPanic(args...)
// }

// func DPanicf(template string, args ...interface{}) {
// 	sugarLogger.DPanicf(template, args...)
// }

// func Panic(args ...interface{}) {
// 	sugarLogger.Panic(args...)
// }

// func Panicf(template string, args ...interface{}) {
// 	sugarLogger.Panicf(template, args...)
// }

// func Fatal(args ...interface{}) {
// 	sugarLogger.Fatal(args...)
// }

// func Fatalf(template string, args ...interface{}) {
// 	msg := srcCodeMsg(2)
// 	sugarLogger.Fatalf(msg+template, args...)
// }
