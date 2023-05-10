package telegram

import "net/http"

type Client struct {
	host     string
	basePath string
	client   http.Client
}

type Update struct {
	ID      int              `json:"id"`
	Message *IncomingMessage `json:"message"`
}

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type IncomingMessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}

type From struct {
	Username string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
}
