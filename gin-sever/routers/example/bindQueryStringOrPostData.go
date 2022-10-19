package example

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func StartPage(c *gin.Context) {
	var person Person
	if c.ShouldBind(&person) == nil {
		log.Printf("person name: %s\n", person.Name)
		log.Printf("person address: %s\n", person.Address)
		log.Printf("person birthday: %s\n", person.Birthday)
	}

	c.String(200, "ok")
}
