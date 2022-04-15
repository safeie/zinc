package storage

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type logger struct{}

var Logger = new(logger)

func init() {
	log.Level(zerolog.ErrorLevel)
}

func (t *logger) Errorf(format string, v ...interface{}) {
	log.Error().Str("module", "badger").Msgf(format, v)
}

func (t *logger) Warningf(format string, v ...interface{}) {}

func (t *logger) Infof(format string, v ...interface{}) {}

func (t *logger) Debugf(format string, v ...interface{}) {}
