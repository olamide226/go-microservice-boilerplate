package logger

import (
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
)

// LogLevel represents different log levels.
type LogLevel string

const (
	// InfoLevel represents the info log level.
	InfoLevel LogLevel = "info"
	// DebugLevel represents the debug log level.
	DebugLevel LogLevel = "debug"
)

type Options struct {
	Environment string
}

// Logger is an interface representing the logging functionality.
type Logger interface {
	Info(string)
	Warn(string)
	Error(string, error)
	Debug(string)
	Fatal(string, error)
}

// loggerImpl is the implementation of the Logger interface using zerolog.
type loggerImpl struct {
	logger zerolog.Logger
}

var Global *loggerImpl

// NewLogger initializes the logger
func NewLogger(options Options) {

	hostname, err := os.Hostname()
	if err != nil {
		panic("could not get hostname")
	}
	if Global == nil {
		Global = &loggerImpl{}
		if options.Environment == "development" {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
			buildInfo, _ := debug.ReadBuildInfo()
			Global.logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
				Level(zerolog.TraceLevel).
				With().
				Timestamp().
				Caller().
				Int("pid", os.Getpid()).
				Str("go_version", buildInfo.GoVersion).
				Logger()
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
			Global.logger = zerolog.New(os.Stdout).
				With().Timestamp().Str("hostname", hostname).Logger()
		}
	}

}

// Info logs an informational message.
func (l *loggerImpl) Info(msg string) {
	l.logger.Info().Msg(msg)
}

// Error logs an error message.
func (l *loggerImpl) Error(msg string, err error) {

	l.logger.Error().Caller().Err(err).Msg(msg)
}

// Warn logs a warning message.
func (l *loggerImpl) Warn(msg string) {
	l.logger.Warn().Caller().Msg(msg)
}

// Debug logs a debug message.
func (l *loggerImpl) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l *loggerImpl) Fatal(msg string, err error) {
	l.logger.Fatal().Caller().Err(err).Msg(msg)
}

//Get the Global logger
func Get() Logger {
	return Global
}
