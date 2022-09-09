package logs

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var _ Logger = (*ZerologLogger)(nil)

type ZerologLogger struct {
	logger zerolog.Logger
}

// NewZerologLogger gets a new zerolog adapter for use in the application context.
func NewZerologLogger(levelStr string) Logger {
	level, err := zerolog.ParseLevel(levelStr)
	if err != nil {
		panic(err)
	}

	loggerOutput := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339Nano,
	}
	zeroLogger := zerolog.New(loggerOutput).
		Level(level).
		With().
		Timestamp().
		// Str("app", "launch-api").
		Logger()

	return &ZerologLogger{
		logger: zeroLogger,
	}
}

func addFieldsData(
	event *zerolog.Event,
	fields LogFields,
) {
	for i, v := range fields {
		event.Interface(i, v)
	}
}

// Fatal logs an error message and terminates the application immediately.
func (l *ZerologLogger) Fatal(
	msg string,
	err error,
	fields LogFields,
) {
	event := l.logger.Fatal().Err(err)

	if fields != nil {
		addFieldsData(event, fields)
	}

	event.Msg(msg)
}

// Error logs an error message.
func (l *ZerologLogger) Error(
	msg string,
	err error,
	fields LogFields,
) {
	event := l.logger.Err(err)

	if fields != nil {
		addFieldsData(event, fields)
	}

	event.Msg(msg)
}

// Info logs an info message.
func (l *ZerologLogger) Info(
	msg string,
	fields LogFields,
) {
	event := l.logger.Info()

	if fields != nil {
		addFieldsData(event, fields)
	}

	event.Msg(msg)
}

// Debug logs a debug message.
func (l *ZerologLogger) Debug(
	msg string,
	fields LogFields,
) {
	event := l.logger.Debug()

	if fields != nil {
		addFieldsData(event, fields)
	}

	event.Msg(msg)
}

// Trace logs a trace message.
func (l *ZerologLogger) Trace(
	msg string,
	fields LogFields,
) {
	event := l.logger.Trace()

	if fields != nil {
		addFieldsData(event, fields)
	}

	event.Msg(msg)
}

// With creates new adapter with the input fields as context.
func (l *ZerologLogger) With(
	fields LogFields,
) Logger {
	if fields == nil {
		return l
	}

	subLog := l.logger.With()

	for i, v := range fields {
		subLog = subLog.Interface(i, v)
	}

	return &ZerologLogger{
		logger: subLog.Logger(),
	}
}
