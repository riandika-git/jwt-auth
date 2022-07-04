package helper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type Errors struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type PaginationResponse struct {
	Pagination Pagination  `json:"pagination"`
	Content    interface{} `json:"data"`
}

//EmptyObj object is used when data doesnt want to be null on json
type EmptyObj struct{}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

//BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}

func ResponseSuccess(data interface{}, ctx *gin.Context) {
	res := BuildResponse(true, "success", data)
	ctx.JSON(http.StatusOK, res)
}

func DialogSuccess(message string, ctx *gin.Context) {
	res := BuildResponse(true, message, nil)
	ctx.JSON(http.StatusOK, res)
}

func DialogError(message string, httpStatus int, ctx *gin.Context) {
	res := BuildResponse(true, message, nil)
	//ctx.JSON(httpStatus, res)
	ctx.AbortWithStatusJSON(400, res)
	panic(nil)
}

func ErrorValidation(err error, ctx *gin.Context) {
	fields := make([]Errors, 0)
	for _, err := range err.(validator.ValidationErrors) {
		errV := Errors{
			ID:      err.Field(),
			Message: ErrorMessageValidation(err),
		}
		fields = append(fields, errV)
	}
	res := Response{
		Status:  false,
		Message: "Error validation",
		Errors:  fields,
		Data:    nil,
	}
	ctx.AbortWithStatusJSON(400, res)
	panic(nil)
}
func ErrorMessageValidation(fe validator.FieldError) string {
	fieldErrMsg := "Field validation for '%s' failed on the '%s' tag"
	return fmt.Sprintf(fieldErrMsg, fe.Field(), fe.Tag())
}

func BuildPaginationResponse(dto interface{}, pagination Pagination) PaginationResponse {
	return PaginationResponse{
		Pagination: pagination,
		Content:    dto,
	}
}
