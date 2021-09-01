package endpoints

type UserRequest struct {
	Uid int64 `json:"uid"`
}

type UserResponse struct {
	Result string `json:"result"`
	Addr   string `json:"addr"`
}
