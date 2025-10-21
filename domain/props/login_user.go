package props

type LoginUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
