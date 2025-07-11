package interceptor

import (
	"context"

	"demo/utils/grpchelper"

	"google.golang.org/grpc"
)

func ServerError(errorHandler *grpchelper.ErrorHandler) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (any, error) {
		res, err := handler(ctx, req)

		if err != nil {
			err = errorHandler.ErrToCode(err)
			return res, err
		}

		return res, err
	}
}
