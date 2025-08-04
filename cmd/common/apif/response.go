package apif

import (
	"elotus/cmd/common/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/modern-go/reflect2"
	"net/http"
)

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

func Respond(c *gin.Context, data interface{}, err error) {
	meta := Meta{
		Code:    http.StatusOK * 1000,
		Message: "success",
	}
	statusCode := http.StatusOK

	if err != nil {
		var customErr *errs.CustomError
		switch {
		case errors.As(err, &customErr):
			meta = Meta{
				Code:    customErr.Code,
				Message: customErr.Message,
			}
			statusCode = customErr.StatusCode
		default:
			meta = Meta{
				Code:    errs.ErrInternalServer,
				Message: errs.ErrService[errs.ErrInternalServer],
			}
			statusCode = http.StatusInternalServerError
		}
	}

	resp := Response{
		Meta: meta,
	}
	if !reflect2.IsNil(data) {
		resp.Data = data
	}
	c.JSON(statusCode, resp)
}
