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

	ctx := context.Background()
	for {
		randNum := randhelper.GenerateRandomNumber(1, 100)
		switch {
		case randNum <= 40:
			_, err := demoClient.Ping(ctx, &pb.PingRequest{})
			if err != nil {
				log.Error().Err(err).Str("method", "Ping").Send()
			}
		case randNum > 40 && randNum <= 50:
			_, err := demoClient.Panic(ctx, &pb.PanicRequest{})
			if err != nil {
				log.Error().Err(err).Str("method", "Panic").Send()
			}
		case randNum > 50:
			_, err := demoClient.Retry(ctx, &pb.RetryRequest{})
			if err != nil {
				log.Error().Err(err).Str("method", "Retry").Send()
			}

		default:
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
