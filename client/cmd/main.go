package main

import (
	"context"
	"demo/client/conn"
	"demo/utils/loghelper"
	"demo/utils/randhelper"
	"time"

	"github.com/rs/zerolog/log"

	pb "demo/proto"
)

func main() {
	loghelper.InitLogger("demo_client")

	demoClient, err := conn.DemoClient("demo-server:8082")
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	var round int64
	ctx := context.Background()
	for {
		start := time.Now()
		randNum := randhelper.GenerateRandomNumber(1, 100)
		switch {
		case randNum <= 40:
			log.Info().Msg("Ping")
			_, err := demoClient.Ping(ctx, &pb.PingRequest{})
			if err != nil {
				log.Error().Err(err).Str("method", "Ping").Send()
			}
		case randNum > 40 && randNum <= 50:
			log.Info().Msg("Panic")
			_, err := demoClient.Panic(ctx, &pb.PanicRequest{})
			if err != nil {
				log.Error().Err(err).Str("method", "Panic").Send()
			}
		case randNum > 50:
			log.Info().Msg("Retry")
			_, err := demoClient.Retry(ctx, &pb.RetryRequest{})
			if err != nil {
				log.Error().Err(err).Str("method", "Retry").Send()
			}
		default:
		}
		end := time.Now()
		round++
		log.Info().TimeDiff("cost_time", end, start).Int64("round", round).Msg("send request")
		time.Sleep(5000 * time.Millisecond)
	}
}
