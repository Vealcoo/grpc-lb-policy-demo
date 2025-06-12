package server

import (
	"errors"

	"google.golang.org/grpc/codes"
)

var (
	errorMap = map[error]codes.Code{
		ErrInvalidArgument: codes.InvalidArgument,
	}
	ErrInvalidArgument = errors.New("InvalidArgument")
)
