package shared

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Response struct {
	StatusCode int
	Data       interface{} `json:"data"`
}

type Error struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func (res *Response) BuildSuccess(data interface{}, status ...int) {
	if len(status) == 0 {
		status = append(status, http.StatusOK)
	}
	res.StatusCode, res.Data = status[0], data
}

func (res *Response) BuildError(err error, status ...int) {
	if len(status) == 0 {
		status = append(status, http.StatusInternalServerError)
	}

	switch err := err.(type) {
	default:
		res.buildError(status[0], "Error", err)
	case validator.ValidationErrors:
		res.buildValidationError(http.StatusUnprocessableEntity, "Validation error", err)
	case error:
		switch err {
		default:
			res.buildError(status[0], "Error", err)
		case gorm.ErrRecordNotFound:
			res.buildError(http.StatusNotFound, "Record not found", err)
		}
	}
}

func (res *Response) buildError(status int, message string, e error) {
	err := Error{
		Message: message,
		Errors:  e.Error(),
	}
	res.StatusCode, res.Data = status, err
}

func (res *Response) buildValidationError(status int, message string, errors validator.ValidationErrors) {
	err := Error{
		Message: message,
		Errors:  mapValidationErrors(errors),
	}
	res.StatusCode, res.Data = status, err
}

func mapValidationErrors(errors validator.ValidationErrors) map[string]string {
	result := make(map[string]string)
	for _, err := range errors {
		result[err.Field()] = customErrorsByTag(err.Field(), err.ActualTag())
	}
	return result
}

func customErrorsByTag(field, tag string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("Field %s is required", field)
	case "numeric":
		return fmt.Sprintf("Field %s must be a numeric type", field)
	default:
		return fmt.Sprintf("Field %s is invalid", field)
	}
}
