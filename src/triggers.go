package main

/*
	File: triggers.go
	Author: luckyFang
	Date: 2022-04-16
	Desc： 这里用来注册触发器
*/

// 注册触发对象
func registerTriggers() *Subject {
	// 观察者对象
	subject := NewSubject()
	subject.add(NewScheduleTask("早上7点推送早安", TTime{
		Year: 0,
		Mon:  0,
		Day:  0,
		Hour: 7,
		Min:  0,
		Sec:  0,
	}, func() {
		goodMorning()
	}))

	subject.add(NewScheduleTask("tip1", TTime{
		Year: 0,
		Mon:  0,
		Day:  0,
		Hour: 15,
		Min:  0,
		Sec:  0,
	}, func() {
		pushMsg("提示", "今天学习了吗？")
	}))

	subject.add(NewScheduleTask("tip1", TTime{
		Year: 0,
		Mon:  0,
		Day:  0,
		Hour: 20,
		Min:  0,
		Sec:  0,
	}, func() {
		pushMsg("提示", "今天的任务都完成了吗？")
	}))

	subject.add(NewScheduleTask("晚上23点推送晚安", TTime{
		Year: 0,
		Mon:  0,
		Day:  0,
		Hour: 23,
		Min:  0,
		Sec:  0,
	}, func() {
		goodEvening()
	}))

	return &subject
}
