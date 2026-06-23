package main

import (
	"fmt"

	"github.com/wailsapp/wails/v3/pkg/application"
)

/**
* 窗口操作，有个窗口的打开，大小等等涉及到窗口控制的都在这里

**/

// WindowOption 一个窗口的配置，为啥要但写一个出来呢，因为自带的获取配置好像不全
type WindowOption struct {
	IsFullscreen bool
	CantClose    bool // 是否可以关闭,true是不能关闭 一般常用的窗口只隐藏不关闭，频繁地创建删除窗口占用资源
	application.WebviewWindowOptions
}

// Windows 所有的窗口都在这，为啥要写在这呢，因为我也不知道写在哪？
var Windows = make(map[string]*application.WebviewWindow)
var WindowOptions = make(map[string]*WindowOption)

// NewWindowOptions 创建窗口配置
// opts 传多了没用，只取第一个
func NewWindowOptions(name string, opts ...WindowOption) *WindowOption {
	if len(opts) == 0 {
		// 创建默认实例
		WindowOptions[name] = &WindowOption{
			WebviewWindowOptions: application.WebviewWindowOptions{},
		}
	}

	WindowOptions[name] = &opts[0]
	return WindowOptions[name]
}

// CreateWindow 创建窗口 成功后返回一个窗口实例
func CreateWindow(opt WindowOption) (*application.WebviewWindow, error) {
	if opt.Name == "" {
		return nil, fmt.Errorf("窗口名称不能为空")
	}

	opt.BackgroundColour = application.NewRGBA(27, 38, 54, 0)
	opt.BackgroundType = 3
	WindowOptions[opt.Name] = &opt
	return App.Window.NewWithOptions(opt.WebviewWindowOptions), nil
}

// SetWindow 设置窗口属性
func SetWindow(opt WindowOption) error {
	if opt.Name == "" {
		return fmt.Errorf("窗口名称不能为空")
	}

	window := Windows[opt.Name]
	if WindowOptions[opt.Name] == nil {
		WindowOptions[opt.Name] = NewWindowOptions(opt.Name)
	}

	// 标题
	if opt.Title != "" {
		window.SetTitle(opt.Title)
	}

	// 全屏
	if opt.IsFullscreen {
		window.Fullscreen()
	} else {
		window.UnFullscreen()
	}

	// 是否置顶
	if opt.AlwaysOnTop {
		window.SetAlwaysOnTop(true)
	} else {
		window.SetAlwaysOnTop(false)
	}

	// 设置 URL 防止闪烁
	if opt.URL != WindowOptions[opt.Name].URL {
		window.SetURL(opt.URL)
		WindowOptions[opt.Name].URL = opt.URL
	}

	// 设置位置
	if opt.X == 0 || opt.Y == 0 {
		window.Center()
	} else {
		window.SetPosition(opt.X, opt.Y)
	}

	window.SetSize(opt.Width, opt.Height)

	// 是否隐藏，最后设置防止显示以后再设置其他的属性发生闪烁
	if opt.Hidden {
		window.Hide()
	} else {
		window.Show()
	}

	return nil
}
