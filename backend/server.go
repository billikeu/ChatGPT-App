package backend

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/billikeu/ChatGPT-App/backend/middlewares"

	"github.com/billikeu/Go-ChatBot/bot"
	bingunofficial "github.com/billikeu/Go-ChatBot/bot/bingUnofficial"
	"github.com/billikeu/Go-ChatBot/bot/params"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/tidwall/gjson"
)

type Server struct {
	ctx     context.Context
	bot     *bot.Bot
	setting *Setting
}

func NewServer(sk, proxy string) *Server {
	s := &Server{
		bot: bot.NewBot(&bot.Config{
			Proxy: proxy, // socks5://10.0.0.13:3126 , http://10.0.0.13:3127
			ChatGPT: bot.ChatGPTConf{
				SecretKey: sk, // your secret key
			},
			BingUnofficialConfig: &bingunofficial.BingConfig{
				Cookies: []map[string]interface{}{map[string]interface{}{
					"name":  "demo",
					"value": "demo",
				}},
			},
		}),
	}
	return s
}

func (s *Server) SetSetting(jsonStr string) {
	s.setting = NewSetting(jsonStr)
}

func (s *Server) Init(ctx context.Context) {
	s.ctx = ctx
	go s.Start()
}

func (s *Server) Start() {
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/chat-process", s.chatProcess)
	r.POST("/config", s.config)
	r.POST("/session", s.session)
	r.POST("/verify", s.verify)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// /chat-process
/*
req: {"prompt":"2","options":{"conversationId":"2d1fe5e1-fedc-4e2b-bf75-065b1da01abc","parentMessageId":"chatcmpl-70LeQVv1CupSiy6aeio4rt5safFYA"},"systemMessage":"You are ChatGPT, "}
err: {"message":"","data":null,"status":"Fail"}
*/
func (s *Server) chatProcess(c *gin.Context) {
	c.Header("Content-type", "application/octet-stream")

	// c.Request.b
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
	// log.Println(j)
	var chunkIndex int
	var text string
	err = s.bot.Ask(context.Background(), &params.AskParams{
		Proxy:             s.setting.Proxy(),
		SecretKey:         s.setting.ChatSk(),
		RefreshProxy:      true,
		RefreshSecretKey:  true,
		ConversationId:    conversationId,
		Prompt:            prompt,
		ChatEngine:        s.setting.Engine(), // params.ChatGPT
		SystemRoleMessage: j.Get("options.systemMessage").String(),
		Callback: func(_params *params.CallParams, err error) {
			if err != nil {
				return
			}
			chunkIndex += 1
			if _params.Chunk == "" {
				_params.Chunk = strings.TrimPrefix(_params.Text, text)
			}
			text = _params.Text
			m := ChatMessage{
				Role:            "assistant",
				ID:              _params.MsgId,
				ParentMessageID: _params.ParentId,
				ConversationID:  conversationId,
				Text:            _params.Text,
				Detail: DetailInfo{
					ID:      _params.MsgId,
					Object:  "chat.completion.chunk",
					Created: time.Now().Unix(),
					Model:   params.ChatGPT,
					Choices: []ChoiceInfo{
						{
							Delta: DeltaInfo{
								Content: _params.Chunk,
							},
							Index:        0,
							FinishReason: _params.Done,
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
		},
	})
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
}

// /config
func (s *Server) config(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": nil,
		"data": gin.H{
			"apiModel":     "ChatGPTAPI",
			"reverseProxy": "-",
			"timeoutMs":    60000,
			"socksProxy":   "-",
			"httpsProxy":   "-",
			"balance":      "-",
		},
		"status": "Success",
	})
}

// /session
func (s *Server) session(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "",
		"data": gin.H{
			"auth":  false,
			"model": "ChatGPTAPI",
		}})
}

// /verify
func (s *Server) verify(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
