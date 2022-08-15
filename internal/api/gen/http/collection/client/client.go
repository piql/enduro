// Code generated by goa v3.5.5, DO NOT EDIT.
//
// collection client HTTP transport
//
// Command:
// $ goa gen github.com/penwern/enduro/internal/api/design -o internal/api

package client

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the collection service endpoint HTTP clients.
type Client struct {
	// Monitor Doer is the HTTP client used to make requests to the monitor
	// endpoint.
	MonitorDoer goahttp.Doer

	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// Show Doer is the HTTP client used to make requests to the show endpoint.
	ShowDoer goahttp.Doer

	// Delete Doer is the HTTP client used to make requests to the delete endpoint.
	DeleteDoer goahttp.Doer

	// Cancel Doer is the HTTP client used to make requests to the cancel endpoint.
	CancelDoer goahttp.Doer

	// Retry Doer is the HTTP client used to make requests to the retry endpoint.
	RetryDoer goahttp.Doer

	// Workflow Doer is the HTTP client used to make requests to the workflow
	// endpoint.
	WorkflowDoer goahttp.Doer

	// Download Doer is the HTTP client used to make requests to the download
	// endpoint.
	DownloadDoer goahttp.Doer

	// Decide Doer is the HTTP client used to make requests to the decide endpoint.
	DecideDoer goahttp.Doer

	// Bulk Doer is the HTTP client used to make requests to the bulk endpoint.
	BulkDoer goahttp.Doer

	// BulkStatus Doer is the HTTP client used to make requests to the bulk_status
	// endpoint.
	BulkStatusDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme     string
	host       string
	encoder    func(*http.Request) goahttp.Encoder
	decoder    func(*http.Response) goahttp.Decoder
	dialer     goahttp.Dialer
	configurer *ConnConfigurer
}

// NewClient instantiates HTTP clients for all the collection service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
	dialer goahttp.Dialer,
	cfn *ConnConfigurer,
) *Client {
	if cfn == nil {
		cfn = &ConnConfigurer{}
	}
	return &Client{
		MonitorDoer:         doer,
		ListDoer:            doer,
		ShowDoer:            doer,
		DeleteDoer:          doer,
		CancelDoer:          doer,
		RetryDoer:           doer,
		WorkflowDoer:        doer,
		DownloadDoer:        doer,
		DecideDoer:          doer,
		BulkDoer:            doer,
		BulkStatusDoer:      doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
		dialer:              dialer,
		configurer:          cfn,
	}
}

// Monitor returns an endpoint that makes HTTP requests to the collection
// service monitor server.
func (c *Client) Monitor() goa.Endpoint {
	var (
		decodeResponse = DecodeMonitorResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildMonitorRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("collection", "monitor", err)
		}
		if c.configurer.MonitorFn != nil {
			conn = c.configurer.MonitorFn(conn, cancel)
		}
		go func() {
			<-ctx.Done()
			conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, "client closing connection"),
				time.Now().Add(time.Second),
			)
			conn.Close()
		}()
		stream := &MonitorClientStream{conn: conn}
		return stream, nil
	}
}

// List returns an endpoint that makes HTTP requests to the collection service
// list server.
func (c *Client) List() goa.Endpoint {
	var (
		encodeRequest  = EncodeListRequest(c.encoder)
		decodeResponse = DecodeListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("collection", "list", err)
		}
		return decodeResponse(resp)
	}
}

// Show returns an endpoint that makes HTTP requests to the collection service
// show server.
func (c *Client) Show() goa.Endpoint {
	var (
		decodeResponse = DecodeShowResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildShowRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ShowDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("collection", "show", err)
		}
		return decodeResponse(resp)
	}
}

// Delete returns an endpoint that makes HTTP requests to the collection
// service delete server.
func (c *Client) Delete() goa.Endpoint {
	var (
		decodeResponse = DecodeDeleteResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("collection", "delete", err)
		}
		return decodeResponse(resp)
	}
}

// Cancel returns an endpoint that makes HTTP requests to the collection
// service cancel server.
func (c *Client) Cancel() goa.Endpoint {
	var (
		decodeResponse = DecodeCancelResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCancelRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CancelDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("collection", "cancel", err)
		}
		return decodeResponse(resp)
	}
}

// Retry returns an endpoint that makes HTTP requests to the collection service
// retry server.
func (c *Client) Retry() goa.Endpoint {
	var (
		decodeResponse = DecodeRetryResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRetryRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RetryDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("collection", "retry", err)
		}
		return decodeResponse(resp)
	}
}

// Workflow returns an endpoint that makes HTTP requests to the collection
// service workflow server.
func (c *Client) Workflow() goa.Endpoint {
	var (
		decodeResponse = DecodeWorkflowResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildWorkflowRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.WorkflowDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("collection", "workflow", err)
		}
		return decodeResponse(resp)
	}
}

// Download returns an endpoint that makes HTTP requests to the collection
// service download server.
func (c *Client) Download() goa.Endpoint {
	var (
		decodeResponse = DecodeDownloadResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDownloadRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DownloadDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("collection", "download", err)
		}
		return decodeResponse(resp)
	}
}

// Decide returns an endpoint that makes HTTP requests to the collection
// service decide server.
func (c *Client) Decide() goa.Endpoint {
	var (
		encodeRequest  = EncodeDecideRequest(c.encoder)
		decodeResponse = DecodeDecideResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDecideRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DecideDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("collection", "decide", err)
		}
		return decodeResponse(resp)
	}
}

// Bulk returns an endpoint that makes HTTP requests to the collection service
// bulk server.
func (c *Client) Bulk() goa.Endpoint {
	var (
		encodeRequest  = EncodeBulkRequest(c.encoder)
		decodeResponse = DecodeBulkResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildBulkRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.BulkDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("collection", "bulk", err)
		}
		return decodeResponse(resp)
	}
}

// BulkStatus returns an endpoint that makes HTTP requests to the collection
// service bulk_status server.
func (c *Client) BulkStatus() goa.Endpoint {
	var (
		decodeResponse = DecodeBulkStatusResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildBulkStatusRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.BulkStatusDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("collection", "bulk_status", err)
		}
		return decodeResponse(resp)
	}
}
