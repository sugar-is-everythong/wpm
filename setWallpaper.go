package main

import (
	"fmt"
	"math/rand"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

// 顺序切换时每个显示器的当前图片索引，对应 displayMap 的key
var sequentialIndex = 0

// ManualSwitch 手动切换指定显示器壁纸，不走策略，并重置该显示器定时器
func (g *GreetService) ManualSwitch(monitorIndex int) {
	doSwitch(monitorIndex)
	ResetMonitorTimer(monitorIndex)
}

// switchWallpaper 根据策略切换壁纸
func switchWallpaper() {
	// 每次都加载一下显示器的数量，防止显示器有变化，不过一般也没有变化
	if len(App.Screen.GetAll()) != len(displayMap) {
		initDisplayMap()
	}

	switch Config.Strategy {
	case SwitchAllMonitors: // 所有显示器同时切换
		for idx := range Config.MonitorDirs {
			doSwitch(idx)
		}
	case SwitchMouseScreen: // 鼠标所在显示器切换 ,其实这个策略没有用，因为我打算做一个右键切换对应显示器壁纸的功能
		idx := getDisplayIndex()
		if idx >= 0 {
			doSwitch(idx)
		}
	case SwitchSequential: // 每次只切换一个显示器，按顺序轮转
		doSwitch(sequentialIndex)
		sequentialIndex++
		if sequentialIndex >= len(displayMap) {
			sequentialIndex = 0
		}
	case SwitchRandom: // 随机选一个显示器切换
		if len(Config.MonitorDirs) == 0 {
			return
		}

		// 只查找设置过目录的显示器
		s := make([]int, 0, len(Config.MonitorDirs))
		for k, dirs := range Config.MonitorDirs {
			if len(dirs.Dirs) > 0 {
				s = append(s, k)
			}
		}
		doSwitch(s[rand.Intn(len(s))])
	}
}

// doSwitch 对指定显示器执行壁纸切换
func doSwitch(monitorIndex int) {
	imagePath, err := randImg(monitorIndex)
	if err != nil {
		fmt.Printf("获取显示器 %d 随机图片失败: %v\n", monitorIndex, err)
		return
	}

	// 获取显示器设备路径
	if err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_DISABLE_OLE1DDE); err != nil {
		fmt.Printf("COM初始化失败: %v\n", err)
		return
	}
	defer ole.CoUninitialize()

	unknown, err := ole.CreateInstance(CLSID_DesktopWallpaper, nil)
	if err != nil {
		fmt.Printf("创建DesktopWallpaper对象失败: %v\n", err)
		return
	}
	defer unknown.Release()

	dw, err := unknown.QueryInterface(IID_IDesktopWallpaper)
	if err != nil {
		fmt.Printf("查询接口失败: %v\n", err)
		return
	}
	defer dw.Release()

	monitorID, ok := displayMap[monitorIndex]
	if !ok {
		fmt.Printf("获取显示器 %d ID失败\n", monitorIndex)
		return
	}

	if err := setPosition(dw, int32(Config.WallpaperFit)); err != nil {
		fmt.Printf("设置填充模式失败: %v\n", err)
		return
	}

	if err := setWallpaper(dw, monitorID, imagePath); err != nil {
		fmt.Printf("设置壁纸失败: %v\n", err)
		return
	}
	fmt.Printf("显示器 %d 壁纸已切换: %s\n", monitorIndex, imagePath)
}

// 壁纸填充模式
func setPosition(dw *ole.IDispatch, position int32) error {
	return comCall(dw, idwSetPosition, uintptr(position))
}

// 设置壁纸
func setWallpaper(dw *ole.IDispatch, monitorID, imagePath string) error {
	monitorPtr, err := syscall.UTF16PtrFromString(monitorID)
	if err != nil {
		return err
	}
	imagePtr, err := syscall.UTF16PtrFromString(imagePath)
	if err != nil {
		return err
	}
	return comCall(dw, idwSetWallpaper,
		uintptr(unsafe.Pointer(monitorPtr)),
		uintptr(unsafe.Pointer(imagePtr)),
	)
}
