package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/sys/windows/registry"
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

// preference 偏好设置
type preference struct {
	Dirs     []string // 壁纸目录
	Interval int      // 切换间隔（分钟），0 表示不自动切换
}

// AppConfig 应用配置
type AppConfig struct {
	AutoStart           bool               `json:"autoStart"`           // 程序启动后是否立即执行定时切换
	RecursiveDir        bool               `json:"recursiveDir"`        // 扫描壁纸目录时是否递归子目录
	Strategy            SwitchStrategy     `json:"strategy"`            // 自动切换策略
	RegisterContextMenu bool               `json:"registerContextMenu"` // 是否注册到系统右键菜单
	MonitorDirs         map[int]preference `json:"monitorDirs"`         // 每个显示器的壁纸目录
	AutoMatchRatio      bool               `json:"autoMatchRatio"`      // 自动匹配显示器比例
	AutoRun             bool               `json:"autoRun"`             // 是否开机自启动
	WallpaperFit        WallpaperFit       `json:"wallpaperFit"`        // 壁纸应用模式
}

// Config 全局配置实例
var Config = AppConfig{
	AutoStart:           false,
	RecursiveDir:        true,
	Strategy:            SwitchMouseScreen,
	RegisterContextMenu: true,
	MonitorDirs:         make(map[int]preference),
	AutoMatchRatio:      true,
	AutoRun:             false,
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

// defaultConfig 生成默认配置，按当前显示器数量为每个显示器生成配置条目
func defaultConfig() AppConfig {
	cfg := Config
	cfg.MonitorDirs = make(map[int]preference)

	// 默认壁纸目录，首次运行先占位，用户可在界面中修改
	pictures := filepath.Join(os.Getenv("USERPROFILE"), "Pictures")
	if home, err := os.UserHomeDir(); err == nil {
		pictures = filepath.Join(home, "Pictures")
	}

	// 保证前端按显示器索引绑定时每个索引都有值
	count := 1
	if App != nil && App.Screen != nil {
		if screens := App.Screen.GetAll(); len(screens) > 0 {
			count = len(screens)
		}
	}
	for i := 0; i < count; i++ {
		cfg.MonitorDirs[i] = preference{
			Dirs:     []string{pictures},
			Interval: 30,
		}
	}
	return cfg
}

// LoadConfig 从文件加载配置，配置文件不存在时生成默认配置文件
func LoadConfig() error {
	path := configPath()
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			Config = defaultConfig()
			return SaveConfig() // 首次运行写出一份符合格式的配置文件
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

// Save 保存完整配置，同时管理各显示器定时器
func (c *ConfigService) Save(cfg AppConfig) error {
	// 记录旧配置中各显示器的切换间隔
	prevIntervals := make(map[int]int)
	for k, v := range Config.MonitorDirs {
		prevIntervals[k] = v.Interval
	}

	Config = cfg
	if Config.MonitorDirs == nil {
		Config.MonitorDirs = make(map[int]preference)
	}
	if err := SaveConfig(); err != nil {
		return err
	}

	// 对比各显示器间隔变化，按需重启定时器
	for k, pref := range Config.MonitorDirs {
		if pref.Interval != prevIntervals[k] {
			StopMonitorTimer(k)
			if pref.Interval > 0 {
				StartMonitorTimer(k, pref.Interval)
			}
		}
	}
	// 停止已删除显示器的定时器
	for k := range prevIntervals {
		if _, ok := Config.MonitorDirs[k]; !ok {
			StopMonitorTimer(k)
		}
	}

	// 处理开机自启动
	if Config.AutoRun {
		Run()
	} else {
		RemoveAutoRun()
	}

	// TODO: 处理右键菜单注册/注销
	return nil
}

// initConfig 初始化加载配置
func initConfig() {
	configOnce.Do(func() {
		if err := LoadConfig(); err != nil {
			fmt.Printf("加载配置失败: %v，使用默认配置\n", err)
		}
		if Config.MonitorDirs == nil {
			Config.MonitorDirs = make(map[int]preference)
		}
		// 启动时根据 autoStart 决定是否立即切换一次
		if Config.AutoStart {
			for k, _ := range displayMap {
				go doSwitch(k)
			}
		}
		// 根据各显示器的 Interval 启动独立定时器
		for k, pref := range Config.MonitorDirs {
			if pref.Interval > 0 {
				StartMonitorTimer(k, pref.Interval)
			}
		}
	})
}

// Run 注册开机自启动
func Run() {
	appName := "塞壁纸"

	exePath, err := os.Executable()
	if err != nil {
		log.Println("获取可执行文件路径失败:", err)
		return
	}
	absPath, err := filepath.Abs(exePath)
	if err != nil {
		log.Println("获取绝对路径失败:", err)
		return
	}

	key, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.WRITE)
	if err != nil {
		log.Println("注册表打开失败:", err)
		return
	}
	defer key.Close()

	err = key.SetStringValue(appName, absPath)
	if err != nil {
		log.Println("设置自启动项失败:", err)
		return
	}
}

// RemoveAutoRun 移除开机自启动
func RemoveAutoRun() {
	appName := "塞壁纸"

	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.WRITE)
	if err != nil {
		// 键不存在，无需处理
		return
	}
	defer key.Close()

	_ = key.DeleteValue(appName)
}
