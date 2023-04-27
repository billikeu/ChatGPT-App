package conversation

import (
	"encoding/json"
	"log"

	"github.com/billikeu/ChatGPT-App/backend/chatbot/message"
	"github.com/billikeu/ChatGPT-App/backend/chatbot/observer"

	uuid "github.com/satori/go.uuid"
)

type BingConversation struct {
	id             string
	ConversationId string
	ParentId       string // // parentId == res.Message.ID
	observers      []observer.Observer
	strategyName   string
	StrCookie      string
	Cookies        []map[string]interface{}
	Prompt         string
	Proxy          string
}

func NewBingConversation(conversationId, parentId, cookie string) *BingConversation {

	s := &BingConversation{
		id:             uuid.NewV4().String(),
		ConversationId: conversationId,
		ParentId:       parentId,
		observers:      make([]observer.Observer, 0),
		StrCookie:      cookie,
		Cookies:        []map[string]interface{}{},
	}
	err := json.Unmarshal([]byte(cookie), &s.Cookies)
	if err != nil {
		log.Println("parse cookies err ", err)
	}
	return s
}

func (s *BingConversation) ID() string {
	return s.id
}

func (s *BingConversation) ConversionID() string {
	return s.ConversationId
}

func (s *BingConversation) GetMessage() interface{} {
	req := &message.BingRequester{
		Requester: message.Requester{
			ConversionID: s.ConversationId,
			ParentId:     s.ParentId,
			Prompt:       s.Prompt,
			Proxy:        s.Proxy,
			Timeout:      60,
		},
		ModelName: "text-davinci-002-render-sha",
		StrCookie: s.StrCookie,
		Cookies:   s.Cookies,
	}
	return req
}

func (s *BingConversation) StrategyName() string {
	return s.strategyName
}

func (s *BingConversation) SetStrategy(strategyName string) {
	s.strategyName = strategyName
}

func (s *BingConversation) AddMessage(message string) {
	// notify all observers
	s.Prompt = message
	s.NotifyObservers(message)
}

func (s *BingConversation) SetLastMessage(msg string) {

}

func (s *BingConversation) RegisterObserver(observer observer.Observer) {
	s.observers = append(s.observers, observer)
}

func (s *BingConversation) RemoveObserver(observer observer.Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *BingConversation) NotifyObservers(data interface{}) {
	for _, observer := range s.observers {
		go observer.Update(data)
	}
}
