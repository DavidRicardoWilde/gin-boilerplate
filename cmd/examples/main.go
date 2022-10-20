package main

import (
	"fmt"
	"gin-boilerplate/utils/loggers"
)

func main() {
	fmt.Println("starting init")
	loggers.InitWebServerLog()
	fmt.Println("Init finished")
	loggers.Log.Infof("Info log")
	loggers.Log.Debugf("Debug log")
	loggers.Log.Warnf("Warn log")
	loggers.Log.Errorf("Error log")
	//loggers.Log.Panicf("Panic log")
	loggers.Log.Tracef("Trace log")
	//loggers.Log.Fatalf("Fatal log")
}
