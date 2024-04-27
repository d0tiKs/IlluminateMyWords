package logger

import (
	"fmt"
	"os"
)

const (
	// PROGRAM PROPERTIES
	STDIN  = 0
	STDOUT = 1
	STDERR = 2

	STAND_FD = STDERR
	ERROR_FD = STDERR

	// LOG LEVELS
	LOG_DEBUG   = "DEBUG"
	LOG_INFO    = "INFO"
	LOG_ERROR   = "ERROR"
	LOG_WARNING = "WARNING"
)

type LogConfig struct {
	verbose *bool
	outfd   *os.File
}

var logconfig LogConfig

func InitLogger(verbose *bool, outfd int) {
	logconfig.verbose = verbose
	switch outfd {
	case STDERR:
		logconfig.outfd = os.Stderr
	case STDOUT:
	default:
		logconfig.outfd = os.Stdout
	}
}

func LogMessage(logLevel string, format string, vargs ...interface{}) (n int, err error) {
	if !*logconfig.verbose && logLevel == LOG_DEBUG {
		return 0, nil
	}

	logLevelToken := fmt.Sprintf("[%s] ", logLevel)
	message := fmt.Sprintf(format, vargs...)

	return fmt.Fprintln(logconfig.outfd, logLevelToken+message)
}
