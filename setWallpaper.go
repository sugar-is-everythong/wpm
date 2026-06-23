package main

import (
	"log"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

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
