package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

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
	OrderId       string `json:"orderId"`
	QuickAccessId string `json:"-"`
	Type          string `json:"type"`
	Title         string `json:"title"`
	Description   string `json:"description"`
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

// Scan scan value into Jsonb, implements sql.Scanner interface
func (a *Action) Scan(value interface{}) error {
	source, ok := value.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i Action
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}
	*a = i
	return nil
}

// Value return json value, implement driver.Valuer interface
func (a Action) Value() (driver.Value, error) {
	return json.Marshal(a)
}
