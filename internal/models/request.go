package models

type RequestBody struct {
	Message string `json:"message"`
}

type BotResponseString struct {
	UserInput   string `json:"userInput"`
	BotResponse string `json:"botResponse"`
}
