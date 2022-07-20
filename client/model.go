package client

type Common struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type LoginResp struct {
	Common
	Data LoginRespData `json:"data"`
}

type LoginRespData struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Score    int    `json:"score"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
