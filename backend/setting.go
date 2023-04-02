package backend

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
