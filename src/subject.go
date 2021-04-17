package main

import (
	"time"
)

/*
	File: subject.go
	Author: luckyFang
	Date: 2022-04-16
	Desc： 观察者对象
*/

type Subject struct {
	observerList []Observable
}

// 创建一个观察者
func NewSubject() Subject {
	return Subject{}
}

// 添加一个任务到观察列表
func (this *Subject) add(observer ScheduleTask) {
	this.observerList = append(this.observerList, &observer)
}

// 通知所有被观察对象
func (this *Subject) notifyAll() {
	now := time.Now()
	// 推送全部
	for index := range this.observerList {
		go this.observerList[index].trigger(now)
	}
}

// 从观察列表中删除一个观察者
func (this *Subject) remove(index int) {
	this.observerList = append(this.observerList[:index], this.observerList[index+1:]...)
}
