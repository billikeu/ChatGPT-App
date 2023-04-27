package strategy

import (
	"github.com/billikeu/ChatGPT-App/backend/chatbot/factory"
	"github.com/billikeu/ChatGPT-App/backend/chatbot/message"
)

// ChatGPTStrategy
type ChatGPTUnoStrategy struct {
	factory factory.ChatbotFactory
}

// NewChatGPTStrategy 
func NewChatGPTUnoStrategy(factory factory.ChatbotFactory) *ChatGPTUnoStrategy {
	return &ChatGPTUnoStrategy{
		factory: factory,
	}
}

// SendMessage 
func (s *ChatGPTUnoStrategy) SendMessage(req interface{}) (*message.Responser, error) {
	return s.factory.CreateChatbot().HandleMessage(req)
}

func (s *ChatGPTUnoStrategy) GetName() string {
	return StrategyNameChatGPTUno
}
