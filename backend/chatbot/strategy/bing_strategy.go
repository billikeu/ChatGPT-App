package strategy

import (
	"github.com/billikeu/ChatGPT-App/backend/chatbot/factory"
	"github.com/billikeu/ChatGPT-App/backend/chatbot/message"
)

// BingStrategy
type BingStrategy struct {
	factory *factory.BingFactory
}

// NewBingStrategy
func NewBingStrategy(factory *factory.BingFactory) *BingStrategy {
	return &BingStrategy{
		factory: factory,
	}
}

// SendMessage
func (s *BingStrategy) SendMessage(req interface{}) (*message.Responser, error) {
	return s.factory.CreateChatbot().HandleMessage(req)
}

func (s *BingStrategy) GetName() string {
	return StrategyNameBing
}
