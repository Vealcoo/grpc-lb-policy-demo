package server

import (
	"context"
	pb "demo/proto"
	"sync/atomic"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var retryCount atomic.Int64

func (s *demoService) Retry(ctx context.Context, request *pb.RetryRequest) (*pb.RetryReply, error) {
	log.Info().Int64("retry_count", retryCount.Load()).Msg("Retry get request")

	if retryCount.Add(1) == 2 {
		retryCount.Store(0)
		log.Info().Msg("Retry count reset, request succeeded")
		return &pb.RetryReply{}, nil
	}

	return &pb.RetryReply{}, status.New(codes.Unavailable, "for client retry testing").Err()
}
