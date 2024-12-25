package leigod

import (
	"bytes"
	"encoding/json"
	"github.com/XuHandsome/leigod-helper/libs"
	"github.com/XuHandsome/leigod-helper/pkgs/logger"
	"io"
	"net/http"
)

func Pause(accountToken string) (bool, error) {
	url := libs.WebApiHost + libs.WebApiPausePath

	data := map[string]interface{}{
		"account_token": accountToken,
		"lang":          libs.Lang,
	}

	requestBody, err := json.Marshal(data)
	if err != nil {
		logger.Error("请求Body解析Json失败: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		logger.Error("Pause Post请求报错: %v", err)
		return false, err
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

	var response libs.PauseResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		logger.Error("解析 JSON 错误: %v", err)
		return false, err
	}

	logger.Info("暂停状态码: %v", response.Code)
	logger.Info("暂停状态: %v", response.Msg)

	return true, nil
}
