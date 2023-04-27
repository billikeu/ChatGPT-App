package factory

import "github.com/billikeu/ChatGPT-App/backend/chatbot/engine"

type ChatGPTUnoFactory struct{}

func (f *ChatGPTUnoFactory) CreateChatbot() Chatbot {
	// create a new chatbot instance that uses the ChatGPT engine
	chatbot := NewChatBot()
	chatbot.SetChatEngine(engine.NewChatGPTUno())

	return chatbot
}
