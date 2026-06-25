package main

import (
	"fmt"
	"log"
	"math/rand"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

// 顺序切换时每个显示器的当前图片索引
var sequentialIndex = make(map[int]int)

// switchWallpaper 根据策略切换壁纸
func switchWallpaper() {
	switch Config.Strategy {
	case SwitchMouseScreen:
		idx := getDisplayIndex()
		if idx >= 0 {
			doSwitch(idx)
		}
	case SwitchRandom:
		if len(Config.MonitorDirs) == 0 {
			return
		}
		monitors := make([]int, 0, len(Config.MonitorDirs))
		for k := range Config.MonitorDirs {
			monitors = append(monitors, k)
		}
		idx := monitors[rand.Intn(len(monitors))]
		doSwitch(idx)
	case SwitchSequential:
		if Config.SwitchAllMonitor {
			for idx := range Config.MonitorDirs {
				doSwitch(idx)
			}
		} else {
			for idx := range Config.MonitorDirs {
				doSwitch(idx)
				break
			}
		}
	}
}

// doSwitch 对指定显示器执行壁纸切换
func doSwitch(monitorIndex int) {
	images := scanMonitorDirs(monitorIndex)
	if len(images) == 0 {
		fmt.Printf("显示器 %d 没有找到图片\n", monitorIndex)
		return
	}

	var imagePath string
	if Config.Strategy == SwitchSequential {
		idx := sequentialIndex[monitorIndex] % len(images)
		imagePath = images[idx]
		sequentialIndex[monitorIndex] = idx + 1
	} else {
		imagePath = images[rand.Intn(len(images))]
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

	monitorID, err := getMonitorDevicePathAt(dw, uint32(monitorIndex))
	if err != nil {
		fmt.Printf("获取显示器 %d 路径失败: %v\n", monitorIndex, err)
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

// 配置壁纸设置选项
func configSet(monitorID, imagePath string) {
	if err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_DISABLE_OLE1DDE); err != nil {
		log.Fatalf("CoInitializeEx failed: %v", err)
	}
	defer ole.CoUninitialize()

	unknown, err := ole.CreateInstance(CLSID_DesktopWallpaper, nil)
	if err != nil {
		log.Fatalf("CreateInstance failed: %v", err)
	}
	defer unknown.Release()

	dw, err := unknown.QueryInterface(IID_IDesktopWallpaper)
	if err != nil {
		log.Fatalf("QueryInterface failed: %v", err)
	}
	defer dw.Release()

	// 设置填充模式：0=Center, 1=Tile, 2=Stretch, 3=Fit, 4=Fill, 5=Span
	if err := setPosition(dw, 4); err != nil {
		log.Fatalf("SetPosition failed: %v", err)
	}

	if err := setWallpaper(dw, monitorID, imagePath); err != nil {
		log.Fatalf("SetWallpaper failed: %v", err)
	}
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
