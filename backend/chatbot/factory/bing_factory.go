package factory

import "github.com/billikeu/ChatGPT-App/backend/chatbot/engine"

type BingFactory struct{}

func (f *BingFactory) CreateChatbot() Chatbot {
	// create a new chatbot instance that uses the Bing engine
	chatbot := NewChatBot()
	chatbot.SetChatEngine(engine.NewBing())

	return chatbot
}
