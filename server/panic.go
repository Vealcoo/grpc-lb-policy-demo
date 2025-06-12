package server

import (
	"context"
	pb "demo/proto"

	"github.com/rs/zerolog/log"
)

func (s *demoService) Panic(ctx context.Context, request *pb.PanicRequest) (*pb.PanicReply, error) {
	log.Info().Msg("Panic get request")
	panic("panic for interceptor test")
}
