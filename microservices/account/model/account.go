package model

type Account struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ServedBy string `json:"servedBy"`
	Quote    Quote  `json:"quote"`
}

type Quote struct {
	Text     string `json:"quote"`
	ServedBy string `json:"ipAddress"`
	Language string `json:"language"`
}

type VipNotification struct {
	AccountId string `json:"account_id"`
	ReadAt string `json:"read_at"`
}