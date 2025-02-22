// Code generated by goa v3.5.5, DO NOT EDIT.
//
// pipeline endpoints
//
// Command:
// $ goa gen github.com/penwern/enduro/internal/api/design -o internal/api

package pipeline

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "pipeline" service endpoints.
type Endpoints struct {
	List       goa.Endpoint
	Show       goa.Endpoint
	Processing goa.Endpoint
}

// NewEndpoints wraps the methods of the "pipeline" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List:       NewListEndpoint(s),
		Show:       NewShowEndpoint(s),
		Processing: NewProcessingEndpoint(s),
	}
}

// Use applies the given middleware to all the "pipeline" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Show = m(e.Show)
	e.Processing = m(e.Processing)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "pipeline".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListPayload)
		return s.List(ctx, p)
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "pipeline".
func NewShowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowPayload)
		res, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedEnduroStoredPipeline(res, "default")
		return vres, nil
	}
}

// NewProcessingEndpoint returns an endpoint function that calls the method
// "processing" of service "pipeline".
func NewProcessingEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ProcessingPayload)
		return s.Processing(ctx, p)
	}
}
