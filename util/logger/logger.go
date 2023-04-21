package logger

func Errorln(args ...interface{}) {
	Logger.Errorln(args)
}

func Infoln(args ...interface{}) {
	Logger.Infoln(args)
}

func Debugln(args ...interface{}) {
	Logger.Debugln(args)
}
