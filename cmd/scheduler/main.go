package main

import (
	"context"
	"fmt"
	"gin-boilerplate/tasks"
	"gin-boilerplate/tasks/examples"
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
	cronJob := cron.New()
	// jobs
	jobId, err := cronJob.AddFunc("@every 5s", taskWrapper(new(examples.ExampleTask)))
	if err != nil {
		// error log
	}

	fmt.Sprintf("Example task added! Job id: %d", jobId)
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
