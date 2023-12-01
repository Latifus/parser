package internal

import (
	"log"
	"log/slog"
	"os"
)

var logger *slog.Logger

func init() {
	logFile, err := os.OpenFile("parser.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	logger = slog.New(slog.NewJSONHandler(logFile, nil))
}

func Logger() *slog.Logger {
	return logger
}