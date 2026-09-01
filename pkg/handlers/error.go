package handlers

import (
	"log"

	"app/core"
	"app/pkg/messages"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
)

func InternalErrorHandler(c *gin.Context, err any) {
	goErr := errors.Wrap(err, 2)
	log.Printf("internal error: %v\n%s", goErr, goErr.ErrorStack())

	c.AbortWithStatusJSON(
		500,
		core.ToResponse(
			messages.MessageInternalError,
			map[string]any{},
			map[string]any{},
		),
	)
}
