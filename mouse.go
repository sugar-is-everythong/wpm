package main

import (
	"github.com/go-vgo/robotgo"
)

// 获取鼠标所在显示器的索引
func getDisplayIndex() int {
	mx, my := robotgo.Location() // 获取鼠标位置
	num := robotgo.DisplaysNum() // 显示器数量

	for i := 0; i < num; i++ {
		x, y, w, h := robotgo.GetDisplayBounds(i)

		// 判断鼠标是否在这个显示器范围内
		if mx >= x && mx < x+w && my >= y && my < y+h {
			return i
		}
	}

	return -1
}
