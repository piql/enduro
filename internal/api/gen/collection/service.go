// Code generated by goa v3.5.5, DO NOT EDIT.
//
// collection service
//
// Command:
// $ goa gen github.com/penwern/enduro/internal/api/design -o internal/api

package collection

import (
	"context"

	collectionviews "github.com/penwern/enduro/internal/api/gen/collection/views"
	goa "goa.design/goa/v3/pkg"
)

// The collection service manages packages being transferred to Archivematica.
type Service interface {
	// Monitor implements monitor.
	Monitor(context.Context, MonitorServerStream) (err error)
	// List all stored collections
	List(context.Context, *ListPayload) (res *ListResult, err error)
	// Show collection by ID
	Show(context.Context, *ShowPayload) (res *EnduroStoredCollection, err error)
	// Delete collection by ID
	Delete(context.Context, *DeletePayload) (err error)
	// Cancel collection processing by ID
	Cancel(context.Context, *CancelPayload) (err error)
	// Retry collection processing by ID
	Retry(context.Context, *RetryPayload) (err error)
	// Retrieve workflow status by ID
	Workflow(context.Context, *WorkflowPayload) (res *EnduroCollectionWorkflowStatus, err error)
	// Download collection by ID
	Download(context.Context, *DownloadPayload) (res []byte, err error)
	// Make decision for a pending collection by ID
	Decide(context.Context, *DecidePayload) (err error)
	// Bulk operations (retry, cancel...).
	Bulk(context.Context, *BulkPayload) (res *BulkResult, err error)
	// Retrieve status of current bulk operation.
	BulkStatus(context.Context) (res *BulkStatusResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "collection"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [11]string{"monitor", "list", "show", "delete", "cancel", "retry", "workflow", "download", "decide", "bulk", "bulk_status"}

// MonitorServerStream is the interface a "monitor" endpoint server stream must
// satisfy.
type MonitorServerStream interface {
	// Send streams instances of "EnduroMonitorUpdate".
	Send(*EnduroMonitorUpdate) error
	// Close closes the stream.
	Close() error
}

// MonitorClientStream is the interface a "monitor" endpoint client stream must
// satisfy.
type MonitorClientStream interface {
	// Recv reads instances of "EnduroMonitorUpdate" from the stream.
	Recv() (*EnduroMonitorUpdate, error)
}

// EnduroMonitorUpdate is the result type of the collection service monitor
// method.
type EnduroMonitorUpdate struct {
	// Identifier of collection
	ID uint
	// Type of the event
	Type string
	// Collection
	Item *EnduroStoredCollection
}

// ListPayload is the payload type of the collection service list method.
type ListPayload struct {
	Name                *string
	OriginalID          *string
	TransferID          *string
	AipID               *string
	PipelineID          *string
	EarliestCreatedTime *string
	LatestCreatedTime   *string
	Status              *string
	// Pagination cursor
	Cursor *string
}

// ListResult is the result type of the collection service list method.
type ListResult struct {
	Items      EnduroStoredCollectionCollection
	NextCursor *string
}

// ShowPayload is the payload type of the collection service show method.
type ShowPayload struct {
	// Identifier of collection to show
	ID uint
}

// EnduroStoredCollection is the result type of the collection service show
// method.
type EnduroStoredCollection struct {
	// Identifier of collection
	ID uint
	// Name of the collection
	Name *string
	// Status of the collection
	Status string
	// Identifier of processing workflow
	WorkflowID *string
	// Identifier of latest processing workflow run
	RunID *string
	// Identifier of Archivematica transfer
	TransferID *string
	// Identifier of Archivematica AIP
	AipID *string
	// Identifier provided by the client
	OriginalID *string
	// Identifier of Archivematica pipeline
	PipelineID *string
	// Creation datetime
	CreatedAt string
	// Start datetime
	StartedAt *string
	// Completion datetime
	CompletedAt *string
}

// DeletePayload is the payload type of the collection service delete method.
type DeletePayload struct {
	// Identifier of collection to delete
	ID uint
}

// CancelPayload is the payload type of the collection service cancel method.
type CancelPayload struct {
	// Identifier of collection to remove
	ID uint
}

// RetryPayload is the payload type of the collection service retry method.
type RetryPayload struct {
	// Identifier of collection to retry
	ID uint
}

// WorkflowPayload is the payload type of the collection service workflow
// method.
type WorkflowPayload struct {
	// Identifier of collection to look up
	ID uint
}

// EnduroCollectionWorkflowStatus is the result type of the collection service
// workflow method.
type EnduroCollectionWorkflowStatus struct {
	Status  *string
	History EnduroCollectionWorkflowHistoryCollection
}

// DownloadPayload is the payload type of the collection service download
// method.
type DownloadPayload struct {
	// Identifier of collection to look up
	ID uint
}

// DecidePayload is the payload type of the collection service decide method.
type DecidePayload struct {
	// Identifier of collection to look up
	ID uint
	// Decision option to proceed with
	Option string
}

// BulkPayload is the payload type of the collection service bulk method.
type BulkPayload struct {
	Operation string
	Status    string
	Size      uint
}

// BulkResult is the result type of the collection service bulk method.
type BulkResult struct {
	WorkflowID string
	RunID      string
}

// BulkStatusResult is the result type of the collection service bulk_status
// method.
type BulkStatusResult struct {
	Running    bool
	StartedAt  *string
	ClosedAt   *string
	Status     *string
	WorkflowID *string
	RunID      *string
}

type EnduroStoredCollectionCollection []*EnduroStoredCollection

type EnduroCollectionWorkflowHistoryCollection []*EnduroCollectionWorkflowHistory

// WorkflowHistoryEvent describes a history event in Cadence.
type EnduroCollectionWorkflowHistory struct {
	// Identifier of collection
	ID *uint
	// Type of the event
	Type *string
	// Contents of the event
	Details interface{}
}

// Collection not found.
type CollectionNotfound struct {
	// Message of error
	Message string
	// Identifier of missing collection
	ID uint
}

// Error returns an error description.
func (e *CollectionNotfound) Error() string {
	return "Collection not found."
}

// ErrorName returns "CollectionNotfound".
func (e *CollectionNotfound) ErrorName() string {
	return e.Message
}

// MakeNotRunning builds a goa.ServiceError from an error.
func MakeNotRunning(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not_running",
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

// MakeNotAvailable builds a goa.ServiceError from an error.
func MakeNotAvailable(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not_available",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewEnduroMonitorUpdate initializes result type EnduroMonitorUpdate from
// viewed result type EnduroMonitorUpdate.
func NewEnduroMonitorUpdate(vres *collectionviews.EnduroMonitorUpdate) *EnduroMonitorUpdate {
	return newEnduroMonitorUpdate(vres.Projected)
}

// NewViewedEnduroMonitorUpdate initializes viewed result type
// EnduroMonitorUpdate from result type EnduroMonitorUpdate using the given
// view.
func NewViewedEnduroMonitorUpdate(res *EnduroMonitorUpdate, view string) *collectionviews.EnduroMonitorUpdate {
	p := newEnduroMonitorUpdateView(res)
	return &collectionviews.EnduroMonitorUpdate{Projected: p, View: "default"}
}

// NewEnduroStoredCollection initializes result type EnduroStoredCollection
// from viewed result type EnduroStoredCollection.
func NewEnduroStoredCollection(vres *collectionviews.EnduroStoredCollection) *EnduroStoredCollection {
	return newEnduroStoredCollection(vres.Projected)
}

// NewViewedEnduroStoredCollection initializes viewed result type
// EnduroStoredCollection from result type EnduroStoredCollection using the
// given view.
func NewViewedEnduroStoredCollection(res *EnduroStoredCollection, view string) *collectionviews.EnduroStoredCollection {
	p := newEnduroStoredCollectionView(res)
	return &collectionviews.EnduroStoredCollection{Projected: p, View: "default"}
}

// NewEnduroCollectionWorkflowStatus initializes result type
// EnduroCollectionWorkflowStatus from viewed result type
// EnduroCollectionWorkflowStatus.
func NewEnduroCollectionWorkflowStatus(vres *collectionviews.EnduroCollectionWorkflowStatus) *EnduroCollectionWorkflowStatus {
	return newEnduroCollectionWorkflowStatus(vres.Projected)
}

// NewViewedEnduroCollectionWorkflowStatus initializes viewed result type
// EnduroCollectionWorkflowStatus from result type
// EnduroCollectionWorkflowStatus using the given view.
func NewViewedEnduroCollectionWorkflowStatus(res *EnduroCollectionWorkflowStatus, view string) *collectionviews.EnduroCollectionWorkflowStatus {
	p := newEnduroCollectionWorkflowStatusView(res)
	return &collectionviews.EnduroCollectionWorkflowStatus{Projected: p, View: "default"}
}

// newEnduroMonitorUpdate converts projected type EnduroMonitorUpdate to
// service type EnduroMonitorUpdate.
func newEnduroMonitorUpdate(vres *collectionviews.EnduroMonitorUpdateView) *EnduroMonitorUpdate {
	res := &EnduroMonitorUpdate{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	if vres.Item != nil {
		res.Item = newEnduroStoredCollection(vres.Item)
	}
	return res
}

// newEnduroMonitorUpdateView projects result type EnduroMonitorUpdate to
// projected type EnduroMonitorUpdateView using the "default" view.
func newEnduroMonitorUpdateView(res *EnduroMonitorUpdate) *collectionviews.EnduroMonitorUpdateView {
	vres := &collectionviews.EnduroMonitorUpdateView{
		ID:   &res.ID,
		Type: &res.Type,
	}
	if res.Item != nil {
		vres.Item = newEnduroStoredCollectionView(res.Item)
	}
	return vres
}

// newEnduroStoredCollection converts projected type EnduroStoredCollection to
// service type EnduroStoredCollection.
func newEnduroStoredCollection(vres *collectionviews.EnduroStoredCollectionView) *EnduroStoredCollection {
	res := &EnduroStoredCollection{
		Name:        vres.Name,
		WorkflowID:  vres.WorkflowID,
		RunID:       vres.RunID,
		TransferID:  vres.TransferID,
		AipID:       vres.AipID,
		OriginalID:  vres.OriginalID,
		PipelineID:  vres.PipelineID,
		StartedAt:   vres.StartedAt,
		CompletedAt: vres.CompletedAt,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Status != nil {
		res.Status = *vres.Status
	}
	if vres.CreatedAt != nil {
		res.CreatedAt = *vres.CreatedAt
	}
	if vres.Status == nil {
		res.Status = "new"
	}
	return res
}

// newEnduroStoredCollectionView projects result type EnduroStoredCollection to
// projected type EnduroStoredCollectionView using the "default" view.
func newEnduroStoredCollectionView(res *EnduroStoredCollection) *collectionviews.EnduroStoredCollectionView {
	vres := &collectionviews.EnduroStoredCollectionView{
		ID:          &res.ID,
		Name:        res.Name,
		Status:      &res.Status,
		WorkflowID:  res.WorkflowID,
		RunID:       res.RunID,
		TransferID:  res.TransferID,
		AipID:       res.AipID,
		OriginalID:  res.OriginalID,
		PipelineID:  res.PipelineID,
		CreatedAt:   &res.CreatedAt,
		StartedAt:   res.StartedAt,
		CompletedAt: res.CompletedAt,
	}
	return vres
}

// newEnduroCollectionWorkflowStatus converts projected type
// EnduroCollectionWorkflowStatus to service type
// EnduroCollectionWorkflowStatus.
func newEnduroCollectionWorkflowStatus(vres *collectionviews.EnduroCollectionWorkflowStatusView) *EnduroCollectionWorkflowStatus {
	res := &EnduroCollectionWorkflowStatus{
		Status: vres.Status,
	}
	if vres.History != nil {
		res.History = newEnduroCollectionWorkflowHistoryCollection(vres.History)
	}
	return res
}

// newEnduroCollectionWorkflowStatusView projects result type
// EnduroCollectionWorkflowStatus to projected type
// EnduroCollectionWorkflowStatusView using the "default" view.
func newEnduroCollectionWorkflowStatusView(res *EnduroCollectionWorkflowStatus) *collectionviews.EnduroCollectionWorkflowStatusView {
	vres := &collectionviews.EnduroCollectionWorkflowStatusView{
		Status: res.Status,
	}
	if res.History != nil {
		vres.History = newEnduroCollectionWorkflowHistoryCollectionView(res.History)
	}
	return vres
}

// newEnduroCollectionWorkflowHistoryCollection converts projected type
// EnduroCollectionWorkflowHistoryCollection to service type
// EnduroCollectionWorkflowHistoryCollection.
func newEnduroCollectionWorkflowHistoryCollection(vres collectionviews.EnduroCollectionWorkflowHistoryCollectionView) EnduroCollectionWorkflowHistoryCollection {
	res := make(EnduroCollectionWorkflowHistoryCollection, len(vres))
	for i, n := range vres {
		res[i] = newEnduroCollectionWorkflowHistory(n)
	}
	return res
}

// newEnduroCollectionWorkflowHistoryCollectionView projects result type
// EnduroCollectionWorkflowHistoryCollection to projected type
// EnduroCollectionWorkflowHistoryCollectionView using the "default" view.
func newEnduroCollectionWorkflowHistoryCollectionView(res EnduroCollectionWorkflowHistoryCollection) collectionviews.EnduroCollectionWorkflowHistoryCollectionView {
	vres := make(collectionviews.EnduroCollectionWorkflowHistoryCollectionView, len(res))
	for i, n := range res {
		vres[i] = newEnduroCollectionWorkflowHistoryView(n)
	}
	return vres
}

// newEnduroCollectionWorkflowHistory converts projected type
// EnduroCollectionWorkflowHistory to service type
// EnduroCollectionWorkflowHistory.
func newEnduroCollectionWorkflowHistory(vres *collectionviews.EnduroCollectionWorkflowHistoryView) *EnduroCollectionWorkflowHistory {
	res := &EnduroCollectionWorkflowHistory{
		ID:      vres.ID,
		Type:    vres.Type,
		Details: vres.Details,
	}
	return res
}

// newEnduroCollectionWorkflowHistoryView projects result type
// EnduroCollectionWorkflowHistory to projected type
// EnduroCollectionWorkflowHistoryView using the "default" view.
func newEnduroCollectionWorkflowHistoryView(res *EnduroCollectionWorkflowHistory) *collectionviews.EnduroCollectionWorkflowHistoryView {
	vres := &collectionviews.EnduroCollectionWorkflowHistoryView{
		ID:      res.ID,
		Type:    res.Type,
		Details: res.Details,
	}
	return vres
}
