package factory

import (
	"fmt"

	"github.com/billikeu/ChatGPT-App/backend/chatbot/message"

	"github.com/billikeu/ChatGPT-App/backend/chatbot/engine"
)

type ChatEngineFactory interface {
	CreateChatEngine() (engine.ChatEngine, error)
}

// type ChatEngine interface {
// 	SendMessage(msg string) (string, error)
// }

type ChatbotFactory interface {
	CreateChatbot() Chatbot
}

type Chatbot interface {
	HandleMessage(req interface{}) (*message.Responser, error)
}

type chatbot struct {
	engine engine.ChatEngine
}

func NewChatBot() *chatbot {
	return &chatbot{}
}

func (c *chatbot) SetChatEngine(engine engine.ChatEngine) {
	c.engine = engine
}

func (c *chatbot) HandleMessage(req interface{}) (*message.Responser, error) {
	return c.engine.SendMessage(req)
}

func CreateFactory(engineName string) (ChatbotFactory, error) {
	// create the appropriate factory based on the chosen chat engine
	switch engineName {
	case engine.EngineNameChatGPT:
		return &ChatGPTFactory{}, nil
	case engine.EngineNameBing:
		return &BingFactory{}, nil
	case engine.EngineNameBard:
		return &BardFactory{}, nil
	case engine.EngineNameChatGPTUno:
		return &ChatGPTUnoFactory{}, nil
	default:
		return nil, fmt.Errorf("invalid chat engine '%s'", engineName)
	}
}
