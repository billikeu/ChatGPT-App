package conversation

import (
	"github.com/billikeu/ChatGPT-App/backend/chatbot/message"
	"github.com/billikeu/ChatGPT-App/backend/chatbot/observer"

	uuid "github.com/satori/go.uuid"
)

type ChatGPTUnoConversation struct {
	id             string
	ConversationId string
	ParentId       string // // parentId == res.Message.ID
	observers      []observer.Observer
	strategyName   string
	AccessToken    string
	Prompt         string
}

func NewChatGPTUnoConversation(conversationId, parentId, accessToken string) *ChatGPTUnoConversation {
	s := &ChatGPTUnoConversation{
		id:             uuid.NewV4().String(),
		ConversationId: conversationId,
		ParentId:       parentId,
		observers:      make([]observer.Observer, 0),
		AccessToken:    accessToken,
	}
	return s
}

func (s *ChatGPTUnoConversation) ID() string {
	return s.id
}

func (s *ChatGPTUnoConversation) ConversionID() string {
	return s.ConversationId
}

func (s *ChatGPTUnoConversation) GetMessage() interface{} {
	req := &message.ChatGPTUnoRequester{
		Requester: message.Requester{
			ConversionID: s.ConversationId,
			ParentId:     s.ParentId,
			Prompt:       s.Prompt,
			Proxy:        "",
			Timeout:      60,
		},
		ModelName:   "text-davinci-002-render-sha",
		AccessToken: s.AccessToken,
	}
	return req
}

func (s *ChatGPTUnoConversation) StrategyName() string {
	return s.strategyName
}

func (s *ChatGPTUnoConversation) SetStrategy(strategyName string) {
	s.strategyName = strategyName
}

func (s *ChatGPTUnoConversation) AddMessage(message string) {
	// notify all observers
	s.Prompt = message
	s.NotifyObservers(message)
}

func (s *ChatGPTUnoConversation) SetLastMessage(msg string) {

}

func (s *ChatGPTUnoConversation) RegisterObserver(observer observer.Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ChatGPTUnoConversation) RemoveObserver(observer observer.Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ChatGPTUnoConversation) NotifyObservers(data interface{}) {
	for _, observer := range s.observers {
		go observer.Update(data)
	}
}
