<<<<<<<< HEAD:packages/model/error.go
package model
========
package core
>>>>>>>> slhmy/user_service2:packages/core/error.go

import (
	"fmt"
	"runtime"
)

type SeviceError struct {
	Code       int       `json:"code"`
	Msg        string    `json:"msg"`
	stackTrace []uintptr `json:"-"`
}

func (se *SeviceError) CaptureStackTrace() *SeviceError {
	se.stackTrace = []uintptr{}
	runtime.Callers(2, se.stackTrace)

	return se
}

func IsServiceError(err interface{}) bool {
	_, ok := err.(*SeviceError)
	return ok
}

func WrapToServiceError(err interface{}) *SeviceError {
	if IsServiceError(err) {
		return err.(*SeviceError)
	} else {
		serviceErr := NewInternalError(fmt.Sprintf("%v", err))
		serviceErr.CaptureStackTrace()
		return serviceErr
	}
}

func NewInternalError(msg string) *SeviceError {
	return &SeviceError{
		Code: 500,
		Msg:  msg,
	}
}

func NewUnAuthorizedError(msg string) *SeviceError {
	return &SeviceError{
		Code: 401,
		Msg:  msg,
	}
}
