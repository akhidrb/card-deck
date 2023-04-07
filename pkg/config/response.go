package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleResponseError(c *gin.Context, err error) {
	switch cErr := err.(type) {
	case BadRequestError:
		c.JSON(http.StatusBadRequest, errorToDtoError(cErr))
	case NotFoundResourceError:
		c.JSON(http.StatusNotFound, errorToDtoError(cErr))
	case ValidationError:
		c.JSON(http.StatusForbidden, errorToDtoError(cErr))
	default:
		c.JSON(
			http.StatusInternalServerError, ErrorResponse{
				Message: "server error",
			},
		)
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func errorToDtoError(e error) ErrorResponse {
	return ErrorResponse{
		Message: e.Error(),
	}
}

type NotFoundResourceError struct {
	error
}

func NewNotFoundResourceError(err error) NotFoundResourceError {
	return NotFoundResourceError{
		error: err,
	}
}

type BadRequestError struct {
	error
}

func NewBadRequestError(err error) BadRequestError {
	return BadRequestError{
		error: err,
	}
}

type ValidationError struct {
	error
}

func NewValidationError(err error) ValidationError {
	return ValidationError{
		error: err,
	}
}
