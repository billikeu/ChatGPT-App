package engine

import "github.com/billikeu/ChatGPT-App/backend/chatbot/message"

type Bard struct{}

func NewBard() *Bard {
	return &Bard{}
}

func (b *Bard) SendMessage(req interface{}) (response *message.Responser, err error) {

	return
}

func (b *Bard) Close() error {

	return nil
}
