package server

import (
	"context"
	pb "demo/proto"
	"sync"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var retryMu sync.Mutex
var retryCount int64

func (s *demoService) Retry(ctx context.Context, request *pb.RetryRequest) (*pb.RetryReply, error) {
	retryMu.Lock()
	defer retryMu.Unlock()

	log.Info().Int64("retry_count", retryCount).Msg("Retry get request")

	if retryCount += 1; retryCount >= 2 {
		retryCount = 0
		log.Info().Msg("Retry count reset, request succeeded")
		return &pb.RetryReply{}, nil
	}

	return &pb.RetryReply{}, status.New(codes.Unavailable, "for client retry testing").Err()
}
