#Logger 

Logger wraps around the logrus logging package giving options for including information
about where a log message was called. It includes the function name, file, and line number.

```
{
    "file": "/Users/realugbun/go/src/github.com/realugbun/somepackage/main.go",
    "func": "main.someFunc",
    "level": "info",
    "line": 87,
    "msg": "example log message",
    "time": "2022-02-06T12:50:44-05:00"
}
```

It also includes the option to provide a stack trace.

```
{
    "file": "/Users/realugbun/go/src/github.com/realugbun/somepackage/main.go",
    "func": "main.someFunc",
    "level": "info",
    "line": 87,
    "msg": "example log message",
    "time": "2022-02-06T12:55:51-05:00",
    "trace": [
        {
            "file": "/Users/realugbun/go/src/github.com/realugbun/mongodb/main.go",
            "function": "main.main",
            "line": 31
        },
        {
            "file": "/usr/local/opt/go/libexec/src/runtime/proc.go",
            "function": "runtime.main",
            "line": 255
        },
        {
            "file": "/usr/local/opt/go/libexec/src/runtime/asm_amd64.s",
            "function": "runtime.goexit",
            "line": 1581
        }
    ]
}
```

## Initialization

Logger has two options for initing.
 
`Init()` starts the logger with default options. 

- Loggs to standard output
- Log level set to info
- Information about where the log was called is included

`InitWithOptions()` allows customizing the logging behavior.


```
options := logger.NewOptions() 		// Creates a new options struct
options.SetFile("filename.log")		// Sets the file where logs should be written
options.SetIncludeFunc(true)		// Include information about the calling function
options.SetLevel("info")			// Set the log level

st := logger.NewStackTrace()		// Creates a new struct for stack trace options and enables stack tracing
st.SetMaxEntries(5)					// Set the maximum number of entries to include in the trace
st.SetStopFile("main.go")			// Set a filename once reached the stack trace will stop
st.SetStopFunction("main.main")		// Set a function name once reached the stack trace will stop
st.SetLambda(true)					// Sets the stack trace to stop when it reaches the function AWS uses to invoke Lambda

options.SetStackTrace(*st)			// Sets the stack trace options on logger options

logger.InitWithOptions(options)		// Init the logger with options
```

## Usage

Logger supports two types of logging which match closely with logrus. `logger.Info()`, `logger.Trace()` etc.

The second option allows adding custom fields which are a slice of `map[string]interface{}`. These follow the naming convention `logger.InfoWithFields(fields)` etc.
 