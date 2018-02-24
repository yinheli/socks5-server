package util

import (
	"github.com/op/go-logging"
	"os"
)

const module = "kungfu-lite"

var (
	log          = logging.MustGetLogger(module)
	levelBackend logging.LeveledBackend
)

func init() {
	format := logging.MustStringFormatter(
		`%{color}%{time:2006-01-02 15:04:05.000} [%{level:.4s}]%{color:reset} - %{message}`,
	)

	backend := logging.NewLogBackend(os.Stdout, "", 0)
	levelBackend = logging.AddModuleLevel(logging.NewBackendFormatter(backend, format))
	levelBackend.SetLevel(logging.INFO, module)
	log.SetBackend(levelBackend)
}

// GetLog is log util
func GetLog() *logging.Logger {
	return log
}

// SetLogLevelDebug is for setting global log level
func SetLogLevelDebug() {
	levelBackend.SetLevel(logging.DEBUG, module)
}
