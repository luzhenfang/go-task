package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
	File: util.go
	Author: luckyFang
	Date: 2022-04-16
	Desc： 工具类
*/

const (
	token = "your token here"
)

/**
推送消息到微信
	title 推送标题
	content 推送内容
*/
func pushMsg(title string, content string) {
	if len(title) == 0 {
		title = "通知"
	}
	resp, err := http.PostForm("http://pushplus.hxtrip.com/send", url.Values{"token": {token}, "title": {title}, "content": {content}})
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(body))
}

func getWeather() string {
	resp, _ := http.Get("http://lzfblog.cn:8888/weather.txt")
	body, _ := ioutil.ReadAll(resp.Body)
	replace := strings.Replace(string(body), " ", "", -1)
	return replace
}

// 早上好
func goodMorning() {
	pushMsg("早安", getWeather())
}

// 晚安
func goodEvening() {
	pushMsg("晚安", "夜深了早点睡哦!")
}
