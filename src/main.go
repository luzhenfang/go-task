package main

import (
	"fmt"
)

/*
	File: main.go
	Author: luckyFang
	Date: 2022-04-16
	Desc： 创建定时任务服务
*/

func init() {
	banner := `                    __                __    
.-----.-----.______|  |_.---.-.-----.|  |--.
|  _  |  _  |______|   _|  _  |__ --||    < 
|___  |_____|      |____|___._|_____||__|__|
|_____|                                     
`
	fmt.Println(banner)
	fmt.Println("go-task v1.1")
}

func main() {
	NewServer().registerTriggers(registerTriggers()).start()
}
