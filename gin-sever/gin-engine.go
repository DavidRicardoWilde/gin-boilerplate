package gins

import (
	"gin-boilerplate/gin-sever/middlewares"
	"gin-boilerplate/gin-sever/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

var GinEngine *gin.Engine

// Init sets GinEngine as custom gin engine and uses custom middlewares
func Init() {
	// Init a gin engine
	GinEngine = gin.New()

	// Add preset cors middleware, you can add your own middleware
	GinEngine.Use(middlewares.Cors())
	// Add default recovery middleware, you can add your own middleware
	GinEngine.Use(gin.Recovery())
}

// InitDefault init a default gin engine
func InitDefault() {
	GinEngine = gin.Default()
}

// LoadApiGroups loads api groups
func LoadApiGroups() {
	// Set base uri
	baseUri := "api"

	// add different api groups
	GinEngine.MaxMultipartMemory = 8 << 20 // set custom memory 8 MiB
	routers.LoadExampleApiGroup(baseUri, GinEngine)
}

func LoadFileSystem() {
	GinEngine.Static("/assets", "./assets")
	GinEngine.StaticFS("/more_static", http.Dir("my_file_system"))
	GinEngine.StaticFile("/favicon.ico", "./resources/favicon.ico")
}
