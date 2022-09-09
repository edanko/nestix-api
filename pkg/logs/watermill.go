package logs

import (
	"github.com/ThreeDotsLabs/watermill"
)

var _ watermill.LoggerAdapter = (*WatermillLogger)(nil)

type WatermillLogger struct {
	logger Logger
}

// NewWatermillLogger gets a new zerolog adapter for use in the watermill context.
func NewWatermillLogger(l Logger) *WatermillLogger {
	return &WatermillLogger{
		logger: l,
	}
}

// Logs an error message.
func (l *WatermillLogger) Error(
	msg string,
	err error,
	fields watermill.LogFields,
) {
	l.logger.Error(msg, err, LogFields(fields))
}

// Info logs an info message.
func (l *WatermillLogger) Info(
	msg string,
	fields watermill.LogFields,
) {
	l.logger.Info(msg, LogFields(fields))
}

// Debug logs a debug message.
func (l *WatermillLogger) Debug(
	msg string,
	fields watermill.LogFields,
) {
	l.logger.Debug(msg, LogFields(fields))
}

// Trace logs a trace message.
func (l *WatermillLogger) Trace(
	msg string,
	fields watermill.LogFields,
) {
	l.logger.Trace(msg, LogFields(fields))
}

// With creates new adapter with the input fields as context.
func (l *WatermillLogger) With(
	fields watermill.LogFields,
) watermill.LoggerAdapter {
	return &WatermillLogger{
		logger: l.logger.With(LogFields(fields)),
	}
}
