package server

import (
	"context"
	pb "demo/proto"

	"github.com/rs/zerolog/log"
)

func (s *demoService) Ping(ctx context.Context, request *pb.PingRequest) (*pb.PingReply, error) {
	log.Info().Msg("Ping get request")
	return &pb.PingReply{}, nil
}
