package routers

import (
	"fmt"
	"gin-boilerplate/gin-sever/routers/example"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoadExampleApiGroup(baseUri string, ginEngine *gin.Engine) {
	exGroup := ginEngine.Group(baseUri + "/example/v1")
	{
		// add specified middleware for this group, for example: custom logger
		//exGroup.Use(ginlogrus.Logger(logs.Log))

		// routers with all example in gin official documentation
		exGroup.GET("/ping", example.Ping)
		exGroup.GET("/get", example.Get)
		exGroup.GET("/asciiJson", example.AsciiJson)
		exGroup.GET("/getDataB", example.GetDataB)
		exGroup.GET("/getDataC", example.GetDataC)
		exGroup.GET("/getDataD", example.GetDataD)
		exGroup.GET("/startPage", example.StartPage)
		exGroup.GET("/bindUri/:id/:name", example.BindUri)
		exGroup.POST("/post", example.Post)
		exGroup.POST("/login", example.Login)
		exGroup.POST("/postForm", example.PostForm)
		exGroup.POST("/postAndQuery", example.PostAndQuery)
		exGroup.GET("/redirect", func(context *gin.Context) {
			context.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
		})
		exGroup.POST("/redirectForPost", func(context *gin.Context) {
			context.Redirect(http.StatusFound, "/foo")
		})
		exGroup.GET("/someJson", func(context *gin.Context) {
			name := []string{"len", "foo", "austin"}
			context.SecureJSON(http.StatusOK, name)
		})
		exGroup.GET("/cookie", func(context *gin.Context) {
			cookie, err := context.Cookie("gin_cookie")

			if err != nil {
				cookie = "NotSet"
				context.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
			}

			fmt.Printf("Cookie value: %s \n", cookie)
		})
		exGroup.POST("/uploads", func(context *gin.Context) {
			// Multipart form
			form, _ := context.MultipartForm()
			files := form.File["upload[]"]

			for _, file := range files {
				log.Println(file.Filename)

				// Upload the file to specific dst.
				context.SaveUploadedFile(file, "upload-to-this-path")
			}
			context.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
		})
		exGroup.POST("/upload", func(context *gin.Context) {
			// single file
			file, _ := context.FormFile("file")
			log.Println(file.Filename)

			// Upload the file to specific dst.
			context.SaveUploadedFile(file, "dst")

			context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		})
	}
}
