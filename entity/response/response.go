package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	helloWorld     = "Hello World"
	successMessage = "Success"

	failedMessage = "Failed"

	errorInvalidParameter = "Invalid parameter"
)

type ResponseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type EmptyStruct struct{}

func buildResponse(status string, message string, data any) ResponseBody {
	return ResponseBody{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func Error(c *gin.Context, httpCode int, status string, message string) {
	c.JSON(httpCode, buildResponse(status, message, EmptyStruct{}))
	c.Abort()
}

func OKHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, buildResponse(successMessage, helloWorld, nil))
}

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, buildResponse(successMessage, successMessage, data))
}

func OKCreated(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, buildResponse(successMessage, successMessage, data))
}

func ErrorInvalidParameter(c *gin.Context) {
	c.JSON(http.StatusBadRequest, buildResponse(failedMessage, errorInvalidParameter, nil))
	c.Abort()
}
