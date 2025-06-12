package grpchelper

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorHandler struct {
	errMap map[error]codes.Code
}

func NewGRPCErrorHandler() *ErrorHandler {
	return &ErrorHandler{
		errMap: make(map[error]codes.Code),
	}
}

func (h *ErrorHandler) SetErrorMap(m map[error]codes.Code) {
	merge := make(map[error]codes.Code)

	for key, value := range h.errMap {
		merge[key] = value
	}

	for key, value := range m {
		merge[key] = value
	}

	h.errMap = merge
}

func (h *ErrorHandler) ErrToCode(err error) error {
	v, ok := h.errMap[err]
	if !ok {
		v = codes.Internal
	}

	if _, ok := status.FromError(err); ok {
		return err
	}
	return status.Errorf(v, "%s", err.Error())
}

func IsErrNotFound(err error) bool {
	if s, ok := status.FromError(err); ok && s.Code() == codes.NotFound {
		return true
	}

	return false
}

func IsErrInternal(err error) bool {
	if s, ok := status.FromError(err); ok && s.Code() == codes.Internal {
		return true
	}

	return false
}

func IsErrResourceExhausted(err error) bool {
	if s, ok := status.FromError(err); ok && s.Code() == codes.ResourceExhausted {
		return true
	}

	return false
}
