package core

import (
	"log"
	"net/http"

	"app/pkg/messages"
	"app/pkg/validation"

	"github.com/gin-gonic/gin"
)

// Controller is stateless: response values are always passed per call, never
// stored on the controller, so it is safe for concurrent request handlers.
type Controller struct{}

var controller = &Controller{}

func GetController() *Controller {
	return controller
}

func ToResponse(message string, data any, errors any) map[string]any {
	return gin.H{
		"message": message,
		"errors":  errors,
		"data":    data,
	}
}

func (ctrl *Controller) Success(c *gin.Context, data any) {
	if data == nil {
		data = map[string]any{}
	}

	c.SecureJSON(http.StatusOK, ToResponse(messages.MessageOk, data, map[string]any{}))
}

func (ctrl *Controller) UpstreamError(c *gin.Context, err error) {
	log.Printf("upstream request failed: %v", err)

	c.SecureJSON(
		http.StatusBadGateway,
		ToResponse(messages.MessageBadGateway, map[string]any{}, map[string]any{}),
	)
}

func (ctrl *Controller) JsonBindError(c *gin.Context, err error) {
	response := ToResponse(
		messages.MessageUnprocessable,
		map[string]any{},
		validation.Handle(err),
	)
	c.SecureJSON(http.StatusUnprocessableEntity, response)
}

func (ctrl *Controller) UnprocessableError(c *gin.Context, errs map[string]string) {
	response := ToResponse(
		messages.MessageUnprocessable,
		map[string]any{},
		errs,
	)
	c.SecureJSON(http.StatusUnprocessableEntity, response)
}

func (ctrl *Controller) NotFoundError(c *gin.Context, errs map[string]string) {
	response := ToResponse(
		messages.MessageNotFound,
		map[string]any{},
		errs,
	)
	c.SecureJSON(http.StatusNotFound, response)
}

func (ctrl *Controller) BadRequestError(c *gin.Context, errs map[string]string) {
	response := ToResponse(
		messages.MessageBadRequest,
		map[string]any{},
		errs,
	)
	c.SecureJSON(http.StatusBadRequest, response)
}
