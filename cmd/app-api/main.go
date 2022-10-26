package main

import (
	"context"
	"gin-boilerplate/configs"
	gins "gin-boilerplate/gin-sever"
	cfgs "gin-boilerplate/utils/configs"
	"gin-boilerplate/utils/dbs"
	"gin-boilerplate/utils/loggers"
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Get pid when started
	pid := os.Getpid()

	// Start your app with cli, like: gin-boilerplate --config config-file-path --env dev
	app := cli.NewApp()
	app.Name = "gin-boilerplate"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Usage:       "input config file name you want to use, or using default",
			Value:       "./configs/config.toml",
			DefaultText: "./configs/config.toml",
			Required:    true,
			Action: func(c *cli.Context, path string) error {
				// Load config file
				configs.LoadConfigFile(path)
				return nil
			},
		},
		&cli.StringFlag{
			Name:        "env",
			Usage:       "input env you want to use or app will running on dev. env.",
			Value:       "prod",
			DefaultText: "prod",
			Destination: &configs.Env,
			Required:    true,
		},
		// Add other flags you need...
	}
	app.Action = func(c *cli.Context) error {
		// Init configs
		cfgs.InitAllConfigs()

		// Init log system, set your customized logger config
		loggers.InitApiServerLog()

		// Init gorm database client which will be according to the config you set.
		dbs.InitGlobalDBClient()

		// Set api server with gin engine
		gins.Init()

		// Load api group routers
		gins.LoadApiGroups()
		// Load file server
		//gins.LoadFileServer()

		// Setting http api server
		server := http.Server{
			Addr:           cfgs.GetGlobalAppServerCfg().ServerPort,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			Handler:        gins.GinEngine,
			MaxHeaderBytes: 1 << 20,
		}

		// Start http api server
		go func() {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				loggers.ApiLog.Fatal("listen: ", err)
			}
		}()
		loggers.ApiLog.Infoln("Web server started, pid: ", pid)

		// Quit server gracefully
		quitSignal := make(chan os.Signal, 20)
		signal.Notify(quitSignal, os.Interrupt)
		<-quitSignal

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			loggers.ApiLog.WithError(err).Error("Server shutdown failed")
		}
		loggers.ApiLog.Info("Server shutdown")

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		loggers.ApiLog.WithError(err).Error("error when run app")

	}
}
