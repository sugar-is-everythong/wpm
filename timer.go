package main

import (
	"time"

	"github.com/go-co-op/gocron/v2"
)

func init() {
	s, _ := gocron.NewScheduler()

	s.NewJob(
		gocron.DurationJob(5*time.Second),
		gocron.NewTask(func() {
			//fmt.Println("任务执行:", time.Now().Format("15:04:05"))
		}),
	)

	s.Start()
}
