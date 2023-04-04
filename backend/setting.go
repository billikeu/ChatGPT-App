package backend

import (
	"log"

	"github.com/billikeu/Go-ChatBot/bot/params"
	"github.com/tidwall/gjson"
)

type AccountType string

const (
	ChatGPT           AccountType = "ChatGPT"
	ChatGPTUnofficial AccountType = "ChatGPTUnofficial"
	NewBingUnofficial AccountType = "NewBingUnofficial"
)

type AccountInfo struct {
	CurrentType       AccountType `json:"currentType"`
	OpenaiApiKey      string      `json:"openaiApiKey"`
	BaseURL           string      `json:"baseURL"`
	OpenaiAccessToken string      `json:"openaiAccessToken"`
	NewbingCookies    string      `json:"newbingCookies"`
	Proxy             string      `json:"proxy"`
}

type AccountState struct {
	AccountInfo AccountInfo `json:"accountInfo"`
}

type Setting struct {
	setting gjson.Result
}

func NewSetting(str string) *Setting {
	s := &Setting{
		setting: gjson.Parse(str),
	}
	return s
}

func (s *Setting) Proxy() string {
	return s.setting.Get("accountInfo.proxy").String()
}

// newbingCookies
func (s *Setting) ChatSk() string {
	var sk string
	if s.Engine() == params.ChatGPT {
		sk = s.setting.Get("accountInfo.openaiApiKey").String()
	} else if s.Engine() == params.NewBingUnofficial {
		sk = s.setting.Get("accountInfo.newbingCookies").String()
		// demo := []map[string]interface{}{}
		// err := json.Unmarshal([]byte(sk), &demo)
		// if err != nil {
		// 	log.Println(err)
		// }
		// log.Println(demo, len(demo))
	}
	return sk
}

// ChatGPT
// ChatGPT
func (s *Setting) Engine() string {
	engine := s.setting.Get("accountInfo.currentType").String()
	if engine != params.ChatGPT && engine != params.ChatGPTUnofficial && engine != params.NewBingUnofficial {
		log.Println("engine err:", engine)
		return ""
	}
	return engine
}
