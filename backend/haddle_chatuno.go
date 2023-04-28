package backend

import (
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/billikeu/ChatGPT-App/backend/chatbot/conversation"
	"github.com/billikeu/ChatGPT-App/backend/chatbot/engine"
	"github.com/billikeu/ChatGPT-App/backend/chatbot/strategy"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// /chat-process
/*
req: {"prompt":"2","options":{"conversationId":"2d1fe5e1-fedc-4e2b-bf75-065b1da01abc","parentMessageId":"chatcmpl-70LeQVv1CupSiy6aeio4rt5safFYA"},"systemMessage":"You are ChatGPT, "}
err: {"message":"","data":null,"status":"Fail"}
*/
func (s *Server) chatProcessChatGPTUno(c *gin.Context) {
	c.Header("Content-type", "application/octet-stream")
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// log.Println(string(bodyBytes))
	defer c.Request.Body.Close()
	j := gjson.ParseBytes(bodyBytes)
	prompt := j.Get("prompt").String()
	conversationId := j.Get("options.conversationId").String()
	parentMessageId := j.Get("options.parentMessageId").String()
	// log.Println(j)
	var chunkIndex int
	var text string

	// chatbotSessionObserver := observer.NewChatbotConversationObserver()

	// create sessions and register observers
	// 6361952b-e0ba-4222-a40e-f16713d2105b
	var conv1 *conversation.ChatGPTUnoConversation
	conv := s.botV2.GetConversation(conversationId)
	if conv == nil {
		conv1 = conversation.NewChatGPTUnoConversation(conversationId, parentMessageId, s.setting.accountState.AccountInfo.OpenaiAccessToken)
		// conv1.RegisterObserver(chatbotSessionObserver)
		s.botV2.AddConversation(conv1)
		conv1.SetStrategy(strategy.StrategyNameChatGPTUno)
	} else {
		conv1 = conv.(*conversation.ChatGPTUnoConversation)
	}

	// simulate chat between user1 and chatbot
	conv1.AddMessage(prompt)
	response, err := s.botV2.GetResponse(conv1)
	if err != nil {
		log.Printf("get response err:%s", err.Error())
		return
	}
	// var text string
	for {
		chunkIndex += 1
		rmsg := <-response.Msg
		if rmsg == nil {
			break
		}
		if rmsg.Err != nil {
			log.Println(rmsg.Err)
			if err != nil {
				m := FailMessage{
					Message: err.Error(),
					Status:  "fail",
					Data:    nil,
				}
				msg := m.String()
				if chunkIndex > 1 {
					msg = "\n" + msg
				}
				c.Writer.Write([]byte(msg))
				c.Writer.Flush()
			}
			break
		}
		data, _ := rmsg.ChatGPTUno()
		if data == nil {
			continue
		}

		// log.Println(data.Raw)
		chunk := strings.TrimPrefix(data.Message.Content.Parts[0], text)
		text = data.Message.Content.Parts[0]
		m := ChatMessage{
			Role:            "assistant",
			ID:              data.Message.ID,
			ParentMessageID: conv1.ParentId,
			ConversationID:  data.ConversationID,
			Text:            data.Message.Content.Parts[0],
			Detail: DetailInfo{
				ID:      data.Message.ID,
				Object:  "chat.completion.chunk",
				Created: time.Now().Unix(),
				Model:   engine.EngineNameChatGPTUno,
				Choices: []ChoiceInfo{
					{
						Delta: DeltaInfo{
							Content: chunk,
						},
						Index:        0,
						FinishReason: data.Message.EndTurn,
					},
				},
			},
		}
		msg := m.String()
		if chunkIndex > 1 {
			msg = "\n" + msg
		}
		c.Writer.Write([]byte(msg))
		c.Writer.Flush()

		if conv1.ConversationId != data.ConversationID && data.Message.Author.Role == "assistant" {
			conv1.ConversationId = data.ConversationID
		}
		if conv1.ParentId != data.Message.ID && data.Message.Author.Role == "assistant" {
			conv1.ParentId = data.Message.ID
		}
	}
}
