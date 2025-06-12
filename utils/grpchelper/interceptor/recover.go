package interceptor

import (
	"errors"
	"fmt"
	"runtime/debug"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ServerRecovery() grpc.UnaryServerInterceptor {
	return grpc_recovery.UnaryServerInterceptor(
		grpc_recovery.WithRecoveryHandler(
			func(p interface{}) (err error) {
				msg := fmt.Sprintf("panic: %v", p)
				stack := string(debug.Stack())

				if pe, ok := p.(error); ok {
					log.Error().Err(pe).Str("stack", stack).Msg("server panic")
				} else {
					log.Error().Err(errors.New("unformat panic error")).Str("stack", stack).Msg("server panic")
				}

				return status.New(codes.Internal, msg).Err()
			},
		),
	)
}
