package backend

import (
	"log"

	"github.com/billikeu/Go-ChatBot/bot/params"
	"github.com/tidwall/gjson"
)

const (
	ChatGPT           = "ChatGPT"
	ChatGPTUnofficial = "ChatGPTUnofficial"
	NewBingUnofficial = "NewBingUnofficial"
)

type AccountInfo struct {
	ChatEngine        string `json:"chat_engine"`
	OpenaiApiKey      string `json:"openai_api_key"`
	BaseURL           string `json:"base_url"`
	OpenaiAccessToken string `json:"openai_access_token"`
	NewbingCookies    string `json:"newbing_cookies"`
	Proxy             string `json:"proxy"`
}

type AccountState struct {
	AccountInfo AccountInfo `json:"account_info"`
}

type Setting struct {
	setting      gjson.Result
	accountState *AccountState `json:"account_state"`
}

func NewSetting(str string) *Setting {
	s := &Setting{
		setting:      gjson.Parse(str),
		accountState: &AccountState{},
	}
	return s
}

func (s *Setting) Proxy() string {
	return s.accountState.AccountInfo.Proxy
}

// newbingCookies
func (s *Setting) ChatSk() string {
	var sk string
	if s.Engine() == params.ChatGPT {
		sk = s.accountState.AccountInfo.OpenaiApiKey
		// sk = s.setting.Get("accountInfo.openaiApiKey").String()
	} else if s.Engine() == params.NewBingUnofficial {
		sk = s.accountState.AccountInfo.NewbingCookies
		// sk = s.setting.Get("accountInfo.newbingCookies").String()
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
	engine := s.accountState.AccountInfo.ChatEngine
	if engine != params.ChatGPT && engine != params.ChatGPTUnofficial && engine != params.NewBingUnofficial {
		log.Println("engine err:", engine)
		return ""
	}
	return engine
}
