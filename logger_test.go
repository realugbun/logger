package logger

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InitWithOptions(t *testing.T) {

	for _, tc := range []struct {
		name     string
		logLevel string
		expLog   string
	}{
		{
			name:   "happy path",
			expLog: "logging started at level info",
		},
		{
			name:     "invalid level",
			logLevel: "invalid",
			expLog:   "invalid log level using defult",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {

			options := NewOptions().SetFile("./Test_InitWithOptions.log")
			defer os.Remove(options.GetFile())

			if tc.logLevel != "" {
				options.SetLevel(tc.logLevel)
			}

			InitWithOptions(options)

			data, err := os.ReadFile(options.GetFile())
			assert.NoError(t, err)

			assert.True(t, strings.Contains(string(data), tc.expLog))

		})
	}
}

func Test_stackTrace(t *testing.T) {

	for _, tc := range []struct {
		name     string
		options  *StackTrace
		expFunc  bool
		expTrace bool
	}{
		{
			name:     "No trace",
			expFunc:  false,
			expTrace: false,
		},
		{
			name:     "With trace",
			options:  &StackTrace{},
			expFunc:  true,
			expTrace: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {

			options := NewOptions()

			if tc.options != nil {
				options = options.SetStackTrace(*tc.options)
			}

			if tc.expFunc {
				options.SetIncludeFunc(true)
			}

			InitWithOptions(options)
			actOut := stackTrace()

			_, isTrace := actOut["trace"]
			function := actOut["func"]
			line := actOut["line"]
			file := actOut["file"]
			assert.Equal(t, tc.expTrace, isTrace)

			if tc.expFunc {
				assert.NotEmpty(t, function)
				assert.NotEmpty(t, line)
				assert.NotEmpty(t, file)
			}

		})
	}
}

func Test_test(t *testing.T) {
	Init()
	Info("example log")
}
