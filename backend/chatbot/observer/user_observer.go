package observer

import (
	"log"
	"time"
)

// UserObserver is an observer for user input
type UserObserver struct{}

type ObserverSessionData struct {
	UserID        string
	BotName       string
	Message       string      // message from user
	Response      string      // Reply from the robot
	Timestamp     time.Time   // message timestamp
	ExtraMetadata interface{} // optional extra metadata
}

// NewUserObserver
func NewUserObserver() *UserObserver {
	return &UserObserver{}
}

// Update notifies the observer of an event
func (u *UserObserver) Update(data interface{}) {
	if session, ok := data.(*ObserverSessionData); ok {
		log.Printf("User session %s has a new message: %s\n", session.UserID, session.Message)
	}
}
