package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron/v2"
)

var (
	scheduler  gocron.Scheduler
	timerJob   gocron.Job
	timerMutex sync.Mutex
)

func init() {
	var err error
	scheduler, err = gocron.NewScheduler()
	if err != nil {
		fmt.Printf("创建调度器失败: %v\n", err)
	}
}

// StartTimer 启动壁纸切换定时器（分钟）
func StartTimer(minutes int) {
	timerMutex.Lock()
	defer timerMutex.Unlock()

	if scheduler == nil {
		return
	}

	StopTimer()

	if minutes <= 0 {
		fmt.Println("切换间隔无效，不启动定时器")
		return
	}

	job, err := scheduler.NewJob(
		gocron.DurationJob(time.Duration(minutes)*time.Minute),
		gocron.NewTask(func() {
			switchWallpaper()
		}),
	)
	if err != nil {
		fmt.Printf("创建定时任务失败: %v\n", err)
		return
	}

	timerJob = job
	scheduler.Start()
	fmt.Printf("定时器已启动，间隔 %d 分钟\n", minutes)
}

// StopTimer 停止定时器
func StopTimer() {
	if timerJob != nil {
		_ = scheduler.RemoveJob(timerJob.ID())
		timerJob = nil
	}
}

// RestartTimer 根据当前配置重启定时器
func RestartTimer() {
	if Config.Interval > 0 {
		StartTimer(Config.Interval)
	} else {
		StopTimer()
	}
}
