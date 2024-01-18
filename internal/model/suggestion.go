package model

type Suggestion struct {
	Id          string `json:"id"` //uuid
	Type        string `json:"type"`
	Title       string `json:"title"`
	Picture     string `json:"picture"`
	DeepLink    string `json:"deepLink"`
	Description string `json:"description"`
}
