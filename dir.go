package main

import (
	"os"
	"path/filepath"
	"strings"
)

// isImageFile 判断是否是支持的图片格式
func isImageFile(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".bmp", ".webp", ".gif", ".tif", ".tiff":
		return true
	}
	return false
}

// scanDir 扫描单个目录，查找图片文件
func scanDir(dir string, recursive bool) []string {
	var images []string
	if recursive {
		_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if !info.IsDir() && isImageFile(info.Name()) {
				images = append(images, path)
			}
			return nil
		})
	} else {
		entries, err := os.ReadDir(dir)
		if err != nil {
			return images
		}
		for _, entry := range entries {
			if !entry.IsDir() && isImageFile(entry.Name()) {
				images = append(images, filepath.Join(dir, entry.Name()))
			}
		}
	}
	return images
}

// scanMonitorDirs 扫描指定显示器的所有目录，返回图片列表
func scanMonitorDirs(monitorIndex int) []string {
	dirs, ok := Config.MonitorDirs[monitorIndex]
	if !ok {
		return nil
	}
	var images []string
	for _, dir := range dirs {
		images = append(images, scanDir(dir, Config.RecursiveDir)...)
	}
	return images
}
