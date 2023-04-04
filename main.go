package main

import (
	"context"
	"embed"
	"log"

	"github.com/billikeu/ChatGPT-App/backend"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app := NewApp()
	server := backend.NewServer("", "")

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "ChatGPT-App",
		Width:     1024,
		Height:    768,
		Frameless: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			server.Init(ctx)
		},
		Bind: []interface{}{
			app,
			server,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
