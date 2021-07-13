package tasks

import (
	"errors"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/robfig/cron/v3"
)

var DefaultTasks = newTasks()

// tasks 所有任务的列表
type tasks struct {
	Scheduler *gocron.Scheduler
	Tasks     []*Task
}

// 任务
type Task struct {
	job         *gocron.Job
	Name        string
	Description string
	CronTab     string
	Enable      bool
	JobFunc     interface{}
	Params      []interface{}
}

func newTasks() *tasks {
	ts := &tasks{
		Scheduler: gocron.NewScheduler(time.FixedZone("Asia/Shanghai", 3600*8)),
		Tasks:     []*Task{},
	}
	// 不利己执行定时任务而是等任务被调度
	ts.Scheduler.WaitForScheduleAll()
	return ts
}

// 新增任务
func (t *tasks) AddTask(name, description, cronTab string, enable bool, jobFunc interface{}, params []interface{}, singletonMode bool) error {
	var (
		job *gocron.Job
		err error
	)
	for _, v := range t.Tasks {
		if v.Name == name {
			return errors.New("任务重名")
		}
	}
	// 解析cronTab
	_, err = cron.ParseStandard(cronTab)
	if err != nil {
		return err
	}
	// 如果scheduler为空
	if t.Scheduler == nil {
		t.Scheduler = gocron.NewScheduler(time.FixedZone("Asia/Shanghai", 3600*8))
		t.Scheduler.WaitForScheduleAll()
	}
	if enable {
		job, err = t.Scheduler.Cron(cronTab).Do(jobFunc, params...)
		if singletonMode {
			job.SingletonMode()
		}
	}
	if err != nil {
		return err
	}

	t.Tasks = append(t.Tasks, &Task{
		job:         job,
		Name:        name,
		Description: description,
		CronTab:     cronTab,
		Enable:      enable,
		JobFunc:     jobFunc,
		Params:      params,
	})
	return nil
}

// EnableTask 根据任务名启用任务
func (t *tasks) EnableTask(jobName string) error {
	for _, task := range t.Tasks {
		if task.Name == jobName {
			job, err := t.Scheduler.Cron(task.CronTab).Do(task.JobFunc, task.Params...)
			if err != nil {
				return err
			}
			task.job = job
			task.Enable = true
			return nil
		}
	}
	return nil
}

// StartAsync 非阻塞启动
func (t tasks) StartAsync() {
	t.Scheduler.StartAsync()
}

// DisableJob 通过任务名停用任务
func (t tasks) DisableJob(jobName string) {
	for _, task := range t.Tasks {
		if task.Name == jobName {
			t.Scheduler.RemoveByReference(task.job)
			task.job = nil
			task.Enable = false
			return
		}
	}
}
