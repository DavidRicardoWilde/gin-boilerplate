package example

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AsciiJson(c *gin.Context) {
	data := map[string]interface{}{
		"lang": "Go语言",
		"tag":  "<br>",
	}

	c.AsciiJSON(http.StatusOK, data)
}
