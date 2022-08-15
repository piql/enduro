package eark

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-logr/logr"
	cadencesdk_gen_shared "go.uber.org/cadence/.gen/go/shared"
	cadencesdk_client "go.uber.org/cadence/client"

	goaeark "github.com/penwern/enduro/internal/api/gen/eark"
	"github.com/penwern/enduro/internal/cadence"
	"github.com/penwern/enduro/internal/collection"
	"github.com/penwern/enduro/internal/validation"
)

var ErrearkStatusUnavailable = errors.New("eark status unavailable")

type Service interface {
	Submit(context.Context/**, *goaeark.SubmitPayload*/) (res *goaeark.earkResult, err error)
	Status(context.Context) (res *goaeark.earkStatusResult, err error)
	Hints(context.Context) (res *goaeark.earkHintsResult, err error)
	InitProcessingWorkflow(ctx context.Context, req *collection.ProcessingWorkflowRequest) error
}

type earkImpl struct {
	logger logr.Logger
	cc     cadencesdk_client.Client

	// A list of completedDirs reported by the watcher configuration. This is
	// used to provide the user with possible known values.
	completedDirs []string
}

var _ Service = (*earkImpl)(nil)

func NewService(logger logr.Logger, cc cadencesdk_client.Client, completedDirs []string) *earkImpl {
	return &earkImpl{
		logger:        logger,
		cc:            cc,
		completedDirs: completedDirs,
	}
}

func (s *earkImpl) Submit(ctx context.Context/**, payload *goaeark.SubmitPayload*/) (*goaeark.earkResult, error) {
	/**
	if payload.Path == "" {
		return nil, goaeark.MakeNotValid(errors.New("error starting eark - path is empty"))
	}
	input := earkWorkflowInput{
		Path: payload.Path,
	}
	if payload.Pipeline != nil {
		input.PipelineName = *payload.Pipeline
	}
	if payload.ProcessingConfig != nil {
		input.ProcessingConfig = *payload.ProcessingConfig
	}
	if payload.CompletedDir != nil {
		input.CompletedDir = *payload.CompletedDir
	}
	if payload.RetentionPeriod != nil {
		dur, err := time.ParseDuration(*payload.RetentionPeriod)
		if err != nil {
			return nil, goaeark.MakeNotValid(errors.New("error starting eark - retention period format is invalid"))
		}
		input.RetentionPeriod = &dur
	}
	*/
	opts := cadencesdk_client.StartWorkflowOptions{
		ID:                              earkWorkflowID,
		WorkflowIDReusePolicy:           cadencesdk_client.WorkflowIDReusePolicyAllowDuplicate,
		TaskList:                        cadence.GlobalTaskListName,
		DecisionTaskStartToCloseTimeout: time.Second * 10,
		ExecutionStartToCloseTimeout:    time.Hour,
	}
	exec, err := s.cc.StartWorkflow(ctx, opts, earkWorkflowName/**, input*/)
	if err != nil {
		switch err := err.(type) {
		case *cadencesdk_gen_shared.WorkflowExecutionAlreadyStartedError:
			return nil, goaeark.MakeNotAvailable(
				fmt.Errorf("error starting eark - operation is already in progress (workflowID=%s runID=%s)",
					earkWorkflowID, *err.RunId))
		default:
			s.logger.Info("error starting eark", "err", err)
			return nil, fmt.Errorf("error starting eark")
		}
	}
	result := &goaeark.earkResult{
		WorkflowID: exec.ID,
		RunID:      exec.RunID,
	}
	return result, nil
}

func (s *earkImpl) Status(ctx context.Context) (*goaeark.earkStatusResult, error) {
	result := &goaeark.earkStatusResult{}
	resp, err := s.cc.DescribeWorkflowExecution(ctx, earkWorkflowID, "")
	if err != nil {
		switch err := err.(type) {
		case *cadencesdk_gen_shared.EntityNotExistsError:
			return result, nil
		default:
			s.logger.Info("error retrieving workflow", "err", err)
			return nil, ErrearkStatusUnavailable
		}
	}
	if resp.WorkflowExecutionInfo == nil {
		s.logger.Info("error retrieving workflow execution details")
		return nil, ErrearkStatusUnavailable
	}
	result.WorkflowID = resp.WorkflowExecutionInfo.Execution.WorkflowId
	result.RunID = resp.WorkflowExecutionInfo.Execution.RunId
	if resp.WorkflowExecutionInfo.CloseStatus != nil {
		st := strings.ToLower(resp.WorkflowExecutionInfo.CloseStatus.String())
		result.Status = &st
		return result, nil
	}
	result.Running = true
	return result, nil
}

func (s *earkImpl) Hints(ctx context.Context) (*goaeark.earkHintsResult, error) {
	result := &goaeark.earkHintsResult{
		CompletedDirs: s.completedDirs,
	}
	return result, nil
}

func (s *earkImpl) InitProcessingWorkflow(ctx context.Context, req *collection.ProcessingWorkflowRequest) error {
	req.ValidationConfig = validation.Config{}
	err := collection.InitProcessingWorkflow(ctx, s.cc, req)
	if err != nil {
		s.logger.Error(err, "Error initializing processing workflow.")
	}
	return err
}
