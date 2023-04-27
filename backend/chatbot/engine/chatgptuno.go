package engine

import (
	"errors"
	"log"
	"time"

	"github.com/billikeu/ChatGPT-App/backend/chatbot/message"

	"github.com/billikeu/go-chatgpt/chatgptuno"
)

var chatgptUnoEngine *ChatGPTUno = &ChatGPTUno{
	chat: nil,
}

type ChatGPTUno struct {
	chat *chatgptuno.ChatGPTUnoBot
}

func NewChatGPTUno() *ChatGPTUno {
	return chatgptUnoEngine
}

func (cg *ChatGPTUno) SendMessage(req interface{}) (response *message.Responser, err error) {
	reqData, ok := req.(*message.ChatGPTUnoRequester)
	if !ok {
		err = errors.New("mismatched request data type, the request data type must be message.ChatGPTUnoRequester")
		return nil, err
	}
	if cg.chat == nil {
		cg.chat = chatgptuno.NewChatGPTUnoBot(&chatgptuno.ChatGPTUnoConfig{
			Model:       reqData.ModelName,
			AccessToken: reqData.AccessToken,
			BaseUrl:     "https://chat-proxy.omba.cc/backend-api/",
		})
	}
	response = &message.Responser{
		Msg: make(chan *message.Response, 9999),
	}
	done := make(chan struct{})
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		select {
		case <-time.After(time.Second * time.Duration(reqData.Timeout)):
			close(response.Msg)
			break
		case <-done:
			break
		}
	}()
	go cg.chat.Ask(reqData.Prompt, reqData.ConversionID, reqData.ParentId, reqData.ModelName, reqData.Timeout, func(chatRes *chatgptuno.Response, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("callback panic:", r)
			}
		}()
		response.Msg <- &message.Response{
			Body: chatRes,
			Err:  err,
		}
		if err != nil {
			return
		}
		if chatRes.Message.EndTurn && chatRes.Message.Author.Role == "assistant" {
			close(response.Msg)
			done <- struct{}{}
		}
	})
	return
}

func (cg *ChatGPTUno) Close() error {
	return nil
}
