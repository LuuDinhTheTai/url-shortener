package response

import (
	"time"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Meta struct {
	RequestId string    `json:"request_id"`
	Timestamp time.Time `json:"timestamp"`
}

type ApiResponse struct {
	Status  string          `json:"status"`
	Code    int             `json:"code"`
	Message string          `json:"message,omitempty"`
	Data    any             `json:"data,omitempty"`
	Errors  []ErrorResponse `json:"errors,omitempty"`
	Meta    Meta            `json:"meta,omitempty"`
}

func Success(c *gin.Context, code int, message string, data any) {
	c.JSON(code,
		ApiResponse{
			Status:  "success",
			Code:    code,
			Message: message,
			Data:    data,
			Meta: Meta{
				RequestId: c.GetHeader("X-Request-Id"),
				Timestamp: time.Now(),
			},
		})
}

func Error(c *gin.Context, code int, message string, errors []ErrorResponse) {
	c.JSON(code,
		ApiResponse{
			Status:  "error",
			Code:    code,
			Message: message,
			Errors:  errors,
			Meta: Meta{
				RequestId: c.GetHeader("X-Request-Id"),
				Timestamp: time.Now(),
			},
		})
}
