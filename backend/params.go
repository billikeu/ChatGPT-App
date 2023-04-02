package backend

import "encoding/json"

type ChatMessage struct {
	Role            string     `json:"role"`
	ID              string     `json:"id"`
	ParentMessageID string     `json:"parentMessageId"`
	ConversationID  string     `json:"conversationId"`
	Text            string     `json:"text"`
	Delta           string     `json:"delta"`
	Detail          DetailInfo `json:"detail"`
}

type DetailInfo struct {
	ID      string       `json:"id"`
	Object  string       `json:"object"`
	Created int64        `json:"created"`
	Model   string       `json:"model"`
	Choices []ChoiceInfo `json:"choices"`
}

type ChoiceInfo struct {
	Delta        DeltaInfo   `json:"delta"`
	Index        int         `json:"index"`
	FinishReason interface{} `json:"finish_reason"`
}

type DeltaInfo struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

func (msg *ChatMessage) String() string {
	// b, err := json.MarshalIndent(msg, "", "  ")
	b, err := json.Marshal(msg)
	if err != nil {
		return ""
	}
	return string(b)
}

func (msg *ChatMessage) Byte() []byte {
	// b, err := json.MarshalIndent(msg, "", "  ")
	b, err := json.Marshal(msg)
	if err != nil {
		return nil
	}
	return b
}

// Fail Message
// {"message":"tnames: Host: bypass.duti.tech. is not in the","data":null,"status":"Fail"}
type FailMessage struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  string      `json:"status"`
}

func (msg *FailMessage) String() string {
	// b, err := json.MarshalIndent(msg, "", "  ")
	b, err := json.Marshal(msg)
	if err != nil {
		return ""
	}
	return string(b)
}

func (msg *FailMessage) Byte() []byte {
	// b, err := json.MarshalIndent(msg, "", "  ")
	b, err := json.Marshal(msg)
	if err != nil {
		return nil
	}
	return b
}
