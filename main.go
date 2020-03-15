package main

import (
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {
	for true {
		time.Sleep(5 * time.Second)
		robotgo.KeyTap("w")
		time.Sleep(1 * time.Second)
		robotgo.KeyTap("w")
		time.Sleep(1 * time.Second)
		robotgo.KeyTap("w")
		time.Sleep(1 * time.Second)
		robotgo.KeyTap("1", "lshift")
		time.Sleep(1 * time.Second)
		robotgo.KeyTap("w")
		time.Sleep(1 * time.Second)
		robotgo.KeyTap("space")
		time.Sleep(3 * time.Second)
		robotgo.KeyTap("s", "lctrl")
		time.Sleep(3 * time.Minute)
	}
}
