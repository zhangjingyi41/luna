package logger

import (
	"fmt"
	"luna/conf"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _log *zap.Logger

func Init(c *conf.LogConfig, mode string) (err error) {
	fmt.Println("日志初始化配置中......")
	// 获取日志写入器
	// lumberjack 可以帮助实现日志文件的分割和压缩
	lumberjackLogger := &lumberjack.Logger{
		Filename:   c.Filename,
		MaxSize:    c.MaxSize,
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,
	}
	logWriteSyncer := zapcore.AddSync(lumberjackLogger)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 使用 JSON 编码器
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte(c.Level)); err != nil {
		return err
	}

	var core zapcore.Core
	// 开发模式
	if mode == "dev" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, logWriteSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, logWriteSyncer, l)
	}
	_log = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(_log)
	fmt.Println("日志初始化配置完成")
	return nil
}

// gin 日志中间件
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		action := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next() // 分界点，c.Next后的代码是在接口响应后执行

		cost := time.Since(start)
		_log.Info(
			action,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("action", action),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery 是一个 Gin 框架的中间件，用于捕获并恢复 panic，防止服务器崩溃
// 参数 stack 控制是否在日志中记录堆栈信息
// 返回一个 gin.HandlerFunc 类型的中间件函数
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 检查是否为网络连接断开导致的 panic
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						// 检查错误信息是否包含 "broken pipe" 或 "connection reset by peer"
						// 这些通常表示客户端主动断开连接
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// 获取当前请求的详细信息（不包含请求体）
				httpRequest, _ := httputil.DumpRequest(c.Request, false)

				// 处理连接断开的情况
				if brokenPipe {
					_log.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					c.Error(err.(error))
					c.Abort() // 终止请求,后续的中间件将不会执行
					return
				}

				// 根据 stack 参数决定是否记录堆栈信息
				if stack {
					_log.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())), // 记录完整堆栈信息
					)
				} else {
					_log.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				// 返回 500 内部服务器错误状态码
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next() // 继续处理后续的中间件和请求处理函数
	}
}
