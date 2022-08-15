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

// SubmitResponseBody is the type of the "eark" service "submit" endpoint HTTP
// response body.
type SubmitResponseBody struct {
	WorkflowID string `form:"workflow_id" json:"workflow_id" xml:"workflow_id"`
	RunID      string `form:"run_id" json:"run_id" xml:"run_id"`
}

// StatusResponseBody is the type of the "eark" service "status" endpoint HTTP
// response body.
type StatusResponseBody struct {
	Running    bool    `form:"running" json:"running" xml:"running"`
	Status     *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	WorkflowID *string `form:"workflow_id,omitempty" json:"workflow_id,omitempty" xml:"workflow_id,omitempty"`
	RunID      *string `form:"run_id,omitempty" json:"run_id,omitempty" xml:"run_id,omitempty"`
}

// SubmitNotAvailableResponseBody is the type of the "eark" service "submit"
// endpoint HTTP response body for the "not_available" error.
type SubmitNotAvailableResponseBody struct {
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

// SubmitNotValidResponseBody is the type of the "eark" service "submit"
// endpoint HTTP response body for the "not_valid" error.
type SubmitNotValidResponseBody struct {
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

// NewSubmitResponseBody builds the HTTP response body from the result of the
// "submit" endpoint of the "eark" service.
func NewSubmitResponseBody(res *eark.EarkResult) *SubmitResponseBody {
	body := &SubmitResponseBody{
		WorkflowID: res.WorkflowID,
		RunID:      res.RunID,
	}
	return body
}

// NewStatusResponseBody builds the HTTP response body from the result of the
// "status" endpoint of the "eark" service.
func NewStatusResponseBody(res *eark.EarkStatusResult) *StatusResponseBody {
	body := &StatusResponseBody{
		Running:    res.Running,
		Status:     res.Status,
		WorkflowID: res.WorkflowID,
		RunID:      res.RunID,
	}
	return body
}

// NewSubmitNotAvailableResponseBody builds the HTTP response body from the
// result of the "submit" endpoint of the "eark" service.
func NewSubmitNotAvailableResponseBody(res *goa.ServiceError) *SubmitNotAvailableResponseBody {
	body := &SubmitNotAvailableResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewSubmitNotValidResponseBody builds the HTTP response body from the result
// of the "submit" endpoint of the "eark" service.
func NewSubmitNotValidResponseBody(res *goa.ServiceError) *SubmitNotValidResponseBody {
	body := &SubmitNotValidResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}
