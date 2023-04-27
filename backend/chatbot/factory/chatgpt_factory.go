package factory

import "github.com/billikeu/ChatGPT-App/backend/chatbot/engine"

type ChatGPTFactory struct{}

func (f *ChatGPTFactory) CreateChatbot() Chatbot {
	// create a new chatbot instance that uses the ChatGPT engine
	chatbot := NewChatBot()
	chatbot.SetChatEngine(engine.NewChatGPT())

	return chatbot
}
