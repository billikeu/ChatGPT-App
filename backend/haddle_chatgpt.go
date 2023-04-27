package backend

import (
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/billikeu/ChatGPT-App/backend/chatbot/conversation"
	"github.com/billikeu/ChatGPT-App/backend/chatbot/strategy"
	"github.com/billikeu/Go-ChatBot/bot/params"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/tidwall/gjson"
)

// /chat-process
/*
req: {"prompt":"2","options":{"conversationId":"2d1fe5e1-fedc-4e2b-bf75-065b1da01abc","parentMessageId":"chatcmpl-70LeQVv1CupSiy6aeio4rt5safFYA"},"systemMessage":"You are ChatGPT, "}
err: {"message":"","data":null,"status":"Fail"}
*/
func (s *Server) chatProcessChatGPT(c *gin.Context) {
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
	if conversationId == "" {
		conversationId = uuid.NewV4().String()
	}
	systemMessage := j.Get("options.systemMessage").String()
	// parentMessageId := j.Get("options.parentMessageId").String()
	// log.Println(j)
	var chunkIndex int
	var text string

	secretKey := s.setting.accountState.AccountInfo.OpenaiApiKey
	// 6361952b-e0ba-4222-a40e-f16713d2105b
	var conv1 *conversation.ChatConversation
	conv := s.botV2.GetConversation(conversationId)
	if conv == nil {
		conv1 = conversation.NewChatConversation(conversationId, secretKey)
		conv1.SetStrategy(strategy.StrategyNameChatGPT)
		conv1.SystemMsg = systemMessage
		s.botV2.AddConversation(conv1)

	} else {
		conv1 = conv.(*conversation.ChatConversation)
	}
	conv1.Proxy = s.setting.Proxy()
	conv1.SecretKey = secretKey

	// simulate chat between user1 and chatbot
	conv1.AddMessage(prompt)
	response, err := s.botV2.GetResponse(conv1)
	if err != nil {
		log.Printf("get response err:%s", err.Error())
		return
	}
	msgId := uuid.NewV4().String()
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
		data, _ := rmsg.ChatGPT()
		if data == nil {
			continue
		}

		chunk := strings.TrimPrefix(data.Text, text)
		text = data.Text

		m := ChatMessage{
			Role:            "assistant",
			ID:              msgId,
			ParentMessageID: "",
			ConversationID:  conversationId,
			Text:            text,
			Detail: DetailInfo{
				ID:      msgId,
				Object:  "chat.completion.chunk",
				Created: time.Now().Unix(),
				Model:   params.NewBingUnofficial,
				Choices: []ChoiceInfo{
					{
						Delta: DeltaInfo{
							Content: chunk,
						},
						Index:        0,
						FinishReason: data.Done,
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
	}
}
