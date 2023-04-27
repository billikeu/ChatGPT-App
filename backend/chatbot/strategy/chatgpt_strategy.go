package strategy

import (
	"github.com/billikeu/ChatGPT-App/backend/chatbot/factory"
	"github.com/billikeu/ChatGPT-App/backend/chatbot/message"
)

// ChatGPTStrategy
type ChatGPTStrategy struct {
	factory factory.ChatbotFactory
}

// NewChatGPTStrategy
func NewChatGPTStrategy(factory factory.ChatbotFactory) *ChatGPTStrategy {
	return &ChatGPTStrategy{
		factory: factory,
	}
}

// SendMessage
func (s *ChatGPTStrategy) SendMessage(req interface{}) (*message.Responser, error) {
	return s.factory.CreateChatbot().HandleMessage(req)
}

func (s *ChatGPTStrategy) GetName() string {
	return StrategyNameChatGPT
}
