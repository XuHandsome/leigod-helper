package leigod

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/XuHandsome/leigod-helper/libs"
	"github.com/XuHandsome/leigod-helper/pkgs/logger"
	"sort"
	"time"
)

// Signature 生成请求的签名
func Signature(data map[string]interface{}) map[string]interface{} {
	// 添加时间戳到请求体
	ts := time.Now().Unix()
	data["ts"] = ts

	// 排序参数
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 构建查询字符串
	queryString := ""
	suffix := "&"
	for i, k := range keys {

		if i == len(keys)-1 {
			suffix = ""
		}

		var value string
		switch v := data[k].(type) {
		case string:
			value = v
		case int:
			value = fmt.Sprintf("%d", v)
		default:
			value = fmt.Sprintf("%v", v) // 如果是其他类型，转换为字符串
		}

		queryString += k + "=" + value + suffix
	}

	// 添加密钥
	queryString += "&key=" + libs.SignSecret

	logger.Debug("Signal: %v", queryString)

	// 计算MD5签名
	hash := md5.New()
	hash.Write([]byte(queryString))
	signature := hex.EncodeToString(hash.Sum(nil))

	data["sign"] = signature

	return data
}
