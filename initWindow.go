package main

import "github.com/wailsapp/wails/v3/pkg/application"

// 创建并初始化窗口
func initWindow() {
	App.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:            "塞壁纸",
		Name:             "main window",
		Width:            1000,
		Height:           618,
		BackgroundColour: application.NewRGBA(6, 7, 15, 0),
		Frameless:        true,
		URL:              "/",
		BackgroundType:   1,
	})

	App.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:            "右键菜单",
		Name:             "options window",
		Width:            150,
		Height:           127,
		Hidden:           true,
		AlwaysOnTop:      true,
		BackgroundType:   1,
		Frameless:        true,
		BackgroundColour: application.NewRGBA(6, 7, 15, 0),
		URL:              "/tip",
		DisableResize:    true,
		Windows: application.WindowsWindow{
			HiddenOnTaskbar:                   true,
			BackdropType:                      2,
			DisableFramelessWindowDecorations: true,
		},
	})

}
