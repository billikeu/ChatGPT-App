package engine

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/billikeu/ChatGPT-App/backend/chatbot/message"
	"github.com/billikeu/go-chatgpt/chatgpt"
	"github.com/billikeu/go-chatgpt/params"
)

var chatGPTEngine *ChatGPT = &ChatGPT{
	bots: make(map[string]*chatgpt.ChatGPTConversion, 0),
}

type ChatGPT struct {
	bots map[string]*chatgpt.ChatGPTConversion
}

func NewChatGPT() *ChatGPT {
	return chatGPTEngine
}

func (cg *ChatGPT) SendMessage(req interface{}) (response *message.Responser, err error) {
	reqData, ok := req.(*message.ChatGPTRequester)
	if !ok {
		err = errors.New("mismatched request data type, the request data type must be message.ChatGPTRequester")
		return nil, err
	}

	chat := cg.bots[reqData.ConversionID]
	if chat == nil {
		chat = chatgpt.NewChatGPTConversion(reqData.SecretKey)
		cg.bots[reqData.ConversionID] = chat
	}
	err = chat.SetProxy(reqData.Proxy)
	if err != nil {
		return nil, err
	}
	chat.SetSystemMsg(reqData.SystemMsg)

	err = chat.Init()
	if err != nil {
		return nil, err
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
	go chat.Ask(context.Background(), reqData.Prompt, func(answer *params.Answer, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("callback panic:", r)
			}
		}()
		response.Msg <- &message.Response{
			Body: answer,
			Err:  err,
		}
		if err != nil {
			return
		}
		if answer.Done {
			close(response.Msg)
			done <- struct{}{}
		}
	})
	return
}

func (cg *ChatGPT) Close() error {
	return nil
}
