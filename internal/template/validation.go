package template

func ValidationRoot() string {
	return `package validation

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type APIError struct {
	Code    string            ` + "`json:\"code\"`" + `
	Message string            ` + "`json:\"message\"`" + `
	Details map[string]string ` + "`json:\"details,omitempty\"`" + `
}

var (
	ErrValidation = errors.New("validation error")
	ErrNotFound   = errors.New("not found")
	ErrConflict   = errors.New("conflict")
)

func ErrorHandler(err error, c echo.Context) {
	fmt.Printf("error type: %T\n", err)

	status := http.StatusInternalServerError
	resp := APIError{
		Code:    "INTERNAL_ERROR",
		Message: "Internal server error",
	}

	// ---- Validation errors ----
	if ve, ok := err.(validator.ValidationErrors); ok {
		status = http.StatusBadRequest
		resp.Code = "VALIDATION_ERROR"
		resp.Message = "Validation failed"
		resp.Details = make(map[string]string)

		for _, fe := range ve {
			resp.Details[fe.Field()] = validationMessage(fe)
		}
	}

	// ---- Echo HTTP errors ----
	if he, ok := err.(*echo.HTTPError); ok {
		status = he.Code
		resp.Code = "HTTP_ERROR"
		resp.Message = he.Message.(string)
	}

	// ---- Domain errors ----
	switch err {
	case ErrNotFound:
		status = http.StatusNotFound
		resp.Code = "NOT_FOUND"
		resp.Message = "Resource not found"

	case ErrConflict:
		status = http.StatusConflict
		resp.Code = "CONFLICT"
		resp.Message = "Resource already exists"
	}

	// ---- Send response ----
	if !c.Response().Committed {
		_ = c.JSON(status, map[string]any{
			"error": resp,
		})
	}
}

func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is required"
	case "min":
		return "is too short"
	case "max":
		return "is too long"
	case "email":
		return "must be a valid email"
	default:
		return "is invalid"
	}
}
`
}
