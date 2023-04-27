package backend

import (
	"context"
	"net/http"

	"github.com/billikeu/ChatGPT-App/backend/middlewares"

	"github.com/billikeu/ChatGPT-App/backend/chatbot/bot"
	"github.com/billikeu/ChatGPT-App/backend/chatbot/factory"
	"github.com/billikeu/ChatGPT-App/backend/chatbot/strategy"
	"github.com/gin-gonic/gin"
)

type Server struct {
	ctx     context.Context
	setting *Setting
	botV2   *bot.Bot
}

func NewServer(sk, proxy string) *Server {

	// create strategies for each engine
	gptStrategy := strategy.NewChatGPTStrategy(&factory.ChatGPTFactory{})
	bingStrategy := strategy.NewBingStrategy(&factory.BingFactory{})
	bardStrategy := strategy.NewBardStrategy(&factory.BardFactory{})
	gptunoStrategy := strategy.NewChatGPTUnoStrategy(&factory.ChatGPTUnoFactory{})

	// create a list of strategies
	strategies := []strategy.Strategy{
		gptStrategy,
		bingStrategy,
		bardStrategy,
		gptunoStrategy,
	}

	// create a chatbot
	botV2 := bot.NewChatbot(strategies)

	s := &Server{
		botV2:   botV2,
		setting: NewSetting(""),
	}
	return s
}

func (s *Server) SetAccountState(accountState AccountState) {
	s.setting.accountState = &accountState
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
	if s.setting.Engine() == ChatGPT {
		s.chatProcessChatGPT(c)
	} else if s.setting.Engine() == ChatGPTUnofficial {
		s.chatProcessChatGPTUno(c)
	} else if s.setting.Engine() == NewBingUnofficial {
		s.chatProcessBing(c)
	} else {
		c.JSON(403, gin.H{
			"msg": "unsupported model engine",
		})
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
