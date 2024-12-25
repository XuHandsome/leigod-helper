//go:generate goversioninfo
package main

import (
	"fmt"
	"github.com/XuHandsome/leigod-helper/pkgs/config"
	"github.com/XuHandsome/leigod-helper/pkgs/leigod"
	"github.com/XuHandsome/leigod-helper/pkgs/logger"
	"github.com/XuHandsome/leigod-helper/pkgs/shutdown"
)

func main() {
	// 初始化日志库
	err := logger.InitLogger()
	if err != nil {
		fmt.Printf("初始化Logger失败: %v\n", err)
		return
	}

	User, callStop, err := config.InitConf()
	if err != nil {
		return
	}

	if callStop {
		// 调用webapi停止加速
		logger.Info("开始调用WebApi停止加速服务...")
		_ = leigod.Stop(User.UserName, User.PassWord)
	}

	// 关机
	logger.Info("关机...")
	shutdown.Shutdown()
}
