package engine

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/billikeu/ChatGPT-App/backend/chatbot/message"
	"github.com/billikeu/Go-EdgeGPT/edgegpt"
)

var bingEngine *Bing = &Bing{
	bots: make(map[string]*edgegpt.ChatBot, 0),
}

type Bing struct {
	bots map[string]*edgegpt.ChatBot
}

func NewBing() *Bing {
	return bingEngine
}

func (b *Bing) SendMessage(req interface{}) (response *message.Responser, err error) {
	reqData, ok := req.(*message.BingRequester)
	if !ok {
		err = errors.New("mismatched request data type, the request data type must be message.BingRequester")
		return nil, err
	}
	bot := b.bots[reqData.ConversionID]
	if bot == nil {
		bot = edgegpt.NewChatBot("", reqData.Cookies, reqData.Proxy)
		err = bot.Init()
		if err != nil {
			return nil, fmt.Errorf("init edge bot err %s", err.Error())
		}
		b.bots[reqData.ConversionID] = bot
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

	go bot.Ask(reqData.Prompt, edgegpt.Creative, func(answer *edgegpt.Answer) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("callback panic:", r)
			}
		}()
		response.Msg <- &message.Response{
			Body: answer,
			Err:  nil,
		}
		if answer.IsDone() {
			close(response.Msg)
			done <- struct{}{}
		}
	})
	return
}

func (b *Bing) Close() error {
	return nil
}
