package example

import "github.com/gin-gonic/gin"

func PostAndQuery(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}
