package leigod

import (
	"github.com/XuHandsome/leigod-helper/libs"
)

func Stop(username string, password string) error {

	data := map[string]interface{}{
		"country_code": libs.CountryCode,
		"lang":         libs.Lang,
		"mobile_num":   username,
		"os_type":      libs.OsType,
		"password":     password,
		"region_code":  libs.RegionCode,
		"user_type":    libs.UserType,
		"src_channel":  libs.SrcChannel,
		"username":     username,
	}

	// 添加签名
	signatureData := Signature(data)

	// 登录获取AccountToken
	accountSession, err := Login(signatureData)
	if err != nil {
		return err
	}
	accountToken := accountSession.Data.LoginInfo.AccountToken

	// 调用暂停接口
	_, err = Pause(accountToken)
	if err != nil {
		return err
	}

	return nil
}
