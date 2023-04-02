package chatgpt

import (
	"context"
	"errors"
	"fmt"
	"io"

	openai "github.com/sashabaranov/go-openai"
)

type ChatGPT struct {
	secretKey string
}

func NewChatGPT(secretKey string) *ChatGPT {
	chat := &ChatGPT{
		secretKey: secretKey,
	}
	return chat
}

func (chat *ChatGPT) Ask() {
	c := openai.NewClient(chat.secretKey)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 20,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Lorem ipsum",
			},
		},
		Stream: true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}
		// response.r
		fmt.Printf(response.Choices[0].Delta.Content)
		fmt.Sprintln(response.Object)
	}
}
