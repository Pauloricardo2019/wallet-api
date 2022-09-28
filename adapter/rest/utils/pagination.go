package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wallet-api/internal/model"
)

// @Security ApiKeyAuth

func GeneratePaginationFromRequest(c *gin.Context) model.Pagination {
	var limit, page int
	var sort string
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break

		}
	}
	return model.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

}
