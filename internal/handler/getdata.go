package handler

import (
	"github.com/gin-gonic/gin"
)

func (h TestHandler) getData(c *gin.Context) {
	id := c.Param("id")
	user := h.testService.GetData(id)

	c.JSON(200, gin.H{"message": user})
}
