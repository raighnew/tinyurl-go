package app

import (
	"log"
	"net/http"

	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"

	"github.com/raighneweng/tinyurl-go/pkg/e"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.Println(err.Key, err.Message)
	}

	return
}

func BindAndValid(c *gin.Context, params interface{}) (int, int) {
	err := c.Bind(params)

	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(params)

	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}

	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
