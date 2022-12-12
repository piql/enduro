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
	AIPWorkflowName = "eark-aip-generator"
	AIPWorkflowID   = "eark-aip-generator"
	DIPWorkflowName = "eark-dip-generator"
	DIPWorkflowID = "eark-dip-generator"
	EarkActivityName = "eark-activity"
	
)

var ErrEarkStatusUnavailable = errors.New("eark status unavailable")

type Service interface {
	GenEarkAips(context.Context) (res *goaeark.EarkAIPResult, err error)
	AipGenStatus(context.Context) (res *goaeark.EarkAIPStatusResult, err error)
	GenEarkDips(context.Context) (res *goaeark.EarkDIPResult, err error)
	DipGenStatus(context.Context) (res *goaeark.EarkDIPStatusResult, err error)
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

func (s *earkImpl) GenEarkAips(ctx context.Context) (*goaeark.EarkAIPResult, error) {
	opts := cadencesdk_client.StartWorkflowOptions{
		ID:                              AIPWorkflowID,
		WorkflowIDReusePolicy:           cadencesdk_client.WorkflowIDReusePolicyAllowDuplicate,
		TaskList:                        cadence.GlobalTaskListName,
		DecisionTaskStartToCloseTimeout: time.Second * 10,
		ExecutionStartToCloseTimeout:    time.Hour,
	}
	exec, err := s.cc.StartWorkflow(ctx, opts, AIPWorkflowName)
	if err != nil {
		switch err := err.(type) {
		case *cadencesdk_gen_shared.WorkflowExecutionAlreadyStartedError:
			return nil, goaeark.MakeNotAvailable(
				fmt.Errorf("error starting eark - operation is already in progress (workflowID=%s runID=%s)",
					AIPWorkflowID, *err.RunId))
		default:
			s.logger.Info("error starting eark", "err", err)
			return nil, fmt.Errorf("error starting eark")
		}
	}
	result := &goaeark.EarkAIPResult{
		WorkflowID: exec.ID,
		RunID:      exec.RunID,
	}
	return result, nil
}

func (s *earkImpl) AipGenStatus(ctx context.Context) (*goaeark.EarkAIPStatusResult, error) {
	result := &goaeark.EarkAIPStatusResult{}
	resp, err := s.cc.DescribeWorkflowExecution(ctx, AIPWorkflowID, "")
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

func (s *earkImpl) GenEarkDips(ctx context.Context) (*goaeark.EarkDIPResult, error) {
	opts := cadencesdk_client.StartWorkflowOptions{
		ID:                              DIPWorkflowID,
		WorkflowIDReusePolicy:           cadencesdk_client.WorkflowIDReusePolicyAllowDuplicate,
		TaskList:                        cadence.GlobalTaskListName,
		DecisionTaskStartToCloseTimeout: time.Second * 10,
		ExecutionStartToCloseTimeout:    time.Hour,
	}
	exec, err := s.cc.StartWorkflow(ctx, opts, DIPWorkflowName)
	if err != nil {
		switch err := err.(type) {
		case *cadencesdk_gen_shared.WorkflowExecutionAlreadyStartedError:
			return nil, goaeark.MakeNotAvailable(
				fmt.Errorf("error starting eark - operation is already in progress (workflowID=%s runID=%s)",
					DIPWorkflowID, *err.RunId))
		default:
			s.logger.Info("error starting eark", "err", err)
			return nil, fmt.Errorf("error starting eark")
		}
	}
	result := &goaeark.EarkDIPResult{
		WorkflowID: exec.ID,
		RunID:      exec.RunID,
	}
	return result, nil
}

func (s *earkImpl) DipGenStatus(ctx context.Context) (*goaeark.EarkDIPStatusResult, error) {
	result := &goaeark.EarkDIPStatusResult{}
	resp, err := s.cc.DescribeWorkflowExecution(ctx, DIPWorkflowID, "")
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