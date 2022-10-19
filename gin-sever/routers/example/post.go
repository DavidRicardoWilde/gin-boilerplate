package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")

	fmt.Printf("ids: %v; names: %v", ids, names)
}
