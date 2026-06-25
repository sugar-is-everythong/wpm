package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// SwitchStrategy 自动切换策略
type SwitchStrategy int

const (
	SwitchAllMonitors SwitchStrategy = iota // 0 协同切换（所有显示器同时切换）
	SwitchSequential                        // 1 顺序切换
	SwitchRandom                            // 2 随机切换
	SwitchMouseScreen                       // 3 鼠标所在屏幕
)

// WallpaperFit 壁纸应用模式
type WallpaperFit int

const (
	FitCenter  WallpaperFit = iota // 0 居中
	FitTile                        // 1 平铺
	FitStretch                     // 2 拉伸
	FitFit                         // 3 适应
	FitFill                        // 4 填充
	FitSpan                        // 5 跨屏
)

// AppConfig 应用配置
type AppConfig struct {
	AutoStart           bool             `json:"autoStart"`           // 程序启动后是否立即执行定时切换
	RecursiveDir        bool             `json:"recursiveDir"`        // 扫描壁纸目录时是否递归子目录
	Strategy            SwitchStrategy   `json:"strategy"`            // 自动切换策略
	RegisterContextMenu bool             `json:"registerContextMenu"` // 是否注册到系统右键菜单
	MonitorDirs         map[int][]string `json:"monitorDirs"`         // 每个显示器的壁纸目录
	Interval            int              `json:"interval"`            // 切换间隔（分钟）
	AutoMatchRatio      bool             `json:"autoMatchRatio"`      // 自动匹配显示器比例
	WallpaperFit        WallpaperFit     `json:"wallpaperFit"`        // 壁纸应用模式
}

// Config 全局配置实例
var Config = AppConfig{
	AutoStart:           true,
	RecursiveDir:        true,
	Strategy:            SwitchMouseScreen,
	RegisterContextMenu: true,
	MonitorDirs:         make(map[int][]string),
	Interval:            30,
	AutoMatchRatio:      true,
	WallpaperFit:        FitFill,
}

var configOnce sync.Once

// configPath 获取配置文件路径
func configPath() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		dir = "."
	}
	return filepath.Join(dir, "wpm", "config.json")
}

// LoadConfig 从文件加载配置
func LoadConfig() error {
	path := configPath()
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // 配置文件不存在，使用默认值
		}
		return fmt.Errorf("读取配置文件失败: %w", err)
	}
	return json.Unmarshal(data, &Config)
}

// SaveConfig 保存配置到文件
func SaveConfig() error {
	path := configPath()
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建配置目录失败: %w", err)
	}
	data, err := json.MarshalIndent(Config, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化配置失败: %w", err)
	}
	return os.WriteFile(path, data, 0644)
}

// ConfigService 配置服务，供前端调用
type ConfigService struct{}

// GetConfig 获取当前配置
func (c *ConfigService) GetConfig() AppConfig {
	return Config
}

// Save 保存完整配置，同时重启定时器
func (c *ConfigService) Save(cfg AppConfig) error {
	prevInterval := Config.Interval
	Config = cfg
	if Config.MonitorDirs == nil {
		Config.MonitorDirs = make(map[int][]string)
	}
	if err := SaveConfig(); err != nil {
		return err
	}
	// 间隔变化时重启定时器
	if Config.Interval != prevInterval || Config.Interval > 0 {
		RestartTimer()
	}
	return nil
}

// initConfig 初始化加载配置
func initConfig() {
	configOnce.Do(func() {
		if err := LoadConfig(); err != nil {
			fmt.Printf("加载配置失败: %v，使用默认配置\n", err)
		}
		if Config.MonitorDirs == nil {
			Config.MonitorDirs = make(map[int][]string)
		}
		// 启动时根据 autoStart 决定是否启动定时器
		if Config.AutoStart && Config.Interval > 0 {
			StartTimer(Config.Interval)
		}
	})
}
