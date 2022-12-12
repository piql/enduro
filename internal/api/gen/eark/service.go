// Code generated by goa v3.5.5, DO NOT EDIT.
//
// eark service
//
// Command:
// $ goa gen github.com/penwern/enduro/internal/api/design -o internal/api

package eark

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// The eark service manages eark workflows.
type Service interface {
	// Submit a new E-Ark AIP Workflow
	GenEarkAips(context.Context) (res *EarkResult, err error)
	// Retrieve status of current E-Ark AIP Workflow operation.
	AipGenStatus(context.Context) (res *EarkStatusResult, err error)
	// Submit a new E-Ark DIP Workflow
	GenEarkDips(context.Context) (res *EarkDIPResult, err error)
	// Retrieve status of current E-Ark DIP Workflow operation.
	DipGenStatus(context.Context) (res *EarkStatusResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "eark"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"gen_eark_aips", "aip_gen_status", "gen_eark_dips", "dip_gen_status"}

// EarkResult is the result type of the eark service gen_eark_aips method.
type EarkResult struct {
	WorkflowID string
	RunID      string
}

// EarkStatusResult is the result type of the eark service aip_gen_status
// method.
type EarkStatusResult struct {
	Running    bool
	Status     *string
	WorkflowID *string
	RunID      *string
}

// EarkDIPResult is the result type of the eark service gen_eark_dips method.
type EarkDIPResult struct {
	Success bool
	AipName *string
	DipName *string
}

// MakeNotAvailable builds a goa.ServiceError from an error.
func MakeNotAvailable(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not_available",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotValid builds a goa.ServiceError from an error.
func MakeNotValid(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not_valid",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
