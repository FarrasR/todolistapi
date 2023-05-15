package request

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryHelper struct {
	c *gin.Context
}

func NewQueryHelper(c *gin.Context) *QueryHelper {
	return &QueryHelper{c}
}

func (h *QueryHelper) GetInt(key string, defValue int) int {

	value, ok := h.c.GetQuery(key)

	if !ok {
		return defValue
	}

	if value != "" {
		if v, err := strconv.Atoi(value); err == nil {
			return v
		}
	}

	return defValue
}

func (h *QueryHelper) GetString(key string, defValue string) string {
	value, ok := h.c.GetQuery(key)

	if !ok {
		return defValue
	}

	return value
}

type RequestCreateActivity struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type RequestUpdateActivity struct {
	Title string `json:"title"`
}

type RequestCreateTodo struct {
	ActivityGroupId int    `json:"activity_group_id"`
	Title           string `json:"title"`
}

type RequestUpdateTodo struct {
	Title    string `json:"title"`
	IsActive bool   `json:"IsActive"`
	Priority string `json:"priority"`
}
