package http

type Response struct {
	Result string `json:"result"`
	Emoji  string `json:"emoji"`
	IsNew  bool   `json:"isNew"`
}
