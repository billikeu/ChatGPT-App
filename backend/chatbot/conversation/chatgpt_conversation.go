package conversation

import (
	"github.com/billikeu/ChatGPT-App/backend/chatbot/message"
	"github.com/billikeu/ChatGPT-App/backend/chatbot/observer"

	uuid "github.com/satori/go.uuid"
)

type ChatConversation struct {
	id             string
	ConversationId string
	ParentId       string
	observers      []observer.Observer
	data           []string
	strategyName   string
	SecretKey      string
	SystemMsg      string
	Prompt         string
	Proxy          string
}

func NewChatConversation(conversationID, secretKey string) *ChatConversation {
	s := &ChatConversation{
		id:             uuid.NewV4().String(),
		ConversationId: conversationID,
		observers:      make([]observer.Observer, 0),
		data:           make([]string, 0),
		SecretKey:      secretKey,
	}
	return s
}
func (s *ChatConversation) ID() string {
	return s.id
}

func (s *ChatConversation) ConversionID() string {
	return s.ConversationId
}

func (s *ChatConversation) GetMessage() interface{} {
	// notify all observers
	return &message.ChatGPTRequester{
		Requester: message.Requester{
			ConversionID: s.ConversationId,
			ParentId:     s.ParentId,
			Prompt:       s.Prompt,
			Proxy:        s.Proxy,
			Timeout:      120,
		},
		SecretKey: s.SecretKey,
	}
}

func (s *ChatConversation) StrategyName() string {
	return s.strategyName
}

func (s *ChatConversation) SetStrategy(strategyName string) {
	s.strategyName = strategyName
}

func (s *ChatConversation) AddMessage(message string) {
	// notify all observers
	s.Prompt = message
	s.NotifyObservers(message)
}

func (s *ChatConversation) SetLastMessage(msg string) {

}

func (s *ChatConversation) RegisterObserver(observer observer.Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ChatConversation) RemoveObserver(observer observer.Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ChatConversation) NotifyObservers(data interface{}) {
	for _, observer := range s.observers {
		go observer.Update(data)
	}
}
