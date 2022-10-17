package gins

import (
	"gin-boilerplate/gin-sever/middlewares"
	"gin-boilerplate/gin-sever/routers"
	"github.com/gin-gonic/gin"
)

var GinEngine *gin.Engine

func Init() {
	// Init a gin engine
	GinEngine = gin.New()

	// Add preset cors middleware, you can add your own middleware
	GinEngine.Use(middlewares.Cors())
	// Add default recovery middleware, you can add your own middleware
	GinEngine.Use(gin.Recovery())

	// Set base uri, you can load it from config/env or edite here directly
	baseUri := "api"

	// add different api groups
	routers.LoadExampleApiGroup(baseUri)
}

func InitDefault() {
	GinEngine = gin.Default()
}
