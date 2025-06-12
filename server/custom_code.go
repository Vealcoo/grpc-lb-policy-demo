package server

import (
	"context"
	pb "demo/proto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *demoService) CustomCode(ctx context.Context, request *pb.CustomCodeRequest) (*pb.CustomCodeReply, error) {
	log.Info().Msg("CustomCode get request")

	if codes.Code(request.Code) == codes.OK {
		return &pb.CustomCodeReply{}, nil
	}

	return nil, status.New(codes.Code(request.Code), request.Message).Err()
}
