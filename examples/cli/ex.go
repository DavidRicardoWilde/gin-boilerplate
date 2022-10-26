package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Action = test
	err := app.Run(os.Args)
	if err != nil {
		return
	}
	time.Sleep(5 * time.Second)
}

func test(ctx *cli.Context) error {
	fmt.Println("start")
	return nil
}
