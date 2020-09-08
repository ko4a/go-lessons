package main

type Config struct {
	ApiKey   string
	dbConfig *DbConfig
	LogUrl   string
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	Id int `json:"id"`
}

type UpdateResponse struct {
	Result []Update `json:"result"`
}
