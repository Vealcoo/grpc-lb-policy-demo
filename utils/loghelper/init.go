package loghelper

import (
	"fmt"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger(serviceName string) zerolog.Logger {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}
	log.Logger = log.With().Str("service", serviceName).Caller().Logger()
	return log.Logger
}
