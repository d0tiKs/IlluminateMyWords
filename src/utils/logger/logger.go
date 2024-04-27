package loggerutils

import (
	configutils "IlluminateMyWords/src/utils/config"
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

func LogMessage(logLevel string, format string, vargs ...interface{}) (n int, err error) {
	if !*configutils.Config.Verbose && logLevel == LOG_DEBUG {
		return 0, nil
	}

	logLevelToken := fmt.Sprintf("[%s] ", logLevel)
	message := fmt.Sprintf(format, vargs...)

	outfd := os.Stderr

	switch STAND_FD {
	case STDERR:
		outfd = os.Stderr
	case STDOUT:
	default:
		outfd = os.Stdout
	}

	return fmt.Fprintln(outfd, logLevelToken+message)
}
