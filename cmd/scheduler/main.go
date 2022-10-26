package main

import (
	"context"
	cfgs "gin-boilerplate/configs"
	"gin-boilerplate/tasks"
	"gin-boilerplate/tasks/examples"
	"gin-boilerplate/utils/configs"
	"gin-boilerplate/utils/loggers"
	"github.com/robfig/cron/v3"
	"github.com/urfave/cli/v2"
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
	// Get pid when started
	pid := os.Getpid()

	// Start your app with cli, like: scheduler-boilerplate --config config-file-path --env dev
	app := cli.NewApp()
	app.Name = "scheduler-boilerplate"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Usage:       "input config file name you want to use, or using default",
			Value:       "./configs/config.toml",
			DefaultText: "./configs/config.toml",
			Required:    true,
			Action: func(c *cli.Context, path string) error {
				// Load config file
				cfgs.LoadConfigFile(path)
				return nil
			},
		},
		&cli.StringFlag{
			Name:        "env",
			Usage:       "input env you want to use or app will running on dev. env.",
			Value:       "prod",
			DefaultText: "prod",
			Destination: &cfgs.Env,
			Required:    true,
		},
		// Add other flags you need...
	}
	app.Action = func(c *cli.Context) error {
		// Init global config
		configs.InitAllConfigs()

		// Init log system, set your customized logger config
		loggers.InitScheduleLog()

		// Init cron
		cronJob := cron.New()

		// give me a Cron Object with a UTC location time zone and using logrus instead default logger interface, and custom job wrapper
		//cron.New(cron.WithLocation(time.UTC))

		loggers.ScheduleLog.Infoln("Web server started, pid: ", pid)

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

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		loggers.ScheduleLog.WithError(err).Error("error when run app")
	}

}
