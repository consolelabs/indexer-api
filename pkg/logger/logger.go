package logger

type Fields map[string]interface{}

type Logger interface {
	Fields(data Fields) Logger
	Field(key, value string) Logger
	AddField(key string, value any) Logger
	Debug(msg string)
	Debugf(msg string, args ...interface{})
	Info(msg string)
	Infof(msg string, args ...interface{})
	Warn(msg string)
	Warnf(msg string, args ...interface{})
	Error(err error, msg string)
	Errorf(err error, msg string, args ...interface{})
	Fatal(err error, msg string)
	Fatalf(err error, msg string, args ...interface{})
}
