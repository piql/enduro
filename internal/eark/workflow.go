package eark

import (
	"context"
	"io/ioutil"
	"time"

	cadencesdk_workflow "go.uber.org/cadence/workflow"

	"github.com/penwern/enduro/internal/collection"
	wferrors "github.com/penwern/enduro/internal/workflow/errors"
)

const (
	earkWorkflowName = "eark-aip-generator"
	earkWorkflowID   = "eark-aip-generator"
	earkActivityName = "eark-activity"
)

type earkProgress struct {
	CurrentID uint
}

type earkWorkflowInput struct {
	Path             string
	PipelineName     string
	ProcessingConfig string
	CompletedDir     string
	RetentionPeriod  *time.Duration
}

func earkWorkflow(ctx cadencesdk_workflow.Context, params earkWorkflowInput) error {
	opts := cadencesdk_workflow.WithActivityOptions(ctx, cadencesdk_workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Hour * 24 * 365,
		StartToCloseTimeout:    time.Hour * 24 * 365,
		WaitForCancellation:    true,
	})
	return cadencesdk_workflow.ExecuteActivity(opts, earkActivityName, params).Get(opts, nil)
}

type earkActivity struct {
	earksvc Service
}

func NewearkActivity(earksvc Service) *earkActivity {
	return &earkActivity{
		earksvc: earksvc,
	}
}

func (a *earkActivity) Execute(ctx context.Context, params earkWorkflowInput) error {
	files, err := ioutil.ReadDir(params.Path)
	if err != nil {
		return wferrors.NonRetryableError(err)
	}
	pipelines := []string{}
	if params.PipelineName != "" {
		pipelines = append(pipelines, params.PipelineName)
	}
	for _, file := range files {
		req := collection.ProcessingWorkflowRequest{
			BatchDir:         params.Path,
			Key:              file.Name(),
			IsDir:            file.IsDir(),
			PipelineNames:    pipelines,
			ProcessingConfig: params.ProcessingConfig,
			CompletedDir:     params.CompletedDir,
			RetentionPeriod:  params.RetentionPeriod,
		}
		_ = a.earksvc.InitProcessingWorkflow(ctx, &req)
	}
	return nil
}
