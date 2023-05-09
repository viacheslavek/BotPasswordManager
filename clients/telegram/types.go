package telegram

import "net/http"

type Client struct {
	host     string
	basePath string
	client   http.Client
}

type Update struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}
