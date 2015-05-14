// logger provides a simple, buffered logger.  It composes log.Logger
// for logging, and bytes.Buffer for writing output.  For example:
//
//	bufferedLogger := logger.New()
//	bufferedLogger.Printf("Text")
//	// ...other operations, potentially writing to STDERR
//	bufferedLogger.WriteTo(os.Stderr)
//
// This will write all entries written to bufferedLogger out to STDERR
// together, regardless of what else may have been written in between.
package logger

import (
	"bytes"
	"log"
)

// Logger behaves as a bytes.Buffer that you can write to using the
// log.Logger interface.
type Logger struct {
	*log.Logger
	*bytes.Buffer
}

// New creates a new Logger instance with a fresh buffer.
func New() *Logger {
	buffer := bytes.NewBuffer([]byte{})
	logger := log.New(buffer, "", log.LstdFlags)
	return &Logger{logger, buffer}
}
