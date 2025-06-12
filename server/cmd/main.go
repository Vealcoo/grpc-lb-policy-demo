package main

import (
	"demo/server"
	"demo/utils/loghelper"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func main() {
	loghelper.InitLogger("demo_server")

	grpcServer := server.New()
	go grpcServer.Start()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sc)

	log.Info().Msg("Service starting...")
	<-sc
	grpcServer.Shutdown()
}
