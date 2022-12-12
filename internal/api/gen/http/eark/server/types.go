// Code generated by goa v3.5.5, DO NOT EDIT.
//
// eark HTTP server types
//
// Command:
// $ goa gen github.com/penwern/enduro/internal/api/design -o internal/api

package server

import (
	eark "github.com/penwern/enduro/internal/api/gen/eark"
	goa "goa.design/goa/v3/pkg"
)

// GenEarkAipsResponseBody is the type of the "eark" service "gen_eark_aips"
// endpoint HTTP response body.
type GenEarkAipsResponseBody struct {
	WorkflowID string `form:"workflow_id" json:"workflow_id" xml:"workflow_id"`
	RunID      string `form:"run_id" json:"run_id" xml:"run_id"`
}

// AipGenStatusResponseBody is the type of the "eark" service "aip_gen_status"
// endpoint HTTP response body.
type AipGenStatusResponseBody struct {
	Running    bool    `form:"running" json:"running" xml:"running"`
	Status     *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	WorkflowID *string `form:"workflow_id,omitempty" json:"workflow_id,omitempty" xml:"workflow_id,omitempty"`
	RunID      *string `form:"run_id,omitempty" json:"run_id,omitempty" xml:"run_id,omitempty"`
}

// CreateDipsResponseBody is the type of the "eark" service "create_dips"
// endpoint HTTP response body.
type CreateDipsResponseBody struct {
	Success bool    `form:"success" json:"success" xml:"success"`
	AipName *string `form:"aip_name,omitempty" json:"aip_name,omitempty" xml:"aip_name,omitempty"`
	DipName *string `form:"dip_name,omitempty" json:"dip_name,omitempty" xml:"dip_name,omitempty"`
}

// DipGenStatusResponseBody is the type of the "eark" service "dip_gen_status"
// endpoint HTTP response body.
type DipGenStatusResponseBody struct {
	Running    bool    `form:"running" json:"running" xml:"running"`
	Status     *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	WorkflowID *string `form:"workflow_id,omitempty" json:"workflow_id,omitempty" xml:"workflow_id,omitempty"`
	RunID      *string `form:"run_id,omitempty" json:"run_id,omitempty" xml:"run_id,omitempty"`
}

// GenEarkAipsNotAvailableResponseBody is the type of the "eark" service
// "gen_eark_aips" endpoint HTTP response body for the "not_available" error.
type GenEarkAipsNotAvailableResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GenEarkAipsNotValidResponseBody is the type of the "eark" service
// "gen_eark_aips" endpoint HTTP response body for the "not_valid" error.
type GenEarkAipsNotValidResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateDipsNotAvailableResponseBody is the type of the "eark" service
// "create_dips" endpoint HTTP response body for the "not_available" error.
type CreateDipsNotAvailableResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateDipsNotValidResponseBody is the type of the "eark" service
// "create_dips" endpoint HTTP response body for the "not_valid" error.
type CreateDipsNotValidResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewGenEarkAipsResponseBody builds the HTTP response body from the result of
// the "gen_eark_aips" endpoint of the "eark" service.
func NewGenEarkAipsResponseBody(res *eark.EarkResult) *GenEarkAipsResponseBody {
	body := &GenEarkAipsResponseBody{
		WorkflowID: res.WorkflowID,
		RunID:      res.RunID,
	}
	return body
}

// NewAipGenStatusResponseBody builds the HTTP response body from the result of
// the "aip_gen_status" endpoint of the "eark" service.
func NewAipGenStatusResponseBody(res *eark.EarkStatusResult) *AipGenStatusResponseBody {
	body := &AipGenStatusResponseBody{
		Running:    res.Running,
		Status:     res.Status,
		WorkflowID: res.WorkflowID,
		RunID:      res.RunID,
	}
	return body
}

// NewCreateDipsResponseBody builds the HTTP response body from the result of
// the "create_dips" endpoint of the "eark" service.
func NewCreateDipsResponseBody(res *eark.EarkDIPResult) *CreateDipsResponseBody {
	body := &CreateDipsResponseBody{
		Success: res.Success,
		AipName: res.AipName,
		DipName: res.DipName,
	}
	return body
}

// NewDipGenStatusResponseBody builds the HTTP response body from the result of
// the "dip_gen_status" endpoint of the "eark" service.
func NewDipGenStatusResponseBody(res *eark.EarkStatusResult) *DipGenStatusResponseBody {
	body := &DipGenStatusResponseBody{
		Running:    res.Running,
		Status:     res.Status,
		WorkflowID: res.WorkflowID,
		RunID:      res.RunID,
	}
	return body
}

// NewGenEarkAipsNotAvailableResponseBody builds the HTTP response body from
// the result of the "gen_eark_aips" endpoint of the "eark" service.
func NewGenEarkAipsNotAvailableResponseBody(res *goa.ServiceError) *GenEarkAipsNotAvailableResponseBody {
	body := &GenEarkAipsNotAvailableResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGenEarkAipsNotValidResponseBody builds the HTTP response body from the
// result of the "gen_eark_aips" endpoint of the "eark" service.
func NewGenEarkAipsNotValidResponseBody(res *goa.ServiceError) *GenEarkAipsNotValidResponseBody {
	body := &GenEarkAipsNotValidResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateDipsNotAvailableResponseBody builds the HTTP response body from the
// result of the "create_dips" endpoint of the "eark" service.
func NewCreateDipsNotAvailableResponseBody(res *goa.ServiceError) *CreateDipsNotAvailableResponseBody {
	body := &CreateDipsNotAvailableResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateDipsNotValidResponseBody builds the HTTP response body from the
// result of the "create_dips" endpoint of the "eark" service.
func NewCreateDipsNotValidResponseBody(res *goa.ServiceError) *CreateDipsNotValidResponseBody {
	body := &CreateDipsNotValidResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}
