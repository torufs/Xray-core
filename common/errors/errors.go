// Package errors provides error handling utilities for Xray-core.
// It wraps standard errors with additional context such as severity levels
// and call path information for better debugging and logging.
package errors

import (
	"errors"
	"fmt"
	"strings"
)

// Severity represents the severity level of an error.
type Severity byte

const (
	// SeverityDebug indicates a debug-level error (informational).
	SeverityDebug Severity = iota
	// SeverityInfo indicates an informational error.
	SeverityInfo
	// SeverityWarning indicates a warning-level error.
	SeverityWarning
	// SeverityError indicates a standard error.
	SeverityError
)

// Error is an error type with additional context for Xray-core.
type Error struct {
	pathObj  interface{}
	message  []interface{}
	inner    error
	severity Severity
}

// Error implements the error interface.
func (e *Error) Error() string {
	builder := strings.Builder{}
	if e.pathObj != nil {
		builder.WriteString(fmt.Sprintf("%T", e.pathObj))
		builder.WriteString(": ")
	}
	msgParts := make([]string, 0, len(e.message))
	for _, msg := range e.message {
		msgParts = append(msgParts, fmt.Sprintf("%v", msg))
	}
	builder.WriteString(strings.Join(msgParts, " "))
	if e.inner != nil {
		builder.WriteString(" | caused by: ")
		builder.WriteString(e.inner.Error())
	}
	return builder.String()
}

// Unwrap returns the inner error, implementing the errors.Unwrap interface.
func (e *Error) Unwrap() error {
	return e.inner
}

// Severity returns the severity level of this error.
func (e *Error) GetSeverity() Severity {
	return e.severity
}

// WithSeverity sets the severity level of the error and returns it.
func (e *Error) WithSeverity(s Severity) *Error {
	e.severity = s
	return e
}

// AtDebug sets the error severity to debug.
func (e *Error) AtDebug() *Error {
	return e.WithSeverity(SeverityDebug)
}

// AtInfo sets the error severity to info.
func (e *Error) AtInfo() *Error {
	return e.WithSeverity(SeverityInfo)
}

// AtWarning sets the error severity to warning.
func (e *Error) AtWarning() *Error {
	return e.WithSeverity(SeverityWarning)
}

// AtError sets the error severity to error (default).
func (e *Error) AtError() *Error {
	return e.WithSeverity(SeverityError)
}

// New creates a new Error with the given message parts.
func New(msg ...interface{}) *Error {
	return &Error{
		message:  msg,
		severity: SeverityError,
	}
}

// Cause wraps an existing error with additional context message parts.
func Cause(inner error, msg ...interface{}) *Error {
	return &Error{
		message:  msg,
		inner:    inner,
		severity: SeverityError,
	}
}

// Is reports whether any error in the chain matches the target.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As finds the first error in the chain that matches target.
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// GetSeverityFromError returns the severity of an error if it is an *Error,
// otherwise returns SeverityError as the default.
func GetSeverityFromError(err error) Severity {
	var e *Error
	if errors.As(err, &e) {
		return e.GetSeverity()
	}
	return SeverityError
}
