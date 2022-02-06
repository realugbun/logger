package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	defaultLevel = logrus.InfoLevel
)

type (
	// LogFunction used for passing functions in logrus function logging
	LogFunction func() []interface{}
	// Fields used with WithFields calls ie InfoWithFields
	Fields map[string]interface{}
)

var (
	options *Options
)

// Init sets up the logger with default options: Log level info, IncludeFunc true, and output to standard output.
// For more options use, InitWithOptions
func Init() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Use the default log level
	SetLevel(defaultLevel.String())

	// Defaults to including the function info in the log
	o := NewOptions().SetIncludeFunc(true)
	options = o

}

// InitWithOptions inits the logger using the passed options
func InitWithOptions(o *Options) {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	options = o

	// If a file location is passed, logging will be made to the file. Otherwise it goes to standard output.
	if o.File != nil {
		f, err := os.OpenFile(o.GetFile(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			logrus.Warn("unable to open file", err)
		}
		logrus.SetOutput(f)
	}

	// Set the log level based on options
	if o.Level == nil {
		o.SetLevel(defaultLevel.String())
		logrus.Info("log level not set using default")
	}

	SetLevel(o.GetLevel())

	if options.StackTrace != nil {
		if options.StackTrace.GetLambda() {
			options.StackTrace.SetStopFunction("github.com/aws/aws-lambda-go/lambda.NewHandler")
		}
	}

}

// SetLevel sets the logging level
func SetLevel(level string) {
	l, err := logrus.ParseLevel(strings.ToLower(level))
	if err != nil {
		l = defaultLevel
		logrus.Warn("invalid log level using defult")
	}
	logrus.SetLevel(l)

	logrus.Info("logging started at level " + l.String())
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Print(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Warning(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Fatal(args...)
}

// TraceFn logs a message from a func at level Trace on the standard logger.
func TraceFn(fn LogFunction) {
	logrus.TraceFn(logrus.LogFunction(fn))
}

// DebugFn logs a message from a func at level Debug on the standard logger.
func DebugFn(fn LogFunction) {
	logrus.DebugFn(logrus.LogFunction(fn))
}

// PrintFn logs a message from a func at level Info on the standard logger.
func PrintFn(fn LogFunction) {
	logrus.PrintFn(logrus.LogFunction(fn))
}

// InfoFn logs a message from a func at level Info on the standard logger.
func InfoFn(fn LogFunction) {
	logrus.InfoFn(logrus.LogFunction(fn))
}

// WarnFn logs a message from a func at level Warn on the standard logger.
func WarnFn(fn LogFunction) {
	logrus.WarnFn(logrus.LogFunction(fn))
}

// WarningFn logs a message from a func at level Warn on the standard logger.
func WarningFn(fn LogFunction) {
	logrus.WarningFn(logrus.LogFunction(fn))
}

// ErrorFn logs a message from a func at level Error on the standard logger.
func ErrorFn(fn LogFunction) {
	logrus.ErrorFn(logrus.LogFunction(fn))
}

// PanicFn logs a message from a func at level Panic on the standard logger.
func PanicFn(fn LogFunction) {
	logrus.PanicFn(logrus.LogFunction(fn))
}

// FatalFn logs a message from a func at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalFn(fn LogFunction) {
	logrus.FatalFn(logrus.LogFunction(fn))
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Fatalf(format, args...)
}

// Traceln logs a message at level Trace on the standard logger.
func Traceln(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Traceln(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Debugln(args...)
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Println(args...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Warnln(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Warningln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Errorln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Panicln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(args ...interface{}) {
	f := stackTrace()
	logrus.WithFields(f).Fatalln(args...)
}

func (f *Fields) addFields(fields Fields) Fields {

	ff := *f

	for k, v := range fields {
		ff[k] = v
	}
	f = &ff
	return *f
}

// TraceWithFields logs a message with custom fields at level Trace on the standard logger.
func TraceWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Trace(args...)
}

// DebugWithFields logs a message with custom fields at level Debug on the standard logger.
func DebugWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Debug(args...)
}

// PrintWithFields logs a message with custom fields at level Info on the standard logger.
func PrintWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Print(args...)
}

// InfoWithFields logs a message with custom fields at level Info on the standard logger.
func InfoWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Info(args...)
}

// WarnWithFields logs a message with custom fields at level Warn on the standard logger.
func WarnWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Warn(args...)
}

// WarningWithFields logs a message with custom fields at level Warn on the standard logger.
func WarningWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Warning(args...)
}

// ErrorWithFields logs a message with custom fields at level Error on the standard logger.
func ErrorWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Error(args...)
}

// PanicWithFields logs a message with custom fields at level Panic on the standard logger.
func PanicWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Panic(args...)
}

// FatalWithFields logs a message with custom fields at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Fatal(args...)
}

// TracefWithFields logs a message with custom fields at level Trace on the standard logger.
func TracefWithFields(fields Fields, format string, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Tracef(format, args...)
}

// DebugfWithFields logs a message with custom fields at level Debug on the standard logger.
func DebugfWithFields(fields Fields, format string, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Debugf(format, args...)
}

// PrintfWithFields logs a message with custom fields at level Info on the standard logger.
func PrintfWithFields(fields Fields, format string, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Printf(format, args...)
}

// InfofWithFields logs a message with custom fields at level Info on the standard logger.
func InfofWithFields(fields Fields, format string, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Infof(format, args...)
}

// WarnfWithFields logs a message with custom fields at level Warn on the standard logger.
func WarnfWithFields(fields Fields, format string, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Warnf(format, args...)
}

// WarningfWithFields logs a message with custom fields at level Warn on the standard logger.
func WarningfWithFields(fields Fields, format string, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Warningf(format, args...)
}

// ErrorfWithFields logs a message with custom fields at level Error on the standard logger.
func ErrorfWithFields(fields Fields, format string, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Errorf(format, args...)
}

// PanicfWithFields logs a message with custom fields at level Panic on the standard logger.
func PanicfWithFields(fields Fields, format string, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Panicf(format, args...)
}

// FatalfWithFields logs a message with custom fields at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalfWithFields(fields Fields, format string, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Fatalf(format, args...)
}

// TracelnWithFields logs a message with custom fields at level Trace on the standard logger.
func TracelnWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Traceln(args...)
}

// DebuglnWithFields logs a message with custom fields at level Debug on the standard logger.
func DebuglnWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Debugln(args...)
}

// PrintlnWithFields logs a message with custom fields at level Info on the standard logger.
func PrintlnWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Println(args...)
}

// InfolnWithFields logs a message with custom fields at level Info on the standard logger.
func InfolnWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Infoln(args...)
}

// WarnlnWithFields logs a message with custom fields at level Warn on the standard logger.
func WarnlnWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Warnln(args...)
}

// WarninglnWithFields logs a message with custom fields at level Warn on the standard logger.
func WarninglnWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Warningln(args...)
}

// ErrorlnWithFields logs a message with custom fields at level Error on the standard logger.
func ErrorlnWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Errorln(args...)
}

// PaniclnWithFields logs a message with custom fields at level Panic on the standard logger.
func PaniclnWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Panicln(args...)
}

// FatallnWithFields logs a message with custom fields at level Fatal on the standard logger then the process will exit with status set to 1.
func FatallnWithFields(fields Fields, args ...interface{}) {
	fields.addFields(Fields(stackTrace()))
	logrus.WithFields(logrus.Fields(fields)).Fatalln(args...)
}
