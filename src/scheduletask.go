package main

import "time"

/*
	File: scheduletask.go
	Author: luckyFang
	Date: 2022-04-16
	Desc： 定时任务类
*/

type CALL func()

// 定时任务
type ScheduleTask struct {
	name string
	t    TTime
	proc CALL // trigger
}

// 新建一个定时任务
func NewScheduleTask(name string, t TTime, proc CALL) ScheduleTask {
	return ScheduleTask{name: name, t: t, proc: proc}
}

// 触发器
func (this *ScheduleTask) trigger(now time.Time) {
	// 精确触发
	if this.t.Year == now.Year() && this.t.Mon == int(now.Month()) && this.t.Day == now.Day() &&
		this.t.Hour == now.Hour() && this.t.Min == now.Minute() && this.t.Sec == now.Second() {
		this.proc()
	}
	// 每天触发
	if this.t.Year == 0 && this.t.Mon == 0 && this.t.Day == 0 &&
		this.t.Hour == now.Hour() && this.t.Min == now.Minute() && this.t.Sec == now.Second() {
		this.proc()
	}
	// 每小时触发
	if this.t.Year == 0 && this.t.Mon == 0 && this.t.Day == 0 &&
		this.t.Hour == 0 && this.t.Min == now.Minute() && this.t.Sec == now.Second() {
		this.proc()
	}
	// 每分钟触发
	if this.t.Year == 0 && this.t.Mon == 0 && this.t.Day == 0 &&
		this.t.Hour == 0 && this.t.Min == 0 && this.t.Sec == now.Second() {
		this.proc()
	}
}
