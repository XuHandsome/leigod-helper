package shutdown

import (
	"github.com/XuHandsome/leigod-helper/pkgs/logger"
	"os/exec"
	"runtime"
)

func Shutdown() {

	if runtime.GOOS == "windows" {
		logger.Info("检测到Windows平台，开始执行关机...")
		if err := exec.Command("cmd", "/C", "shutdown", "/p").Run(); err != nil {
			logger.Error("Failed to initiate shutdown: %v", err)
		}
	} else {
		logger.Info("非Windows平台，跳过关机步骤")
	}
}
