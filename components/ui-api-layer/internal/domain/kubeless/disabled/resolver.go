// Code generated by failery v1.0.0. DO NOT EDIT.

package disabled

import context "context"
import gqlschema "github.com/kyma-project/kyma/components/ui-api-layer/internal/gqlschema"

// resolver is an autogenerated failing mock type for the resolver type
type resolver struct {
	err error
}

// NewResolver creates a new resolver type instance
func NewResolver(err error) *resolver {
	return &resolver{err: err}
}

// FunctionsQuery provides a failing mock function with given fields: ctx, environment, first, offset
func (_m *resolver) FunctionsQuery(ctx context.Context, environment string, first *int, offset *int) ([]gqlschema.Function, error) {
	var r0 []gqlschema.Function
	var r1 error
	r1 = _m.err

	return r0, r1
}