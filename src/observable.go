package main

import "time"

/*
	File: observable.go
	Author: luckyFang
	Date: 2022-04-16
	Desc： 触发器接口
*/

// 观察者接口
type Observable interface {
	// 触发器
	trigger(now time.Time)
}
