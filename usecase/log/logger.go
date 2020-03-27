package log

type StdLogger interface {
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warningf(format string, args ...interface{})

	Debug(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Info(args ...interface{})
	Panic(args ...interface{})
	Print(args ...interface{})
	Warning(args ...interface{})

	Debugln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Infoln(args ...interface{})
	Panicln(args ...interface{})
	Println(args ...interface{})
	Warningln(args ...interface{})
}
