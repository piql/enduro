package eark

import (
	//"context"
	//"io/ioutil"
	"time"

	cadencesdk_workflow "go.uber.org/cadence/workflow"

	//"github.com/penwern/enduro/internal/collection"
	//wferrors "github.com/penwern/enduro/internal/workflow/errors"
)

const (
	EarkWorkflowName = "eark-aip-generator"
	EarkWorkflowID   = "eark-aip-generator"
	EarkActivityName = "eark-activity"
)

type EarkProgress struct {
	CurrentID uint
}
/**
type EarkWorkflowInput struct {
	Path             string
	PipelineName     string
	ProcessingConfig string
	CompletedDir     string
	RetentionPeriod  *time.Duration
}
*/
func EarkWorkflow(ctx cadencesdk_workflow.Context/**, params EarkWorkflowInput*/) error {
	opts := cadencesdk_workflow.WithActivityOptions(ctx, cadencesdk_workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Hour * 24 * 365,
		StartToCloseTimeout:    time.Hour * 24 * 365,
		WaitForCancellation:    true,
	})
	return cadencesdk_workflow.ExecuteActivity(opts, EarkActivityName/**, params*/).Get(opts, nil)
}

type EarkActivity struct {
	earksvc Service
}

func NewEarkActivity(earksvc Service) *EarkActivity {
	return &EarkActivity{
		earksvc: earksvc,
	}
}
/**
func (a *EarkActivity) Execute(ctx context.Context/**, params EarkWorkflowInput) error {
	/**
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
*/