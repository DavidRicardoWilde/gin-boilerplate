package main

import (
	"fmt"
	"gin-boilerplate/utils/loggers"
)

func main() {
	fmt.Println("starting init")
	fmt.Println("Init finished")
	loggers.ApiLog.Infof("Info log")
	loggers.ApiLog.Debugf("Debug log")
	loggers.ApiLog.Warnf("Warn log")
	loggers.ApiLog.Errorf("Error log")
	//loggers.ApiLog.Panicf("Panic log")
	loggers.ApiLog.Tracef("Trace log")
	//loggers.ApiLog.Fatalf("Fatal log")
}
