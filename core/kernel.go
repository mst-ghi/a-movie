package core

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func PaginateQueries(c *gin.Context) (string, int) {
	filter := c.DefaultQuery("filter", "created")

	if filter == "" {
		filter = "created"
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	if page < 0 {
		page = 0
	}

	return filter, page
}
