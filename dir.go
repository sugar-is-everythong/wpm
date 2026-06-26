package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
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

// collectImageDirs 收集包含图片的目录列表
func collectImageDirs(dir string, recursive bool) []string {
	var result []string
	if recursive {
		_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil || !info.IsDir() {
				return nil
			}
			entries, e := os.ReadDir(path)
			if e != nil {
				return nil
			}
			for _, entry := range entries {
				if !entry.IsDir() && isImageFile(entry.Name()) {
					result = append(result, path)
					break
				}
			}
			return nil
		})
	} else {
		entries, err := os.ReadDir(dir)
		if err != nil {
			return nil
		}
		for _, entry := range entries {
			if !entry.IsDir() && isImageFile(entry.Name()) {
				result = append(result, dir)
				break
			}
		}
	}
	return result
}

// isLandscape 读取图片文件头，判断是否为横向图片（宽 >= 高）
func isLandscape(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()
	cfg, _, err := image.DecodeConfig(f)
	if err != nil {
		return false, err
	}
	return cfg.Width >= cfg.Height, nil
}

// randImg 从指定显示器的配置目录中随机获取一张图片路径
func randImg(monitorIndex int) (string, error) {
	dirs, ok := Config.MonitorDirs[monitorIndex]
	if !ok || len(dirs.Dirs) == 0 {
		return "", fmt.Errorf("显示器 %d 没有配置目录", monitorIndex)
	}

	// 获取显示器方向（横/竖）
	screens := App.Screen.GetAll()
	var monitorLandscape bool
	if monitorIndex < len(screens) {
		monitorLandscape = screens[monitorIndex].Size.Width >= screens[monitorIndex].Size.Height
	}

	// 收集所有包含图片的目录
	var imageDirs []string
	for _, dir := range dirs.Dirs {
		imageDirs = append(imageDirs, collectImageDirs(dir, Config.RecursiveDir)...)
	}
	if len(imageDirs) == 0 {
		return "", fmt.Errorf("显示器 %d 的目录中没有找到图片", monitorIndex)
	}

	// 随机选一个包含图片的目录
	targetDir := imageDirs[rand.Intn(len(imageDirs))]

	// 从该目录中收集图片
	entries, err := os.ReadDir(targetDir)
	if err != nil {
		return "", fmt.Errorf("读取目录失败: %v", err)
	}
	var images []string
	for _, entry := range entries {
		if !entry.IsDir() && isImageFile(entry.Name()) {
			images = append(images, filepath.Join(targetDir, entry.Name()))
		}
	}
	if len(images) == 0 {
		return "", fmt.Errorf("目录中没有图片")
	}

	// 如果开启比例匹配，过滤方向一致的图片
	if Config.AutoMatchRatio {
		var matched []string
		for _, img := range images {
			landscape, err := isLandscape(img)
			if err != nil {
				continue // 无法读取尺寸的图片跳过
			}
			if landscape == monitorLandscape {
				matched = append(matched, img)
			}
		}
		if len(matched) > 0 {
			images = matched
		}
		// 没有匹配的就用全部图片（fallback）
	}

	return images[rand.Intn(len(images))], nil
}
