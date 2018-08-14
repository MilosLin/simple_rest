package demo

import (
	"log"
	"net/http"
	"simple_rest/api/protocol"
	"time"

	"github.com/gin-gonic/gin"
)

// PostingPersonInput : Input
type PostingPersonInput struct {
	Colors   []string  `form:"Colors[]"`
	Name     string    `form:"Name" binding:"required"`
	Address  string    `form:"Address"`
	Birthday time.Time `form:"Birthday" time_format:"2006-01-02T15:04:05Z07:00"`
}

// PostingPersonOutput : Output
type PostingPersonOutput struct {
	Persion PostingPersonInput
}

// Postting : Post API
func Postting(c *gin.Context) {
	res := protocol.Response{}
	var person PostingPersonInput
	res.Result = &person

	if err := c.Bind(&person); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
