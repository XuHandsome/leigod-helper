package libs

type LoginResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		LoginInfo struct {
			AccountToken string `json:"account_token"`
			ExpiryTime   string `json:"expiry_time"`
			NnToken      string `json:"nn_token"`
		} `json:"login_info"`
		UserInfo struct {
			NickName   string `json:"nickname"`
			Email      string `json:"email"`
			Mobile     string `json:"mobile"`
			Avatar     string `json:"avatar"`
			RegionCode int    `json:"region_code"`
		} `json:"user_info"`
	}
}

type PauseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
