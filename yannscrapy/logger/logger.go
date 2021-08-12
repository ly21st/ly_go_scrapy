package logger

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"yannscrapy/config"

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
var FileName = "scrapy.log"
var MaxSize = 100
var MaxBackups = 30
var MaxAge = 0
var Level = "info"

// 初始化
func Init(cfg *config.LogConfig) (err error) {
	if cfg == nil {
		cfg = new(config.LogConfig)
	}

	if cfg.Filename == "" {
		cfg.Filename = "scrapy.log"
	}

	if cfg.Level == "" {
		cfg.Level = "Info"
	}

	if cfg.MaxBackups == 0 {
		cfg.MaxBackups = 30
	}

	FileAndLine = cfg.FileAndLine
	FileName = cfg.Filename
	MaxSize = cfg.MaxSize
	MaxBackups = cfg.MaxBackups
	MaxAge = cfg.MaxAge
	Level = cfg.Level

	commonInit()
	return nil
}

// 初始化
func init() {
	commonInit()
}

func commonInit()  {
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

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// start := time.Now()
		c.Header("Access-Control-Allow-Origin", "*")
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		sugarLogger.Info(fmt.Sprintf("[Begin] | %s | %s | %s | %s",
			c.ClientIP(),
			c.Request.Method,
			path,
			query))

		c.Next()

		// cost := time.Since(start)
		// sugarLogger.Info(path,
		// 	zap.Int("status", c.Writer.Status()),
		// 	zap.String("method", c.Request.Method),
		// 	zap.String("path", path),
		// 	zap.String("query", query),
		// 	zap.String("ip", c.ClientIP()),
		// 	zap.String("user-agent", c.Request.UserAgent()),
		// 	zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		// 	zap.Duration("cost", cost),
		// )

		// errMsg := c.Errors.ByType(gin.ErrorTypePrivate).String()
		// if errMsg == "" {
		// 	sugarLogger.Info(fmt.Sprintf("[End] | %d | %v | %s | %s | %s | %s",
		// 		c.Writer.Status(),
		// 		cost,
		// 		c.ClientIP(),
		// 		c.Request.Method,
		// 		path,
		// 		query))
		// } else {
		// 	sugarLogger.Info(fmt.Sprintf("[End] | %d | %v | %s | %s | %s | %s | %s",
		// 		c.Writer.Status(),
		// 		cost,
		// 		c.ClientIP(),
		// 		c.Request.Method,
		// 		path,
		// 		query,
		// 		c.Errors.ByType(gin.ErrorTypePrivate).String()))
		// }
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					sugarLogger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					sugarLogger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					sugarLogger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
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
