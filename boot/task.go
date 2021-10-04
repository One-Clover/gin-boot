package boot

import (
	"github.com/robfig/cron/v3"
	"sync"
)

//任务列表
var taskList chan *TaskExecutor

var once sync.Once

func getTaskList() chan *TaskExecutor {
	once.Do(func() {
		//初始化chan
		taskList = make(chan *TaskExecutor)
	})
	return taskList
}

type TaskFunc func(params ...interface{})

type TaskExecutor struct {
	f        TaskFunc
	p        []interface{}
	callback func()
}

func NewTaskExecutor(f TaskFunc, p []interface{}, callback func()) *TaskExecutor {
	return &TaskExecutor{f: f, p: p, callback: callback}
}

func (this *TaskExecutor) Exec() {
	this.f(this.p...)
}

func Task(f TaskFunc, cb func(), params ...interface{}) {
	if f == nil {
		return
	}
	go func() {
		//增加任务队列
		getTaskList() <- NewTaskExecutor(f, params, cb)
	}()
}

func init() {
	chlist := getTaskList()
	go func() {
		for t := range chlist {
			doTask(t)
		}
	}()
}

func doTask(t *TaskExecutor) {
	go func() {
		defer func() {
			if t.callback != nil {
				t.callback()
			}
		}()
		t.Exec()
	}()
}

var onceCron sync.Once

var taskCron *cron.Cron

func getCronTask() *cron.Cron {
	onceCron.Do(func() {
		taskCron = cron.New(cron.WithSeconds())
	})
	return taskCron
}
