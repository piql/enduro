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
)

const (
	EarkWorkflowName = "eark-aip-generator"
	EarkWorkflowID   = "eark-aip-generator"
	EarkActivityName = "eark-activity"
)

var ErrEarkStatusUnavailable = errors.New("eark status unavailable")

type Service interface {
	Submit(context.Context) (res *goaeark.EarkResult, err error)
	Status(context.Context) (res *goaeark.EarkStatusResult, err error)
}

type earkImpl struct {
	logger logr.Logger
	cc     cadencesdk_client.Client
}

var _ Service = (*earkImpl)(nil)


func NewService(logger logr.Logger, cc cadencesdk_client.Client) *earkImpl {
	return &earkImpl{
		logger:        logger,
		cc:            cc,
	}
}

func (s *earkImpl) Submit(ctx context.Context) (*goaeark.EarkResult, error) {
	opts := cadencesdk_client.StartWorkflowOptions{
		ID:                              EarkWorkflowID,
		WorkflowIDReusePolicy:           cadencesdk_client.WorkflowIDReusePolicyAllowDuplicate,
		TaskList:                        cadence.GlobalTaskListName,
		DecisionTaskStartToCloseTimeout: time.Second * 10,
		ExecutionStartToCloseTimeout:    time.Hour,
	}
	exec, err := s.cc.StartWorkflow(ctx, opts, EarkWorkflowName)
	if err != nil {
		switch err := err.(type) {
		case *cadencesdk_gen_shared.WorkflowExecutionAlreadyStartedError:
			return nil, goaeark.MakeNotAvailable(
				fmt.Errorf("error starting eark - operation is already in progress (workflowID=%s runID=%s)",
					EarkWorkflowID, *err.RunId))
		default:
			s.logger.Info("error starting eark", "err", err)
			return nil, fmt.Errorf("error starting eark")
		}
	}
	result := &goaeark.EarkResult{
		WorkflowID: exec.ID,
		RunID:      exec.RunID,
	}
	return result, nil
}

func (s *earkImpl) Status(ctx context.Context) (*goaeark.EarkStatusResult, error) {
	result := &goaeark.EarkStatusResult{}
	resp, err := s.cc.DescribeWorkflowExecution(ctx, EarkWorkflowID, "")
	if err != nil {
		switch err := err.(type) {
		case *cadencesdk_gen_shared.EntityNotExistsError:
			return result, nil
		default:
			s.logger.Info("error retrieving workflow", "err", err)
			return nil, ErrEarkStatusUnavailable
		}
	}
	if resp.WorkflowExecutionInfo == nil {
		s.logger.Info("error retrieving workflow execution details")
		return nil, ErrEarkStatusUnavailable
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
