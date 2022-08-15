// Code generated by goa v3.5.5, DO NOT EDIT.
//
// collection HTTP server types
//
// Command:
// $ goa-v3.5.5 gengithub.com/penwern/enduro/internal/api/design -o
// internal/api

package server

import (
	collection "github.com/artefactual-labs/enduro/internal/api/gen/collection"
	collectionviews "github.com/artefactual-labs/enduro/internal/api/gen/collection/views"
	goa "goa.design/goa/v3/pkg"
)

// BulkRequestBody is the type of the "collection" service "bulk" endpoint HTTP
// request body.
type BulkRequestBody struct {
	Operation *string `form:"operation,omitempty" json:"operation,omitempty" xml:"operation,omitempty"`
	Status    *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	Size      *uint   `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
}

// MonitorResponseBody is the type of the "collection" service "monitor"
// endpoint HTTP response body.
type MonitorResponseBody struct {
	// Identifier of collection
	ID uint `form:"id" json:"id" xml:"id"`
	// Type of the event
	Type string `form:"type" json:"type" xml:"type"`
	// Collection
	Item *EnduroStoredCollectionResponseBody `form:"item,omitempty" json:"item,omitempty" xml:"item,omitempty"`
}

// ListResponseBody is the type of the "collection" service "list" endpoint
// HTTP response body.
type ListResponseBody struct {
	Items      EnduroStoredCollectionCollectionResponseBody `form:"items" json:"items" xml:"items"`
	NextCursor *string                                      `form:"next_cursor,omitempty" json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
}

// ShowResponseBody is the type of the "collection" service "show" endpoint
// HTTP response body.
type ShowResponseBody struct {
	// Identifier of collection
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of the collection
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Status of the collection
	Status string `form:"status" json:"status" xml:"status"`
	// Identifier of processing workflow
	WorkflowID *string `form:"workflow_id,omitempty" json:"workflow_id,omitempty" xml:"workflow_id,omitempty"`
	// Identifier of latest processing workflow run
	RunID *string `form:"run_id,omitempty" json:"run_id,omitempty" xml:"run_id,omitempty"`
	// Identifier of Archivematica transfer
	TransferID *string `form:"transfer_id,omitempty" json:"transfer_id,omitempty" xml:"transfer_id,omitempty"`
	// Identifier of Archivematica AIP
	AipID *string `form:"aip_id,omitempty" json:"aip_id,omitempty" xml:"aip_id,omitempty"`
	// Identifier provided by the client
	OriginalID *string `form:"original_id,omitempty" json:"original_id,omitempty" xml:"original_id,omitempty"`
	// Identifier of Archivematica pipeline
	PipelineID *string `form:"pipeline_id,omitempty" json:"pipeline_id,omitempty" xml:"pipeline_id,omitempty"`
	// Creation datetime
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// Start datetime
	StartedAt *string `form:"started_at,omitempty" json:"started_at,omitempty" xml:"started_at,omitempty"`
	// Completion datetime
	CompletedAt *string `form:"completed_at,omitempty" json:"completed_at,omitempty" xml:"completed_at,omitempty"`
}

// WorkflowResponseBody is the type of the "collection" service "workflow"
// endpoint HTTP response body.
type WorkflowResponseBody struct {
	Status  *string                                               `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	History EnduroCollectionWorkflowHistoryResponseBodyCollection `form:"history,omitempty" json:"history,omitempty" xml:"history,omitempty"`
}

// BulkResponseBody is the type of the "collection" service "bulk" endpoint
// HTTP response body.
type BulkResponseBody struct {
	WorkflowID string `form:"workflow_id" json:"workflow_id" xml:"workflow_id"`
	RunID      string `form:"run_id" json:"run_id" xml:"run_id"`
}

// BulkStatusResponseBody is the type of the "collection" service "bulk_status"
// endpoint HTTP response body.
type BulkStatusResponseBody struct {
	Running    bool    `form:"running" json:"running" xml:"running"`
	StartedAt  *string `form:"started_at,omitempty" json:"started_at,omitempty" xml:"started_at,omitempty"`
	ClosedAt   *string `form:"closed_at,omitempty" json:"closed_at,omitempty" xml:"closed_at,omitempty"`
	Status     *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	WorkflowID *string `form:"workflow_id,omitempty" json:"workflow_id,omitempty" xml:"workflow_id,omitempty"`
	RunID      *string `form:"run_id,omitempty" json:"run_id,omitempty" xml:"run_id,omitempty"`
}

// ShowNotFoundResponseBody is the type of the "collection" service "show"
// endpoint HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// Identifier of missing collection
	ID uint `form:"id" json:"id" xml:"id"`
}

// DeleteNotFoundResponseBody is the type of the "collection" service "delete"
// endpoint HTTP response body for the "not_found" error.
type DeleteNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// Identifier of missing collection
	ID uint `form:"id" json:"id" xml:"id"`
}

// CancelNotFoundResponseBody is the type of the "collection" service "cancel"
// endpoint HTTP response body for the "not_found" error.
type CancelNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// Identifier of missing collection
	ID uint `form:"id" json:"id" xml:"id"`
}

// CancelNotRunningResponseBody is the type of the "collection" service
// "cancel" endpoint HTTP response body for the "not_running" error.
type CancelNotRunningResponseBody struct {
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

// RetryNotFoundResponseBody is the type of the "collection" service "retry"
// endpoint HTTP response body for the "not_found" error.
type RetryNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// Identifier of missing collection
	ID uint `form:"id" json:"id" xml:"id"`
}

// RetryNotRunningResponseBody is the type of the "collection" service "retry"
// endpoint HTTP response body for the "not_running" error.
type RetryNotRunningResponseBody struct {
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

// WorkflowNotFoundResponseBody is the type of the "collection" service
// "workflow" endpoint HTTP response body for the "not_found" error.
type WorkflowNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// Identifier of missing collection
	ID uint `form:"id" json:"id" xml:"id"`
}

// DownloadNotFoundResponseBody is the type of the "collection" service
// "download" endpoint HTTP response body for the "not_found" error.
type DownloadNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// Identifier of missing collection
	ID uint `form:"id" json:"id" xml:"id"`
}

// DecideNotFoundResponseBody is the type of the "collection" service "decide"
// endpoint HTTP response body for the "not_found" error.
type DecideNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// Identifier of missing collection
	ID uint `form:"id" json:"id" xml:"id"`
}

// DecideNotValidResponseBody is the type of the "collection" service "decide"
// endpoint HTTP response body for the "not_valid" error.
type DecideNotValidResponseBody struct {
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

// BulkNotAvailableResponseBody is the type of the "collection" service "bulk"
// endpoint HTTP response body for the "not_available" error.
type BulkNotAvailableResponseBody struct {
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

// BulkNotValidResponseBody is the type of the "collection" service "bulk"
// endpoint HTTP response body for the "not_valid" error.
type BulkNotValidResponseBody struct {
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

// EnduroStoredCollectionResponseBody is used to define fields on response body
// types.
type EnduroStoredCollectionResponseBody struct {
	// Identifier of collection
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of the collection
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Status of the collection
	Status string `form:"status" json:"status" xml:"status"`
	// Identifier of processing workflow
	WorkflowID *string `form:"workflow_id,omitempty" json:"workflow_id,omitempty" xml:"workflow_id,omitempty"`
	// Identifier of latest processing workflow run
	RunID *string `form:"run_id,omitempty" json:"run_id,omitempty" xml:"run_id,omitempty"`
	// Identifier of Archivematica transfer
	TransferID *string `form:"transfer_id,omitempty" json:"transfer_id,omitempty" xml:"transfer_id,omitempty"`
	// Identifier of Archivematica AIP
	AipID *string `form:"aip_id,omitempty" json:"aip_id,omitempty" xml:"aip_id,omitempty"`
	// Identifier provided by the client
	OriginalID *string `form:"original_id,omitempty" json:"original_id,omitempty" xml:"original_id,omitempty"`
	// Identifier of Archivematica pipeline
	PipelineID *string `form:"pipeline_id,omitempty" json:"pipeline_id,omitempty" xml:"pipeline_id,omitempty"`
	// Creation datetime
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// Start datetime
	StartedAt *string `form:"started_at,omitempty" json:"started_at,omitempty" xml:"started_at,omitempty"`
	// Completion datetime
	CompletedAt *string `form:"completed_at,omitempty" json:"completed_at,omitempty" xml:"completed_at,omitempty"`
}

// EnduroStoredCollectionCollectionResponseBody is used to define fields on
// response body types.
type EnduroStoredCollectionCollectionResponseBody []*EnduroStoredCollectionResponseBody

// EnduroCollectionWorkflowHistoryResponseBodyCollection is used to define
// fields on response body types.
type EnduroCollectionWorkflowHistoryResponseBodyCollection []*EnduroCollectionWorkflowHistoryResponseBody

// EnduroCollectionWorkflowHistoryResponseBody is used to define fields on
// response body types.
type EnduroCollectionWorkflowHistoryResponseBody struct {
	// Identifier of collection
	ID *uint `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Type of the event
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Contents of the event
	Details interface{} `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
}

// NewMonitorResponseBody builds the HTTP response body from the result of the
// "monitor" endpoint of the "collection" service.
func NewMonitorResponseBody(res *collectionviews.EnduroMonitorUpdateView) *MonitorResponseBody {
	body := &MonitorResponseBody{
		ID:   *res.ID,
		Type: *res.Type,
	}
	if res.Item != nil {
		body.Item = marshalCollectionviewsEnduroStoredCollectionViewToEnduroStoredCollectionResponseBody(res.Item)
	}
	return body
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "collection" service.
func NewListResponseBody(res *collection.ListResult) *ListResponseBody {
	body := &ListResponseBody{
		NextCursor: res.NextCursor,
	}
	if res.Items != nil {
		body.Items = make([]*EnduroStoredCollectionResponseBody, len(res.Items))
		for i, val := range res.Items {
			body.Items[i] = marshalCollectionEnduroStoredCollectionToEnduroStoredCollectionResponseBody(val)
		}
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "collection" service.
func NewShowResponseBody(res *collectionviews.EnduroStoredCollectionView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:          *res.ID,
		Name:        res.Name,
		Status:      *res.Status,
		WorkflowID:  res.WorkflowID,
		RunID:       res.RunID,
		TransferID:  res.TransferID,
		AipID:       res.AipID,
		OriginalID:  res.OriginalID,
		PipelineID:  res.PipelineID,
		CreatedAt:   *res.CreatedAt,
		StartedAt:   res.StartedAt,
		CompletedAt: res.CompletedAt,
	}
	return body
}

// NewWorkflowResponseBody builds the HTTP response body from the result of the
// "workflow" endpoint of the "collection" service.
func NewWorkflowResponseBody(res *collectionviews.EnduroCollectionWorkflowStatusView) *WorkflowResponseBody {
	body := &WorkflowResponseBody{
		Status: res.Status,
	}
	if res.History != nil {
		body.History = make([]*EnduroCollectionWorkflowHistoryResponseBody, len(res.History))
		for i, val := range res.History {
			body.History[i] = marshalCollectionviewsEnduroCollectionWorkflowHistoryViewToEnduroCollectionWorkflowHistoryResponseBody(val)
		}
	}
	return body
}

// NewBulkResponseBody builds the HTTP response body from the result of the
// "bulk" endpoint of the "collection" service.
func NewBulkResponseBody(res *collection.BulkResult) *BulkResponseBody {
	body := &BulkResponseBody{
		WorkflowID: res.WorkflowID,
		RunID:      res.RunID,
	}
	return body
}

// NewBulkStatusResponseBody builds the HTTP response body from the result of
// the "bulk_status" endpoint of the "collection" service.
func NewBulkStatusResponseBody(res *collection.BulkStatusResult) *BulkStatusResponseBody {
	body := &BulkStatusResponseBody{
		Running:    res.Running,
		StartedAt:  res.StartedAt,
		ClosedAt:   res.ClosedAt,
		Status:     res.Status,
		WorkflowID: res.WorkflowID,
		RunID:      res.RunID,
	}
	return body
}

// NewShowNotFoundResponseBody builds the HTTP response body from the result of
// the "show" endpoint of the "collection" service.
func NewShowNotFoundResponseBody(res *collection.CollectionNotfound) *ShowNotFoundResponseBody {
	body := &ShowNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewDeleteNotFoundResponseBody builds the HTTP response body from the result
// of the "delete" endpoint of the "collection" service.
func NewDeleteNotFoundResponseBody(res *collection.CollectionNotfound) *DeleteNotFoundResponseBody {
	body := &DeleteNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewCancelNotFoundResponseBody builds the HTTP response body from the result
// of the "cancel" endpoint of the "collection" service.
func NewCancelNotFoundResponseBody(res *collection.CollectionNotfound) *CancelNotFoundResponseBody {
	body := &CancelNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewCancelNotRunningResponseBody builds the HTTP response body from the
// result of the "cancel" endpoint of the "collection" service.
func NewCancelNotRunningResponseBody(res *goa.ServiceError) *CancelNotRunningResponseBody {
	body := &CancelNotRunningResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRetryNotFoundResponseBody builds the HTTP response body from the result
// of the "retry" endpoint of the "collection" service.
func NewRetryNotFoundResponseBody(res *collection.CollectionNotfound) *RetryNotFoundResponseBody {
	body := &RetryNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewRetryNotRunningResponseBody builds the HTTP response body from the result
// of the "retry" endpoint of the "collection" service.
func NewRetryNotRunningResponseBody(res *goa.ServiceError) *RetryNotRunningResponseBody {
	body := &RetryNotRunningResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewWorkflowNotFoundResponseBody builds the HTTP response body from the
// result of the "workflow" endpoint of the "collection" service.
func NewWorkflowNotFoundResponseBody(res *collection.CollectionNotfound) *WorkflowNotFoundResponseBody {
	body := &WorkflowNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewDownloadNotFoundResponseBody builds the HTTP response body from the
// result of the "download" endpoint of the "collection" service.
func NewDownloadNotFoundResponseBody(res *collection.CollectionNotfound) *DownloadNotFoundResponseBody {
	body := &DownloadNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewDecideNotFoundResponseBody builds the HTTP response body from the result
// of the "decide" endpoint of the "collection" service.
func NewDecideNotFoundResponseBody(res *collection.CollectionNotfound) *DecideNotFoundResponseBody {
	body := &DecideNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewDecideNotValidResponseBody builds the HTTP response body from the result
// of the "decide" endpoint of the "collection" service.
func NewDecideNotValidResponseBody(res *goa.ServiceError) *DecideNotValidResponseBody {
	body := &DecideNotValidResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewBulkNotAvailableResponseBody builds the HTTP response body from the
// result of the "bulk" endpoint of the "collection" service.
func NewBulkNotAvailableResponseBody(res *goa.ServiceError) *BulkNotAvailableResponseBody {
	body := &BulkNotAvailableResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewBulkNotValidResponseBody builds the HTTP response body from the result of
// the "bulk" endpoint of the "collection" service.
func NewBulkNotValidResponseBody(res *goa.ServiceError) *BulkNotValidResponseBody {
	body := &BulkNotValidResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListPayload builds a collection service list endpoint payload.
func NewListPayload(name *string, originalID *string, transferID *string, aipID *string, pipelineID *string, earliestCreatedTime *string, latestCreatedTime *string, status *string, cursor *string) *collection.ListPayload {
	v := &collection.ListPayload{}
	v.Name = name
	v.OriginalID = originalID
	v.TransferID = transferID
	v.AipID = aipID
	v.PipelineID = pipelineID
	v.EarliestCreatedTime = earliestCreatedTime
	v.LatestCreatedTime = latestCreatedTime
	v.Status = status
	v.Cursor = cursor

	return v
}

// NewShowPayload builds a collection service show endpoint payload.
func NewShowPayload(id uint) *collection.ShowPayload {
	v := &collection.ShowPayload{}
	v.ID = id

	return v
}

// NewDeletePayload builds a collection service delete endpoint payload.
func NewDeletePayload(id uint) *collection.DeletePayload {
	v := &collection.DeletePayload{}
	v.ID = id

	return v
}

// NewCancelPayload builds a collection service cancel endpoint payload.
func NewCancelPayload(id uint) *collection.CancelPayload {
	v := &collection.CancelPayload{}
	v.ID = id

	return v
}

// NewRetryPayload builds a collection service retry endpoint payload.
func NewRetryPayload(id uint) *collection.RetryPayload {
	v := &collection.RetryPayload{}
	v.ID = id

	return v
}

// NewWorkflowPayload builds a collection service workflow endpoint payload.
func NewWorkflowPayload(id uint) *collection.WorkflowPayload {
	v := &collection.WorkflowPayload{}
	v.ID = id

	return v
}

// NewDownloadPayload builds a collection service download endpoint payload.
func NewDownloadPayload(id uint) *collection.DownloadPayload {
	v := &collection.DownloadPayload{}
	v.ID = id

	return v
}

// NewDecidePayload builds a collection service decide endpoint payload.
func NewDecidePayload(body struct {
	// Decision option to proceed with
	Option *string `form:"option" json:"option" xml:"option"`
}, id uint) *collection.DecidePayload {
	v := &collection.DecidePayload{}
	if body.Option != nil {
		v.Option = *body.Option
	}
	v.ID = id

	return v
}

// NewBulkPayload builds a collection service bulk endpoint payload.
func NewBulkPayload(body *BulkRequestBody) *collection.BulkPayload {
	v := &collection.BulkPayload{
		Operation: *body.Operation,
		Status:    *body.Status,
	}
	if body.Size != nil {
		v.Size = *body.Size
	}
	if body.Size == nil {
		v.Size = 100
	}

	return v
}

// ValidateBulkRequestBody runs the validations defined on BulkRequestBody
func ValidateBulkRequestBody(body *BulkRequestBody) (err error) {
	if body.Operation == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("operation", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Operation != nil {
		if !(*body.Operation == "retry" || *body.Operation == "cancel" || *body.Operation == "abandon") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.operation", *body.Operation, []interface{}{"retry", "cancel", "abandon"}))
		}
	}
	if body.Status != nil {
		if !(*body.Status == "new" || *body.Status == "in progress" || *body.Status == "done" || *body.Status == "error" || *body.Status == "unknown" || *body.Status == "queued" || *body.Status == "pending" || *body.Status == "abandoned") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []interface{}{"new", "in progress", "done", "error", "unknown", "queued", "pending", "abandoned"}))
		}
	}
	return
}
