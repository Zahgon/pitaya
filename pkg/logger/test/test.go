package test

import (
	tests "github.com/sirupsen/logrus/hooks/test"
	"github.com/topfreegames/pitaya/v3/pkg/logger/interfaces"
)

// NewNullLogger creates a discarding logger and installs the test hook.
func NewNullLogger() (interfaces.Logger, *tests.Hook) {
	_ = "STUB: not implemented"
	return *new(interfaces.Logger), nil
}
