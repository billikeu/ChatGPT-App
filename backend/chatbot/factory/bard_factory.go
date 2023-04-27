package factory

import "github.com/billikeu/ChatGPT-App/backend/chatbot/engine"

type BardFactory struct{}

func (f *BardFactory) CreateChatbot() Chatbot {
	// create a new chatbot instance that uses the Bard engine
	chatbot := NewChatBot()
	chatbot.SetChatEngine(engine.NewBard())

	return chatbot
}
