package example

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Project struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func BindUri(c *gin.Context) {
	var project Project
	if err := c.ShouldBindUri(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"project": project})
}
