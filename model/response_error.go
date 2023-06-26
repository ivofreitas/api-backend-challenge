package model

import (
	"fmt"
	"net/http"
)

const (
	subPathNotFound = "code=404, message=Not Found"
)

// ResponseError -
type ResponseError struct {
	DeveloperMessage string `json:"developer_message,omitempty"`
	UserMessage      string `json:"user_message,omitempty"`
	StatusCode       int    `json:"status_code,omitempty"`
}

// NewResponseError - Custom errors constructor
func NewResponseError(developerMessage string, userMessage string, statusCode int) *ResponseError {
	return &ResponseError{
		DeveloperMessage: developerMessage,
		UserMessage:      userMessage,
		StatusCode:       statusCode,
	}
}
func (me ResponseError) Error() string {
	return fmt.Sprintf("Error: Status %v and Message %s", me.StatusCode, me.DeveloperMessage)
}

// Conflict - Error for status 409
type Conflict ResponseError

func (me Conflict) Error() string {
	return fmt.Sprintf("Error: Status %v and Message %s", me.StatusCode, me.DeveloperMessage)
}

// NotFound - Error for status 404
type NotFound ResponseError

func (me NotFound) Error() string {
	return fmt.Sprintf("Error: Status %v and Message %s", me.StatusCode, me.DeveloperMessage)
}

// BadRequest - Error for status 400
type BadRequest ResponseError

func (me BadRequest) Error() string {
	return fmt.Sprintf("Error: Status %v and Message %s", me.StatusCode, me.DeveloperMessage)
}

// Unauthorized - Error for status 401
type Unauthorized ResponseError

func (me Unauthorized) Error() string {
	return fmt.Sprintf("Error: Status %v and Message %s", me.StatusCode, me.DeveloperMessage)
}

type Forbidden ResponseError

func (me Forbidden) Error() string {
	return fmt.Sprintf("Error: Status %v and Message %s", me.StatusCode, me.DeveloperMessage)
}

// Custom - Custom errors
type Custom ResponseError

func (me Custom) Error() string {
	return fmt.Sprintf("Error: Status %v and Message %s", me.StatusCode, me.DeveloperMessage)
}

// ErrorDiscover - Error Constructor
func ErrorDiscover(i interface{}) *ResponseError {
	var (
		developerMessage string
		userMessage      string
		statusCode       int
	)
	switch e := i.(type) {
	case NotFound:
		developerMessage = e.DeveloperMessage
		userMessage = "Resource not found"
		statusCode = http.StatusNotFound

	case BadRequest:
		developerMessage = e.DeveloperMessage
		userMessage = "Malformed request"
		statusCode = http.StatusBadRequest

	case Conflict:
		developerMessage = e.DeveloperMessage
		userMessage = "Register already exists"
		statusCode = http.StatusConflict

	case Unauthorized:
		developerMessage = e.DeveloperMessage
		userMessage = "You are not authorized to perform this operation"
		statusCode = http.StatusUnauthorized

	case Forbidden:
		developerMessage = "Forbidden - access denied"
		userMessage = "Make sure you have access to this resource"
		statusCode = http.StatusForbidden

	case Custom:
		developerMessage = e.DeveloperMessage
		userMessage = e.UserMessage
		statusCode = e.StatusCode

	case error:

		switch e.Error() {
		case subPathNotFound:
			developerMessage = "Unknown subpath"
			userMessage = "Resource not found"
			statusCode = http.StatusNotFound
		default:
			developerMessage = e.Error()
			userMessage = "Was encountered an errors when processing your request. We apologize for the inconvenience."
			statusCode = http.StatusInternalServerError
		}
	}
	return NewResponseError(developerMessage, userMessage, statusCode)
}
