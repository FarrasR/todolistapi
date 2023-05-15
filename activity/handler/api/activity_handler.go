package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"todolistapi/domain"
	"todolistapi/entity/model"
	"todolistapi/entity/request"
	"todolistapi/entity/response"

	"github.com/gin-gonic/gin"
)

type ActivityHandler struct {
	ActivityUsecase domain.ActivityUsecase
}

func NewActivityHandler(activityUsecase domain.ActivityUsecase) *ActivityHandler {
	return &ActivityHandler{
		ActivityUsecase: activityUsecase,
	}
}

func (h *ActivityHandler) Register(router *gin.Engine) {
	router.GET("/activity-groups", h.GetAllActivities)
	router.GET("/activity-groups/:id", h.GetActityById)
	router.POST("/activity-groups", h.CreateActivity)
	router.DELETE("/activity-groups/:id", h.DeleteActivity)
	router.PATCH("/activity-groups/:id", h.UpdateActivity)
}

func (h *ActivityHandler) GetAllActivities(c *gin.Context) {
	activities, err := h.ActivityUsecase.GetAllActivity()
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed", err.Error())
		return
	}

	response.OK(c, activities)
}

func (h *ActivityHandler) GetActityById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	activities, err := h.ActivityUsecase.GetOneActivity(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			response.Error(c, http.StatusNotFound, "Not Found", fmt.Sprintf("Activity with ID %d Not Found", id))
			return
		}
		response.Error(c, http.StatusBadRequest, "failed", err.Error())
		return
	}
	response.OK(c, activities)
}

func (h *ActivityHandler) CreateActivity(c *gin.Context) {
	var request request.RequestCreateActivity

	if err := c.ShouldBindJSON(&request); err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	if request.Title == "" {
		response.Error(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "title cannot be null")
		return
	}

	if request.Email == "" {
		response.Error(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "email cannot be null")
		return
	}

	activity := model.Activity{
		Email: request.Email,
		Title: request.Title,
	}

	result, err := h.ActivityUsecase.CreateActivity(activity)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed", err.Error())
		return
	}
	response.OKCreated(c, result)
}

func (h *ActivityHandler) DeleteActivity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	err = h.ActivityUsecase.DeleteActivity(uint(id))

	if err != nil {
		if err.Error() == "record not found" {
			response.Error(c, http.StatusNotFound, "Not Found", fmt.Sprintf("Activity with ID %d Not Found", id))
			return
		}
		response.ErrorInvalidParameter(c)
		return
	}

	response.OK(c, response.EmptyStruct{})
}

func (h *ActivityHandler) UpdateActivity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	var request request.RequestUpdateActivity

	if err := c.ShouldBindJSON(&request); err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	if request.Title == "" {
		response.Error(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "title cannot be null")
		return
	}

	activity, err := h.ActivityUsecase.UpdateActivity(uint(id), request.Title)
	if err != nil {
		if err.Error() == "record not found" {
			response.Error(c, http.StatusNotFound, "Not Found", fmt.Sprintf("Activity with ID %d Not Found", id))
			return
		}
		response.Error(c, http.StatusBadRequest, "failed", err.Error())
		return
	}
	response.OK(c, activity)
}
