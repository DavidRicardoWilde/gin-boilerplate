package routers

import (
	gins "gin-boilerplate/gin-sever"
	"gin-boilerplate/gin-sever/routers/example"
)

func LoadExampleApiGroup(baseUri string) {
	exGroup := gins.GinEngine.Group(baseUri + "/example")
	{
		// add specified middleware for this group, for example: custom logger
		//exGroup.Use(ginlogrus.Logger(logs.Log))

		// routers
		exGroup.GET("/ping", example.Ping)
	}
}
