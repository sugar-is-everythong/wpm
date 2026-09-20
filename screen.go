package main

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type GreetService struct {
}

var (
	CLSID_DesktopWallpaper = ole.NewGUID("C2CF3110-460E-4FC1-B9D0-8A1C0C9CC4BD")
	IID_IDesktopWallpaper  = ole.NewGUID("B92B56A9-8B55-4E14-9A89-0199BBB6F93B")

	// displayMap 缓存所有显示器，key 为索引，程序启动时加载
	displayMap = make(map[int]string)
)

// IDesktopWallpaper 接口方法索引（相对于 IUnknown 3 个方法之后）
const (
	iUnknownMethodCount          = 3
	idwSetWallpaper              = 0
	idwGetWallpaper              = 1
	idwGetMonitorDevicePathAt    = 2
	idwGetMonitorDevicePathCount = 3
	idwGetMonitorRECT            = 4
	idwSetBackgroundColor        = 5
	idwGetBackgroundColor        = 6
	idwSetPosition               = 7
	idwGetPosition               = 8
	idwSetSlideshow              = 9
	idwGetSlideshow              = 10
	idwEnable                    = 11
)

// initDisplayMap 程序启动时调用，缓存所有显示器信息到 displayMap
func initDisplayMap() {
	// 初始化 COM，使用单线程单元（STA）
	if err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_DISABLE_OLE1DDE); err != nil {
		log.Fatalf("CoInitializeEx failed: %v", err)
	}
	defer ole.CoUninitialize()

	// 创建 DesktopWallpaper COM 对象
	unknown, err := ole.CreateInstance(CLSID_DesktopWallpaper, nil)
	if err != nil {
		log.Fatalf("CreateInstance failed: %v", err)
	}
	defer unknown.Release()

	// 查询 IDesktopWallpaper 接口
	dw, err := unknown.QueryInterface(IID_IDesktopWallpaper)
	if err != nil {
		log.Fatalf("QueryInterface failed: %v", err)
	}
	defer dw.Release()

	num, err := getMonitorCount(dw)
	if err != nil {
		log.Fatalf("GetMonitorCount 错误: %v", err)
	}
	for i := uint32(0); i < num; i++ {
		monitorID, err := getMonitorDevicePathAt(dw, i)
		if err != nil {
			log.Printf("  显示器 %d: 获取失败: %v", i, err)
			continue
		}
		displayMap[int(i)] = monitorID
		fmt.Printf("缓存显示器 %d: %s %v\n", i, monitorID, displayMap[int(i)])
	}
}

func (g *GreetService) GetDisplay() map[int]string {
	return displayMap
}

// getMonitorCount 获取显示器数量 调用 IDesktopWallpaper::GetMonitorDevicePathCount
func getMonitorCount(dw *ole.IDispatch) (uint32, error) {
	var count uint32
	if err := comCall(dw, idwGetMonitorDevicePathCount, uintptr(unsafe.Pointer(&count))); err != nil {
		return 0, err
	}
	return count, nil
}

// getMonitorDevicePathAt 获取显示器设备路径/ID 调用 IDesktopWallpaper::GetMonitorDevicePathAt
// 返回的字符串需要用 CoTaskMemFree 释放
func getMonitorDevicePathAt(dw *ole.IDispatch, index uint32) (string, error) {
	var monitorID *uint16
	if err := comCall(dw, idwGetMonitorDevicePathAt, uintptr(index), uintptr(unsafe.Pointer(&monitorID))); err != nil {
		return "", err
	}
	defer ole.CoTaskMemFree(uintptr(unsafe.Pointer(monitorID)))
	return syscall.UTF16ToString((*[1 << 20]uint16)(unsafe.Pointer(monitorID))[:]), nil
}

// comCall 通过虚表调用指定的 COM 接口方法
func comCall(unk *ole.IDispatch, methodIndex int, args ...uintptr) error {
	vtable := uintptr(unsafe.Pointer(unk.RawVTable))
	methodPtr := *(*uintptr)(unsafe.Pointer(vtable + uintptr(iUnknownMethodCount+methodIndex)*unsafe.Sizeof(uintptr(0))))

	// 第一个参数是 this 指针
	allArgs := append([]uintptr{uintptr(unsafe.Pointer(unk))}, args...)

	ret, _, _ := syscall.SyscallN(methodPtr, allArgs...)
	if ret != 0 {
		return ole.NewError(ret)
	}
	return nil
}
