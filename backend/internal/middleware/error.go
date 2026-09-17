package middleware

import (
	"errors"
	"knowledge/internal/apperror"
	"knowledge/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		var appErr *apperror.AppError

		if errors.As(err, &appErr) {

			response.Error(
				c,
				appErr.HTTPStatus,
				appErr.Code,
				appErr.Message,
			)

			return
		}

		// 未知错误
		response.Error(
			c,
			http.StatusInternalServerError,
			50001,
			"internal server error",
		)
	}
}