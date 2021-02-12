// Package dicomlog performs logging for go-dicom or go-netdicom.  It is
// currently just a thin wrapper around the standard log package.
package dicomlog

import (
	"github.com/sirupsen/logrus"
	"log"
	"sync/atomic"
)

// level sets log verbosity. The larger the value, the more verbose.  Setting it
// to -1 disables logging completely.
var (
	level  = int32(0)
	logger *logrus.Logger
)

// SetLevel sets log verbosity. The larger the value, the more verbose. Setting
// it to -1 disables logging completely. Thread safe.
func SetLevel(l int) {
	atomic.StoreInt32(&level, int32(l))
}

func SetLogger(log *logrus.Logger) {
	logger = log
}

// Level returns the current log level. The larger the value, the more verbose.
// Thread safe.
func Level() int {
	return int(atomic.LoadInt32(&level))
}

// Vprintf is shorthand for "if level > Level { log.Printf(...) }".
func Vprintf(l int, format string, args ...interface{}) {
	if Level() >= l {
		if logger != nil {
			logger.Printf(format, args...)
		} else {
			log.Printf(format, args...)
		}

	}
}
