package bot

import (
	"fmt"

	"github.com/billikeu/ChatGPT-App/backend/chatbot/conversation"
	"github.com/billikeu/ChatGPT-App/backend/chatbot/message"
	"github.com/billikeu/ChatGPT-App/backend/chatbot/observer"
	"github.com/billikeu/ChatGPT-App/backend/chatbot/strategy"
)

type Bot struct {
	ID            string
	Name          string
	Conversations map[string]conversation.Conversation
	Observers     []observer.Observer
	strategies    map[string]strategy.Strategy
}

func NewChatbot(strategies []strategy.Strategy) *Bot {
	bot := &Bot{
		strategies:    make(map[string]strategy.Strategy, 0),
		Conversations: make(map[string]conversation.Conversation, 0),
		Observers:     make([]observer.Observer, 0),
	}
	for _, s := range strategies {
		bot.strategies[s.GetName()] = s
	}
	return bot
}

func (bot *Bot) ResisterStrategy(strategy strategy.Strategy) {
	bot.strategies[strategy.GetName()] = strategy
}

func (bot *Bot) getStrategy(strategyName string) strategy.Strategy {
	return bot.strategies[strategyName]
}

func (bot *Bot) GetResponse(conversation conversation.Conversation) (response *message.Responser, err error) {
	// ...
	currentStrategy := bot.getStrategy(conversation.StrategyName())
	if currentStrategy == nil {
		return response, fmt.Errorf("no strategy available for conversation %s %s", conversation.ID(), conversation.StrategyName())
	}
	req := conversation.GetMessage()
	response, err = currentStrategy.SendMessage(req)
	if err != nil {
		return response, fmt.Errorf("send message err %s", err.Error())
	}
	// conversation.SetLastMessage(response)
	bot.NotifyObservers(req, response)
	return
}

// AddConversation
func (bot *Bot) AddConversation(conversation conversation.Conversation) {
	bot.Conversations[conversation.ID()] = conversation
}

// RemoveConversation
func (bot *Bot) RemoveConversation(conversationID string) {
	if bot.Conversations[conversationID] == nil {
		return
	}
	delete(bot.Conversations, conversationID)
}

// GetConversation
func (bot *Bot) GetConversation(conversationID string) conversation.Conversation {
	return bot.Conversations[conversationID]
}

// RegisterObserver
func (bot *Bot) RegisterObserver(observer observer.Observer) {
	bot.Observers = append(bot.Observers, observer)
}

// RemoveObserver
func (bot *Bot) RemoveObserver(observer observer.Observer) {
	var index int = -1
	for i, obs := range bot.Observers {
		if obs == observer {
			index = i
			break
		}
	}
	if index != -1 {
		bot.Observers = append(bot.Observers[:index], bot.Observers[index+1:]...)
	}
}

// NotifyObservers
func (bot *Bot) NotifyObservers(msg interface{}, response *message.Responser) {
	for _, observer := range bot.Observers {
		// go observer.Update(bot, msg, response)
		go observer.Update(response)
	}
}
