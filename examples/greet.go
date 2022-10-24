package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	i := 0

	pid := os.Getppid()

	fmt.Printf("pid: %d \n", pid)

	prc := exec.Command("ps", "-p", strconv.Itoa(pid), "-v")
	out, err := prc.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "start the server",
				Action: func(ctx *cli.Context) error {
					startServer()
					return nil
				},
			},
			{
				Name:  "add",
				Usage: "test",
				Action: func(ctx *cli.Context) error {
					args := ctx.Args().First()
					argsInt, _ := strconv.Atoi(args)
					fmt.Println("i: ", argsInt+i)
					return nil
				},
			},
			//{
			//	Name:  "stop",
			//	Usage: "stop the server",
			//	Action: func(ctx *cli.Context) error {
			//
			//	},
			//},
		},
	}

	//time.Sleep(15 * time.Second)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func startServer() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: r.Handler(),
	}

	defer func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println("server start error: ", err)
		}
	}()
}
