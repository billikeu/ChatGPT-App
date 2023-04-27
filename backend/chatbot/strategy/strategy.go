package strategy

import "github.com/billikeu/ChatGPT-App/backend/chatbot/message"

const (
	StrategyNameChatGPT    = "ChatGPTStrategy"
	StrategyNameBard       = "BardStrategy"
	StrategyNameBing       = "BingStrategy"
	StrategyNameChatGPTUno = "ChatGPTUnoStrategy"
)

type Strategy interface {
	GetName() string
	SendMessage(req interface{}) (*message.Responser, error)
}
