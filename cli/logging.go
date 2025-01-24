package cli

import (
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

var initLogOnce sync.Once

func initLogging(logLevel string) {
	initLogOnce.Do(func() {
		log.SetFormatter(&log.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: time.TimeOnly,
		})

		log.SetOutput(os.Stdout)
		setLogLevel(logLevel)
	})
}

func setLogLevel(logLevel string) {
	if logLevel == "" {
		log.Tracef("No log level configured, defaulting to: '%s'", log.GetLevel())
		return
	}

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.WithError(err).
			Debugf("Invalid log level '%s' configured, defaulting to: '%s'", logLevel, log.GetLevel())
	} else {
		log.SetLevel(level)
		log.Tracef("Log level from configuration: %s", logLevel)
	}
}
