package main

import (
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

type WindowService struct{}

// ShouldQuit 阻止程序退出
func ShouldQuit() bool {
	App.Hide()
	return false
}

// WindowEventsRegister 窗口事件注册，阻止窗口关闭
func WindowEventsRegister() {
	mw, has := App.Window.GetByName("main window")
	if has {
		mw.RegisterHook(events.Windows.WindowClosing, func(event *application.WindowEvent) {
			event.Cancel()
			mw.Hide()
		})
	}

	op, has := App.Window.GetByName("options window")
	if has {
		op.OnWindowEvent(events.Windows.WindowInactive, func(event *application.WindowEvent) {
			op.Hide()
		})
	}
}
