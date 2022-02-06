package logger

// Options for initiating the logger
type Options struct {

	// File location where to store logs. If left nil logging will go to standard output.
	File *string

	// IncludeFunc includes the function name, file name, and line number where the log message was called.
	IncludeFunc *bool

	// Level the level of logging
	Level *string

	// StackTrace options for including stack traces
	StackTrace *StackTrace
}

func NewOptions() *Options {
	return new(Options)
}

func (o *Options) SetFile(file string) *Options {
	o.File = &file
	return o
}

func (o *Options) GetFile() string {
	if o.File == nil {
		return ""
	}
	return *o.File
}

func (o *Options) SetIncludeFunc(b bool) *Options {
	o.IncludeFunc = &b
	return o
}

func (o *Options) GetIncludeFunc() bool {
	if o.IncludeFunc == nil {
		return false
	}
	return *o.IncludeFunc
}

func (o *Options) SetLevel(level string) *Options {
	o.Level = &level
	return o
}

func (o *Options) GetLevel() string {
	return *o.Level
}

func (o *Options) SetStackTrace(options StackTrace) *Options {
	o.StackTrace = &options
	return o
}

// StackTrace sets options for stack traceing
type StackTrace struct {

	// MaxEntries sets the maximum number of trace entries to include
	MaxEntries *int

	// StopFile tells the stack trace to ignore traces before a given file
	StopFile *string

	// StopFunction tells the stack trace to ignore traces before a given function ie main.main
	StopFunction *string

	// Lambda sets stop function or stop file variables for AWS lambda
	Lambda *bool
}

func NewStackTrace() *StackTrace {
	return new(StackTrace)
}

func (s *StackTrace) SetStopFunction(f string) *StackTrace {
	s.StopFunction = &f
	return s
}

func (s *StackTrace) GetStopFunction() string {
	if s.StopFunction == nil {
		return ""
	}
	return *s.StopFunction
}

func (s *StackTrace) SetStopFile(f string) *StackTrace {
	s.StopFile = &f
	return s
}

func (s *StackTrace) GetStopFile() string {
	if s.StopFile == nil {
		return ""
	}
	return *s.StopFile
}

func (s *StackTrace) SetMaxEntries(i int) *StackTrace {
	s.MaxEntries = &i
	return s
}

func (s *StackTrace) GetMaxEntries() int {
	if s.MaxEntries == nil {
		return 0
	}
	return *s.MaxEntries
}

func (s *StackTrace) SetLambda(b bool) *StackTrace {
	s.Lambda = &b
	return s
}

func (s *StackTrace) GetLambda() bool {
	if s.Lambda == nil {
		return false
	}
	return *s.Lambda
}
