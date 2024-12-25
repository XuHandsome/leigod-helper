package libs

// 全局配置
const (
	// LogFile 日志位置，最终执行 = $HOME/.leigod/$LogFilePath
	LogFile = "logs/helper.log"

	// ConfFile 配置文件位置，最终执行 = $HOME/.leigod/$ConfFilePath
	ConfFile = "helper.yml"
)

// 雷神webapi
const (
	WebApiHost      = "https://webapi.leigod.com"
	WebApiLoginPath = "/api/auth/login/v1"
	WebApiPausePath = "/api/user/pause"
)

// SignSecret 签名加密密钥
const SignSecret = "5C5A639C20665313622F51E93E3F2783"

// Login请求体常量
const (
	CountryCode = 86
	Lang        = "zh_CN"
	OsType      = 4
	RegionCode  = 1
	UserType    = "0"
	SrcChannel  = "guanwang"
)
