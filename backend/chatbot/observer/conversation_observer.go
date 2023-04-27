package observer

import (
	"log"
)

// SessionObserver
type ConversationObserver interface {
	Update(data interface{})
}

// ChatbotSessionObserver
type ChatbotConversationObserver struct{}

// NewChatbotSessionObserver creates a new session observer
func NewChatbotConversationObserver() *ChatbotConversationObserver {
	return &ChatbotConversationObserver{}
}

// Update, Implement the observer interface, called when the user session is updated
func (c *ChatbotConversationObserver) Update(data interface{}) {
	// Processing logic can be added here, such as writing chat records to the database, etc.
	log.Printf("ChatbotConversationObserver: %v \n", data)
}

// UserSessionObserver
type UserSessionObserver struct{}

// NewSessionObserver creates a new session observer
func NewUserSessionObserver() *UserSessionObserver {
	return &UserSessionObserver{}
}

// Update , called when the user session is updated
func (u *UserSessionObserver) Update(data interface{}) {
	log.Printf("UserSessionObserver: session %v updated\n", data)
}
