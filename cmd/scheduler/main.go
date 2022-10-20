package main

import (
	"context"
	"gin-boilerplate/tasks"
	"gin-boilerplate/tasks/examples"
	"gin-boilerplate/utils/loggers"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"sync"
	"time"
)

var statusMutex = &sync.Mutex{}
var statusMap = make(map[string]bool)

func taskWrapper(task tasks.ITask) func() {
	return func() {
		flag := false

		(func() {
			statusMutex.Lock()
			defer statusMutex.Unlock()

			flag = statusMap[task.Name()]
			statusMap[task.Name()] = true
		})()

		if !flag {
			return
		}

		task.Exec()

		(func() {
			statusMutex.Lock()
			defer statusMutex.Unlock()

			statusMap[task.Name()] = false
		})()
	}
}

func main() {
	// Init log system, set your customized logger config
	loggers.InitWebServerLog()

	// Init cron
	cronJob := cron.New()

	// give me a Cron Object with a UTC location time zone and using logrus instead default logger interface, and custom job wrapper
	//cron.New(cron.WithLocation(time.UTC))

	// jobs
	jobId, err := cronJob.AddFunc("@every 5s", taskWrapper(new(examples.ExampleTask)))
	// create a job which run at every day 0:00:00
	//jobId, err := cronJob.AddFunc("0 0 0 * * *", taskWrapper(new(examples.ExampleTask)))
	// create a job which run at daily 1 clock
	//jobId, err := cronJob.AddFunc("0 0 1 * * *", taskWrapper(new(examples.ExampleTask)))
	// create a job which run at daily 24 clock
	//jobId, err := cronJob.AddFunc("@daily", taskWrapper(new(examples.ExampleTask)))

	if err != nil {
		loggers.ScheduleLog.WithError(err).Errorln("error when add job")
	}

	loggers.ScheduleLog.Infof("Example task added! Job id: %d \n", jobId)
	cronJob.Start()

	// quit server gracefully
	quitSignal := make(chan os.Signal, 20)
	signal.Notify(quitSignal, os.Interrupt)
	<-quitSignal

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	select {
	case <-cronJob.Stop().Done():
	case <-ctx.Done():
	}

}
