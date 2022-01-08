package log

import (
	"os"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func InitLog() {
	log.Logger = logger
}

func Error(err error) {
	log.Logger.Error().Err(err)
}
