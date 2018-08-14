package demo

import (
	"log"
	"net/http"
	"simple_rest/api/protocol"
	"time"

	"github.com/gin-gonic/gin"
)

// GettingPersonInput : Input參數
type GettingPersonInput struct {
	Colors   []string  `form:"Colors[]"`
	Name     string    `form:"Name" binding:"required"`
	Address  string    `form:"Address"`
	Birthday time.Time `form:"Birthday" time_format:"2006-01-02T15:04:05Z07:00"`
}

// Getting API
func Getting(c *gin.Context) {
	res := protocol.Response{}
	var person GettingPersonInput
	res.Result = &person

	// 綁定Input參數至結構中
	if err := c.Bind(&person); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
