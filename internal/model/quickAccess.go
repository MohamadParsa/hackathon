package model

type Action struct {
	Id            string `json:"id"`
	OrderId       string `json:"orderId"`
	ActionCommand string `json:"actionCommand"`
}
type QuickAccess struct {
	Id      string  `json:"id"` //uuid
	UserId  string  `json:"-"`
	Type    string  `json:"type"`
	Title   string  `json:"title"`
	Picture string  `json:"picture"`
	Action  *Action `json:"action"`
}

type PurcahseHistory struct {
	OrderId     string `json:"orderId"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PurcahseHistoryList []*PurcahseHistory

func (purcahseHistoryList PurcahseHistoryList) FilerByType(serviceType string) PurcahseHistoryList {
	list := PurcahseHistoryList{}
	for _, history := range purcahseHistoryList {
		if history.Type == serviceType {
			list = append(list, history)
		}
	}
	return list
}
