package errs

import (
	"github.com/bombergame/common/consts"
)

//ErrorType is enum constant
type ErrorType int

const (
	//Internal is a corresponding error type
	Internal ErrorType = iota

	//NotFound is a corresponding error type
	NotFound

	//Duplicate is a corresponding error type
	Duplicate

	//InvalidFormat is a corresponding error type
	InvalidFormat

	//BadRequest is a corresponding error type
	BadRequest

	//NotAuthorized is a corresponding error type
	NotAuthorized

	//AccessDenied is a corresponding error type
	AccessDenied
)

const (
	//InternalServiceErrorMessage is the displayed error message
	InternalServiceErrorMessage = "internal service error"

	//NotFoundErrorMessagePrefix is the displayed message prefix
	NotFoundErrorMessagePrefix = "not found: "

	//InvalidFormatErrorMessagePrefix is the displayed message prefix
	InvalidFormatErrorMessagePrefix = "invalid format: "

	//DuplicateErrorMessagePrefix is the displayed message prefix
	DuplicateErrorMessagePrefix = "duplicate: "

	//BadRequestErrorMessagePrefix is the displayed message prefix
	BadRequestErrorMessagePrefix = "bad request: "

	//NotAuthorizedErrorMessage is the displayed message text
	NotAuthorizedErrorMessage = "not authorized"

	//AccessDeniedErrorMessage is the displayed message text
	AccessDeniedErrorMessage = "access denied"
)

//ServiceError contains error data
type ServiceError struct {
	message  string
	innerErr error
	errType  ErrorType
}

//NewInternalServiceError creates error instance
func NewInternalServiceError(err error) error {
	return &ServiceError{
		errType:  Internal,
		message:  InternalServiceErrorMessage,
		innerErr: err,
	}
}

//NewNotFoundError creates error instance
func NewNotFoundError(message string) error {
	return &ServiceError{
		errType: NotFound,
		message: NotFoundErrorMessagePrefix + message,
	}
}

//NewInvalidFormatError creates error instance
func NewInvalidFormatError(message string) error {
	return &ServiceError{
		errType: InvalidFormat,
		message: InvalidFormatErrorMessagePrefix + message,
	}
}

//NewDuplicateError creates error instance
func NewDuplicateError(message string) error {
	return &ServiceError{
		errType: Duplicate,
		message: DuplicateErrorMessagePrefix + message,
	}
}

//NewBadRequestError creates error instance
func NewBadRequestError(message string) error {
	return &ServiceError{
		errType: BadRequest,
		message: BadRequestErrorMessagePrefix + message,
	}
}

//NewNotAuthorizedError creates error instance
func NewNotAuthorizedError() error {
	return &ServiceError{
		errType: NotAuthorized,
		message: NotAuthorizedErrorMessage,
	}
}

//NewAccessDeniedError creates error instance
func NewAccessDeniedError() error {
	return &ServiceError{
		errType: AccessDenied,
		message: AccessDeniedErrorMessage,
	}
}

//Error returns string representation of the error
func (err *ServiceError) Error() string {
	return err.message
}

//InnerError returns string representation of the nested error
func (err *ServiceError) InnerError() string {
	if err.innerErr != nil {
		return err.innerErr.Error()
	}
	return consts.EmptyString
}

//ErrorType returns error type
func (err *ServiceError) ErrorType() ErrorType {
	return err.errType
}
