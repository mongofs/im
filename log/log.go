package log

type Logger interface {
	Error(err error)
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Fatal(interface{})
}


