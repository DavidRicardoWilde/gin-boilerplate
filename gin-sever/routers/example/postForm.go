package example

import "github.com/gin-gonic/gin"

func PostForm(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}
