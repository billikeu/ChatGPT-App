package strategy

import (
	"github.com/billikeu/ChatGPT-App/backend/chatbot/factory"
	"github.com/billikeu/ChatGPT-App/backend/chatbot/message"
)

// BardStrategy
type BardStrategy struct {
	factory *factory.BardFactory
}

// NewChatGPTStrategy
func NewBardStrategy(factory *factory.BardFactory) *BardStrategy {
	return &BardStrategy{
		factory: factory,
	}
}

// SendMessage
func (s *BardStrategy) SendMessage(req interface{}) (*message.Responser, error) {
	return s.factory.CreateChatbot().HandleMessage(req)
}

func (s *BardStrategy) GetName() string {
	return StrategyNameBard
}
