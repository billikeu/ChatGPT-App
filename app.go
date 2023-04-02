package main

import (
	"context"
	"fmt"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) DownloadProcess(prompt string, onDownloadProgress func(map[string]interface{})) ([]byte, error) {

	// 循环读取响应体
	for i := 0; i < 10; i++ {
		// 更新下载进度
		if onDownloadProgress != nil {
			onDownloadProgress(map[string]interface{}{})
		}
		time.Sleep(time.Second * 3)
	}
	return []byte{}, nil
}
