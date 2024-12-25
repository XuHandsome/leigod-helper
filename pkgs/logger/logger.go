package logger

import (
	"fmt"
	"github.com/XuHandsome/leigod-helper/libs"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	Logger *log.Logger
)

// InitLogger 初始化日志库，设置输出到控制台和文件
func InitLogger() error {

	// 准备存放目录 = $HOME/.leigod/$LogFilePath
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error retrieving home directory: %v", err)
	}

	logFilePath := homeDir + "/.leigod/" + libs.LogFile
	logDir := filepath.Dir(logFilePath)
	err = os.MkdirAll(logDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	// 配置 lumberjack Logger 来管理日志切割
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logFilePath, // 日志文件路径
		MaxSize:    1,           // 每个日志文件最大大小 (MB)
		MaxBackups: 7,           // 保留最多 7 个备份
		MaxAge:     7,           // 最多保留 7 天的日志文件
		Compress:   true,        // 是否压缩旧日志
	}

	// 创建 MultiWriter，将日志输出到控制台和文件
	multiWriter := io.MultiWriter(os.Stdout, lumberjackLogger)

	// 初始化 logger
	Logger = log.New(multiWriter, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	return nil
}

func Info(format string, args ...interface{}) {
	Logger.Printf("INFO: "+format, args...)
}

func Error(format string, args ...interface{}) {
	Logger.Printf("ERROR: "+format, args...)
}

func Debug(format string, args ...interface{}) {
	Logger.Printf("DEBUG: "+format, args...)
}

func Warn(format string, args ...interface{}) {
	Logger.Printf("WARN: "+format, args...)
}
