package main

import (
	"embed"

	"log"

	"github.com/go-vgo/robotgo"
	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

var (
	App *application.App
)

func main() {

	app := application.New(application.Options{
		Name:        "塞壁纸",
		Description: "给我的桌面塞进去一些壁纸",
		Services: []application.Service{
			application.NewService(&GreetService{}),
			application.NewService(&ConfigService{}),
		},
		Assets: application.AssetOptions{
			Handler: application.BundledAssetFileServer(assets),
		},
		Windows:    application.WindowsOptions{},
		ShouldQuit: ShouldQuit,
	})
	App = app
	initConfig()           // 加载配置
	app.Show()

	initWindow()           // 创建并初始化窗口
	WindowEventsRegister() // 注册窗口事件

	systray := app.SystemTray.New()
	mw, has := App.Window.GetByName("main window")
	if has {
		systray.OnClick(func() {
			mw.Show().Focus()
		})
	}

	op, has := App.Window.GetByName("options window")
	if has {
		systray.OnRightClick(func() {
			mx, my := robotgo.Location() // 获取鼠标位置
			op.SetPosition(mx-op.Width()/2, my-op.Height()-5)
			op.Show().Focus()
		})
	}

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
