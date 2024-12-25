package config

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/XuHandsome/leigod-helper/libs"
	"github.com/XuHandsome/leigod-helper/pkgs/leigod"
	"github.com/XuHandsome/leigod-helper/pkgs/logger"
	"golang.org/x/term"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"path/filepath"
)

type UserInfo struct {
	UserName string `yaml:"userName"`
	PassWord string `yaml:"passWord"`
}

func InitConf() (UserInfo, bool, error) {
	var User UserInfo
	callStop := true
	var username string
	var password string

	// 准备存放目录 = $HOME/.leigod/$LogFilePath
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logger.Error("error retrieving home directory: %v", err)
		return User, false, err
	}

	configPath := homeDir + "/.leigod/" + libs.ConfFile
	logDir := filepath.Dir(configPath)
	err = os.MkdirAll(logDir, 0755)
	if err != nil {
		logger.Error("failed to create config directory: %v", err)
		return User, false, err
	}

	// 检查配置文件是否存在
	_, err = os.Stat(configPath)
	if err != nil {
		// 文件不存在, 引导输入，并生成配置文件
		if os.IsNotExist(err) {
			logger.Info("配置文件不存在，首次运行请输入雷神加速器用户信息")
			fmt.Print("请输入用户名|手机号: ")
			// 读取用户输入并赋值给 username
			_, err = fmt.Scanln(&username)
			if err != nil {
				logger.Error("读取输入信息发生错误: %v", err)
				return User, false, err
			}

			fmt.Print("请输入密码: ")
			passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
			fmt.Printf("\n")
			if err != nil {
				log.Fatal("读取输入密码时发生错误:", err)
			}
			password = string(passwordBytes)
			ret := md5.Sum([]byte(password))
			password = hex.EncodeToString(ret[:])

			// 验证账号密码
			err = leigod.Stop(username, password)
			if err != nil {
				logger.Error("输入账号密码有误，退出...")
				return User, false, err
			} else {
				// 创建文件
				configFile, err := os.Create(configPath)
				if err != nil {
					logger.Error("创建配置文件发生错误: %v", err)
				}
				defer func(configFile *os.File) {
					err := configFile.Close()
					if err != nil {

					}
				}(configFile)

				// 创建编码器
				encoder := yaml.NewEncoder(configFile)

				// 将配置编码为YAML数据
				User.UserName = username
				User.PassWord = password

				err = encoder.Encode(&User)
				if err != nil {
					logger.Error("写入配置文件失败: %v", err)
				}
				callStop = false
			}
		} else {
			logger.Error("检查文件是否存在过程报错:", err)
			return User, false, err
		}
	} else {
		// 配置文件已存在, 获取用户信息
		logger.Info("配置文件已存在")

		// 读取配置文件
		file, err := os.Open(configPath)
		if err != nil {
			logger.Error("打开配置文件发生错误: %v", err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		configData, err := io.ReadAll(file)
		if err != nil {
			logger.Error("读取配置文件发生错误: %v", err)
		}

		err = yaml.Unmarshal(configData, &User)
	}

	return User, callStop, nil
}
