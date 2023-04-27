package message

// type Requester interface {
// }

type Requester struct {
	Prompt       string
	ConversionID string // used by chatgptuno,
	ParentId     string // used by chatgptuno,
	Proxy        string
	Timeout      int // second
}

type ChatGPTUnoRequester struct {
	Requester
	ModelName   string // used by chatgptuno,
	AccessToken string // used by chatgptuno, https://chat.openai.com/api/auth/session
}

type ChatGPTRequester struct {
	Requester
	ModelName string
	SecretKey string
	SystemMsg string
}

type BingRequester struct {
	Requester
	ModelName string
	StrCookie string
	Cookies   []map[string]interface{}
}
