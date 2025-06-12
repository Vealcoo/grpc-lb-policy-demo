package interceptor

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func ServerLogging(isDebug bool) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		res, err := handler(ctx, req)

		if err != nil {
			log.Error().Err(err).Str("api", info.FullMethod).Interface("req", req).Interface("res", res).Msg("server return error")
		} else if isDebug {
			log.Debug().Str("api", info.FullMethod).Interface("req", req).Interface("res", res).Msg("server return success")
		}

		return res, err
	}
}
