package logger

import (
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

// stackTrace gets the file, line, and function name where the log message was called.
// If stackTraceOptions are defined, it also attaches a stack trace.
// To get an accurate stackTrace the log should be called within the function
// instead of after returning from the function. This is important for errors.
// logger.Error should be called within the function where the error happened
// not after that function returns.
func stackTrace() (fields logrus.Fields) {

	// Don't include func name if disabled
	if !options.GetIncludeFunc() {
		return
	}

	var (
		pc       = make([]uintptr, 15)
		n        = runtime.Callers(2, pc)
		frames   = runtime.CallersFrames(pc[:n])
		isCaller bool
		file     string
		line     int
		function string
		trace    []map[string]interface{}
		counter  int
	)

	isMore := true
	for isMore && counter < 500 {

		// This would prevent an infinate loop in some strage situation
		counter++

		frame, more := frames.Next()
		isMore = more

		// Skip frames to the logger package which should always be the first frame
		if isLoggerCall(frame.File) {
			// The next frame will be the line which called the logger
			isCaller = true
			continue
		}
		// Adds the file, line number, and function name to the main entry
		if isCaller {
			file = frame.File
			line = frame.Line
			function = frame.Function
			isCaller = false
			continue
		}

		// Only add the stack trace if it is enabled
		if options.StackTrace == nil {
			break
		}

		trace = append(trace, map[string]interface{}{
			"file":     frame.File,
			"line":     frame.Line,
			"function": frame.Function,
		})

		if len(trace) == options.StackTrace.GetMaxEntries() {
			break
		}

		// Stop once we reach a particular function name such as main.main
		if frame.Function == options.StackTrace.GetStopFunction() {
			break
		}

		// Stop once we reach a particular file name such as logger.go
		if options.StackTrace.GetStopFile() != "" {
			if strings.HasSuffix(frame.File, options.StackTrace.GetStopFile()) {
				break
			}
		}

		// Stop when we get to the lambda caller if the option is enabled
		if options.StackTrace.GetLambda() {
			if strings.HasPrefix(frame.Function, options.StackTrace.GetStopFunction()) {
				break
			}
		}

	}

	fields = logrus.Fields{
		"file": file,
		"line": line,
		"func": function,
	}

	if len(trace) > 0 {
		fields["trace"] = trace
	}

	return
}

// isLoggerCall checks if the stack trace is a call from the logger
func isLoggerCall(file string) bool {

	fs := strings.Split(file, "/")
	l := fs[len(fs)-2]
	return l == "logger"

}

func test() {

	options := logger.NewOptions()
	options.SetFile("filename.log")
	options.SetIncludeFunc(true)
	options.SetLevel("info")

	st := logger.NewStackTrace()
	st.SetMaxEntries(5)
	st.SetStopFile("main.go")
	st.SetStopFunction("main.main")
	st.SetLambda(true)

	options.SetStackTrace(*st)

	logger.InitWithOptions(options)

}
