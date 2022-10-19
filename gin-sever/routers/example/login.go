package example

import "github.com/gin-gonic/gin"

type LoginForm struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var loginForm LoginForm

	if c.ShouldBind(&loginForm) == nil {
		if loginForm.User == "user" && loginForm.Password == "password" {
			c.JSON(200, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	}
}
