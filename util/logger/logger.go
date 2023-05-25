package logger

type Logger interface {
	Errorln(args ...interface{})
	Warnln(args ...interface{})
	Infoln(args ...interface{})
	Debugln(args ...interface{})
	Sync() error
}

var L Logger
