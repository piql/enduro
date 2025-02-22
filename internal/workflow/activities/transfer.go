package activities

import (
	"context"
	"fmt"
	"net/http"

	"github.com/penwern/enduro/internal/amclient"
	wferrors "github.com/penwern/enduro/internal/workflow/errors"
	"github.com/penwern/enduro/internal/workflow/manager"
)

// TransferActivity submits the transfer to Archivematica and returns its ID.
//
// This is our first interaction with Archivematica. The workflow ends here
// after authentication errors.
type TransferActivity struct {
	manager *manager.Manager
}

func NewTransferActivity(m *manager.Manager) *TransferActivity {
	return &TransferActivity{manager: m}
}

type TransferActivityParams struct {
	PipelineName       string
	TransferLocationID string
	RelPath            string
	Name               string
	ProcessingConfig   string
}

type TransferActivityResponse struct {
	TransferID      string
	PipelineVersion string
	PipelineID      string
}

func (a *TransferActivity) Execute(ctx context.Context, params *TransferActivityParams) (*TransferActivityResponse, error) {
	p, err := a.manager.Pipelines.ByName(params.PipelineName)
	if err != nil {
		return nil, wferrors.NonRetryableError(err)
	}
	amc := p.Client()

	// Transfer path should include the location UUID if defined.
	path := params.RelPath
	if params.TransferLocationID != "" {
		path = fmt.Sprintf("%s:%s", params.TransferLocationID, path)
	}

	resp, httpResp, err := amc.Package.Create(ctx, &amclient.PackageCreateRequest{
		Name:             params.Name,
		Path:             path,
		ProcessingConfig: params.ProcessingConfig,
		AutoApprove:      true,
	})
	if err != nil {
		if httpResp != nil {
			switch {
			case httpResp.StatusCode == http.StatusForbidden:
				return nil, wferrors.NonRetryableError(fmt.Errorf("authentication error: %v", err))
			}
		}
		return nil, err
	}

	return &TransferActivityResponse{
		TransferID:      resp.ID,
		PipelineVersion: httpResp.Header.Get("X-Archivematica-Version"),
		PipelineID:      httpResp.Header.Get("X-Archivematica-ID"),
	}, nil
}
