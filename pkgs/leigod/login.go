package leigod

import (
	"bytes"
	"encoding/json"
	"github.com/XuHandsome/leigod-helper/libs"
	"github.com/XuHandsome/leigod-helper/pkgs/logger"
	"io"
	"net/http"
)

// Login 登录以获取账户信息，包含AccountToken
func Login(data map[string]interface{}) (libs.LoginResponse, error) {
	url := libs.WebApiHost + libs.WebApiLoginPath

	// 存储解析后的JSON响应
	var response libs.LoginResponse

	requestBody, err := json.Marshal(data)
	if err != nil {
		logger.Error("请求Body解析Json失败: %v", err)
	}

	//fmt.Printf("%q", requestBody)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		logger.Error("Login Post请求报错: %v", err)
		return response, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("读取返回失败: %v", err)
	}

	//fmt.Printf("\n响应内容: %s", string(body))
	// 将响应内容解析为 map
	err = json.Unmarshal(body, &response)
	if err != nil {
		logger.Error("解析 JSON 错误: %v", err)
		return response, err
	}

	// 打印登录结果
	logger.Info("登录状态: %v", response.Msg)
	logger.Debug("获取到AccountToken: %v", response.Data.LoginInfo.AccountToken)
	return response, nil
}
