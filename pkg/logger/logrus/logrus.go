package logrus

import (
	"github.com/sirupsen/logrus"
	"github.com/topfreegames/pitaya/v3/pkg/logger/interfaces"
)

type logrusImpl struct {
	logrus.FieldLogger
}

// New returns a new interfaces.Logger implementation based on logrus
func New() interfaces.Logger { _ = "STUB: not implemented"; return *new(interfaces.Logger) }

// NewWithEntry returns a new interfaces.Logger implementation based on a provided logrus entry instance
// Deprecated: NewWithEntry is deprecated.
func NewWithEntry(logger *logrus.Entry) interfaces.Logger {
	_ = "STUB: not implemented"
	return *new(interfaces.Logger)
}

// NewWithLogger returns a new interfaces.Logger implementation based on a provided logrus instance
// Deprecated: NewWithLogger is deprecated.
func NewWithLogger(logger *logrus.Logger) interfaces.Logger {
	_ = "STUB: not implemented"
	return *new(interfaces.Logger)
}

// NewWithFieldLogger returns a new interfaces.Logger implementation based on a provided logrus instance
func NewWithFieldLogger(logger logrus.FieldLogger) interfaces.Logger {
	_ = "STUB: not implemented"
	return *new(interfaces.Logger)
}

func (l *logrusImpl) WithFields(fields map[string]interface{}) interfaces.Logger {
	_ = "STUB: not implemented"
	return *new(interfaces.Logger)
}

func (l *logrusImpl) WithField(key string, value interface{}) interfaces.Logger {
	_ = "STUB: not implemented"
	return *new(interfaces.Logger)
}

func (l *logrusImpl) WithError(err error) interfaces.Logger {
	_ = "STUB: not implemented"
	return *new(interfaces.Logger)
}

func (l *logrusImpl) GetInternalLogger() any { _ = "STUB: not implemented"; return *new(any) }
