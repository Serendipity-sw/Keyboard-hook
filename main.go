package main

import (
	"flag"
	"github.com/go-vgo/robotgo"
	"github.com/smtc/glog"
	"github.com/swgloomy/gutil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	screenWidth  int
	screenHeight int
	debugFlag    = flag.Bool("d", false, "debug mode") //是否为调试模式
)

func main() {
	flag.Parse()

	gutil.LogInit(*debugFlag, "./logs")

	//获取当前显示器大小
	screenWidth, screenHeight = robotgo.GetScaleSize()

	go script()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-c

	glog.Close()
	os.Exit(0)

}

func script() {
	for true {
		roleOperation()
		//每三分钟执行一次方法
		time.Sleep(3 * time.Minute)
	}
}

func roleOperation() {
			//线程等待3秒 使窗口置顶操作正常进行
		time.Sleep(3 * time.Second)

		//将鼠标移动到屏幕中央
		robotgo.MoveMouse(screenWidth/2, screenHeight/2)
		//鼠标左键按住操作
		robotgo.MouseToggle("down")
		//鼠标左键按住向右拖动半个屏幕
		robotgo.DragMouse(screenWidth, screenHeight/2)
		//鼠标左键松开操作 结束视角转向
		robotgo.MouseToggle("up")

		time.Sleep(2 * time.Second)

		//人物动作事件开始
		//操作人物往前移动
		robotgo.KeyToggle("w", "down")
		//移动三秒钟
		time.Sleep(3 * time.Second)
		//结束人物移动
		robotgo.KeyToggle("w", "up")

		time.Sleep(3 * time.Second)
		//操作人物往右移动
		robotgo.KeyToggle("d", "down")
		//移动三秒钟
		time.Sleep(3 * time.Second)
		//结束人物移动
		robotgo.KeyToggle("d", "up")

		time.Sleep(2 * time.Second)
		//控制人物跳动一次
		robotgo.KeyTap("space")
		//等待人物落地.为了防止人物卡顿,等待3秒
		time.Sleep(3 * time.Second)

		//施放一次快捷键技能 对自己  ctrl + s  最好为骑马快捷键或群疗快捷键 或者 3分钟内的大招 凡是一切不需要选中目标施放的技能都可以
		robotgo.KeyTap("s", "lctrl")
}
