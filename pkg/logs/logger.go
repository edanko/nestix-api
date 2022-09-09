package logs

type LogFields map[string]any

type Logger interface {
	Fatal(msg string, err error, fields LogFields)
	Error(msg string, err error, fields LogFields)
	Info(msg string, fields LogFields)
	Debug(msg string, fields LogFields)
	Trace(msg string, fields LogFields)
	With(fields LogFields) Logger
}
