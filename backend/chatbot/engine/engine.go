package engine

import "github.com/billikeu/ChatGPT-App/backend/chatbot/message"

const (
	EngineNameChatGPT    = "ChatGPT"
	EngineNameBing       = "Bing"
	EngineNameBard       = "Bard"
	EngineNameChatGPTUno = "ChatGPTUno"
)

type ChatEngine interface {
	SendMessage(req interface{}) (response *message.Responser, err error)
	Close() error
}
