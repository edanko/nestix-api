package logs

var _ Logger = (*NopLogger)(nil)

// NopLogger is a logger which discards all logs.
type NopLogger struct{}

func (NopLogger) Fatal(msg string, err error, fields LogFields) {}
func (NopLogger) Error(msg string, err error, fields LogFields) {}
func (NopLogger) Info(msg string, fields LogFields)             {}
func (NopLogger) Debug(msg string, fields LogFields)            {}
func (NopLogger) Trace(msg string, fields LogFields)            {}
func (l NopLogger) With(fields LogFields) Logger                { return l }
