package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"todolistapi/domain"
	"todolistapi/entity/request"
	"todolistapi/entity/response"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	TodoUsecase domain.TodoUsecase
}

func NewTodoHandler(todoUsecase domain.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		TodoUsecase: todoUsecase,
	}
}

func (h *TodoHandler) Register(router *gin.Engine) {
	router.GET("/todo-items", h.GetAllTodos)
	router.GET("/todo-items/:id", h.GetTodoById)
	router.POST("/todo-items", h.CreateTodo)
	router.DELETE("/todo-items/:id", h.DeleteTodo)
	router.PATCH("/todo-items/:id", h.UpdateTodo)
}

func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	qh := request.NewQueryHelper(c)
	activityGroupID := qh.GetInt("activity_group_id", 0)

	if activityGroupID == 0 {
		todos, err := h.TodoUsecase.GetAllTodos()
		if err != nil {
			response.Error(c, http.StatusBadRequest, "failed", err.Error())
			return
		}
		response.OK(c, todos)
		return
	}

	todos, err := h.TodoUsecase.GetAllTodosByActivityId(uint(activityGroupID))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed", err.Error())
		return
	}
	response.OK(c, todos)
}

func (h *TodoHandler) GetTodoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	todo, err := h.TodoUsecase.GetOneTodo(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			response.Error(c, http.StatusNotFound, "Not Found", fmt.Sprintf("Todo with ID %d Not Found", id))
			return
		}
		response.Error(c, http.StatusBadRequest, "failed", err.Error())
		return
	}
	response.OK(c, todo)
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var request request.RequestCreateTodo

	if err := c.ShouldBindJSON(&request); err != nil {
		response.ErrorInvalidParameter(c)
		return
	}
	if request.ActivityGroupId == 0 {
		response.Error(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "activity_group_id cannot be null")
		return
	}

	if request.Title == "" {
		response.Error(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "title cannot be null")
		return
	}

	todo, err := h.TodoUsecase.CreateTodo(uint(request.ActivityGroupId), request.Title)

	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed", err.Error())
		return
	}
	response.OK(c, todo)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorInvalidParameter(c)
		return
	}
	err = h.TodoUsecase.DeleteTodo(uint(id))

	if err != nil {
		if err.Error() == "record not found" {
			response.Error(c, http.StatusNotFound, "Not Found", fmt.Sprintf("Todo with ID %d Not Found", id))
			return
		}
		response.ErrorInvalidParameter(c)
		return
	}

	response.OK(c, response.EmptyStruct{})
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	var request request.RequestUpdateTodo
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	if request.Title == "" {
		response.Error(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "title cannot be null")
		return
	}
	todo, err := h.TodoUsecase.UpdateTodo(uint(id), request.Title)
	if err != nil {
		if err.Error() == "record not found" {
			response.Error(c, http.StatusNotFound, "Not Found", fmt.Sprintf("Todo with ID %d Not Found", id))
			return
		}
		response.Error(c, http.StatusBadRequest, "failed", err.Error())
		return
	}
	response.OK(c, todo)
}
