package model

type Action struct {
	Id            string `json:"id"`
	OrderId       string `json:"orderId"`
	ActionCommand string `json:"actionCommand"`
}
type QuickAccess struct {
	Id      string `json:"id"` //uuid
	UserId  string `json:"userId"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	Picture string `json:"picture"`
	Action  `json:"action"`
}
