package main

import "time"

/*
	File: server.go
	Author: luckyFang
	Date: 2022-04-16
	Desc： 定时任务服务类
*/

type Server struct {
	observerSubject *Subject
}

// 创建一个定时任务服务
func NewServer() *Server {
	return &Server{}
}

// 注册触发器
func (this *Server) registerTriggers(subject *Subject) *Server {
	this.observerSubject = subject
	return this
}

// 开启服务
func (this *Server) start() {
	for {
		this.observerSubject.notifyAll()
		time.Sleep(time.Duration(1) * time.Second)
	}
}
