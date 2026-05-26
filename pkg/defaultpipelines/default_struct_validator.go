package defaultpipelines

import (
	"context"
	"sync"

	validator "github.com/go-playground/validator/v10"
)

// DefaultValidator is the default arguments validator for handlers
// in pitaya
type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

// Validate is the the function responsible for validating the 'in' parameter
// based on the struct tags the parameter has.
// This function has the pipeline.Handler signature so
// it is possible to use it as a pipeline function
func (v *DefaultValidator) Validate(ctx context.Context, in interface{}) (context.Context, interface{}, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func (v *DefaultValidator) lazyinit() { _ = "STUB: not implemented"; return }
