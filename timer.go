package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron/v2"
)

var (
	scheduler    gocron.Scheduler
	monitorJobs  = make(map[int]gocron.Job) // 每个显示器的定时任务
	monitorMutex sync.Mutex
)

func init() {
	var err error
	scheduler, err = gocron.NewScheduler()
	if err != nil {
		fmt.Printf("创建调度器失败: %v\n", err)
	}
}

// StartMonitorTimer 启动指定显示器的壁纸切换定时器（分钟）
func StartMonitorTimer(monitorIndex int, minutes int) {
	monitorMutex.Lock()
	defer monitorMutex.Unlock()

	if scheduler == nil {
		return
	}

	// 先停止已有的
	stopMonitorTimerLocked(monitorIndex)

	if minutes <= 0 {
		fmt.Printf("显示器 %d 切换间隔无效，不启动定时器\n", monitorIndex)
		return
	}

	job, err := scheduler.NewJob(
		gocron.DurationJob(time.Duration(minutes)*time.Minute),
		gocron.NewTask(func() {
			switchWallpaper()
		}),
	)
	if err != nil {
		fmt.Printf("创建显示器 %d 定时任务失败: %v\n", monitorIndex, err)
		return
	}

	monitorJobs[monitorIndex] = job
	scheduler.Start()
	fmt.Printf("显示器 %d 定时器已启动，间隔 %d 分钟\n", monitorIndex, minutes)
}

// StopMonitorTimer 停止指定显示器的定时器
func StopMonitorTimer(monitorIndex int) {
	monitorMutex.Lock()
	defer monitorMutex.Unlock()
	stopMonitorTimerLocked(monitorIndex)
}

// stopMonitorTimerLocked 内部方法，调用前需持有锁
func stopMonitorTimerLocked(monitorIndex int) {
	if job, ok := monitorJobs[monitorIndex]; ok {
		_ = scheduler.RemoveJob(job.ID())
		delete(monitorJobs, monitorIndex)
	}
}

// ResetMonitorTimer 重置指定显示器的定时器（手动切换时调用）
func ResetMonitorTimer(monitorIndex int) {
	if pref, ok := Config.MonitorDirs[monitorIndex]; ok && pref.Interval > 0 {
		StartMonitorTimer(monitorIndex, pref.Interval)
	}
}

// StopAllTimers 停止所有显示器定时器
func StopAllTimers() {
	monitorMutex.Lock()
	defer monitorMutex.Unlock()
	for k := range monitorJobs {
		stopMonitorTimerLocked(k)
	}
}

// RestartAllTimers 根据当前配置重启所有定时器
func RestartAllTimers() {
	StopAllTimers()
	for k, pref := range Config.MonitorDirs {
		if pref.Interval > 0 {
			StartMonitorTimer(k, pref.Interval)
		}
	}
}
