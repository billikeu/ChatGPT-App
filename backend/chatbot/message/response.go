package message

import (
	"fmt"

	"github.com/billikeu/Go-EdgeGPT/edgegpt"
	"github.com/billikeu/go-chatgpt/chatgptuno"
	"github.com/billikeu/go-chatgpt/params"
)

type Response struct {
	Body interface{}
	Err  error
}

type Responser struct {
	Msg chan *Response
}

func (r *Response) ChatGPTUno() (*chatgptuno.Response, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	data, ok := r.Body.(*chatgptuno.Response)
	if !ok {
		return nil, fmt.Errorf("error parser type")
	}
	return data, nil
}

func (r *Response) Bing() (*edgegpt.Answer, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	data, ok := r.Body.(*edgegpt.Answer)
	if !ok {
		return nil, fmt.Errorf("error parser type")
	}
	return data, nil
}

func (r *Response) ChatGPT() (*params.Answer, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	data, ok := r.Body.(*params.Answer)
	if !ok {
		return nil, fmt.Errorf("error parser type")
	}
	return data, nil
}
