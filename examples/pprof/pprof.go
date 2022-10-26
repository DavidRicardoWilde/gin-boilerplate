package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func main() {
	go func() {
		// terminal: $ go tool pprof -http=:8081 http://localhost:6060/debug/pprof/heap
		// web:
		// 1、http://localhost:8081/ui
		// 2、http://localhost:6060/debug/charts
		// 3、http://localhost:6060/debug/pprof
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	r := gin.Default()
	pprof.Register(r)
	//ginpprof.Wrap(r)
	r.GET("/ping", Ping)
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8082")
	if err != nil {
		return
	}
}
