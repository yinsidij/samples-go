// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package mocks

import cadence "go.uber.org/cadence/.gen/go/cadence"
import mock "github.com/stretchr/testify/mock"
import shared "go.uber.org/cadence/.gen/go/shared"
import thrift "github.com/uber/tchannel-go/thrift"

// TChanWorkflowService is an autogenerated mock type for the TChanWorkflowService type
type TChanWorkflowService struct {
	mock.Mock
}

// DeprecateDomain provides a mock function with given fields: ctx, deprecateRequest
func (_m *TChanWorkflowService) DeprecateDomain(ctx thrift.Context, deprecateRequest *shared.DeprecateDomainRequest) error {
	ret := _m.Called(ctx, deprecateRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.DeprecateDomainRequest) error); ok {
		r0 = rf(ctx, deprecateRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeDomain provides a mock function with given fields: ctx, describeRequest
func (_m *TChanWorkflowService) DescribeDomain(ctx thrift.Context, describeRequest *shared.DescribeDomainRequest) (*shared.DescribeDomainResponse, error) {
	ret := _m.Called(ctx, describeRequest)

	var r0 *shared.DescribeDomainResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.DescribeDomainRequest) *shared.DescribeDomainResponse); ok {
		r0 = rf(ctx, describeRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.DescribeDomainResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.DescribeDomainRequest) error); ok {
		r1 = rf(ctx, describeRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkflowExecutionHistory provides a mock function with given fields: ctx, getRequest
func (_m *TChanWorkflowService) GetWorkflowExecutionHistory(ctx thrift.Context, getRequest *shared.GetWorkflowExecutionHistoryRequest) (*shared.GetWorkflowExecutionHistoryResponse, error) {
	ret := _m.Called(ctx, getRequest)

	var r0 *shared.GetWorkflowExecutionHistoryResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.GetWorkflowExecutionHistoryRequest) *shared.GetWorkflowExecutionHistoryResponse); ok {
		r0 = rf(ctx, getRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.GetWorkflowExecutionHistoryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.GetWorkflowExecutionHistoryRequest) error); ok {
		r1 = rf(ctx, getRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClosedWorkflowExecutions provides a mock function with given fields: ctx, listRequest
func (_m *TChanWorkflowService) ListClosedWorkflowExecutions(ctx thrift.Context, listRequest *shared.ListClosedWorkflowExecutionsRequest) (*shared.ListClosedWorkflowExecutionsResponse, error) {
	ret := _m.Called(ctx, listRequest)

	var r0 *shared.ListClosedWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.ListClosedWorkflowExecutionsRequest) *shared.ListClosedWorkflowExecutionsResponse); ok {
		r0 = rf(ctx, listRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ListClosedWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.ListClosedWorkflowExecutionsRequest) error); ok {
		r1 = rf(ctx, listRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOpenWorkflowExecutions provides a mock function with given fields: ctx, listRequest
func (_m *TChanWorkflowService) ListOpenWorkflowExecutions(ctx thrift.Context, listRequest *shared.ListOpenWorkflowExecutionsRequest) (*shared.ListOpenWorkflowExecutionsResponse, error) {
	ret := _m.Called(ctx, listRequest)

	var r0 *shared.ListOpenWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.ListOpenWorkflowExecutionsRequest) *shared.ListOpenWorkflowExecutionsResponse); ok {
		r0 = rf(ctx, listRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ListOpenWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.ListOpenWorkflowExecutionsRequest) error); ok {
		r1 = rf(ctx, listRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PollForActivityTask provides a mock function with given fields: ctx, pollRequest
func (_m *TChanWorkflowService) PollForActivityTask(ctx thrift.Context, pollRequest *shared.PollForActivityTaskRequest) (*shared.PollForActivityTaskResponse, error) {
	ret := _m.Called(ctx, pollRequest)

	var r0 *shared.PollForActivityTaskResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.PollForActivityTaskRequest) *shared.PollForActivityTaskResponse); ok {
		r0 = rf(ctx, pollRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.PollForActivityTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.PollForActivityTaskRequest) error); ok {
		r1 = rf(ctx, pollRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PollForDecisionTask provides a mock function with given fields: ctx, pollRequest
func (_m *TChanWorkflowService) PollForDecisionTask(ctx thrift.Context, pollRequest *shared.PollForDecisionTaskRequest) (*shared.PollForDecisionTaskResponse, error) {
	ret := _m.Called(ctx, pollRequest)

	var r0 *shared.PollForDecisionTaskResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.PollForDecisionTaskRequest) *shared.PollForDecisionTaskResponse); ok {
		r0 = rf(ctx, pollRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.PollForDecisionTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.PollForDecisionTaskRequest) error); ok {
		r1 = rf(ctx, pollRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RecordActivityTaskHeartbeat provides a mock function with given fields: ctx, heartbeatRequest
func (_m *TChanWorkflowService) RecordActivityTaskHeartbeat(ctx thrift.Context, heartbeatRequest *shared.RecordActivityTaskHeartbeatRequest) (*shared.RecordActivityTaskHeartbeatResponse, error) {
	ret := _m.Called(ctx, heartbeatRequest)

	var r0 *shared.RecordActivityTaskHeartbeatResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.RecordActivityTaskHeartbeatRequest) *shared.RecordActivityTaskHeartbeatResponse); ok {
		r0 = rf(ctx, heartbeatRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.RecordActivityTaskHeartbeatResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.RecordActivityTaskHeartbeatRequest) error); ok {
		r1 = rf(ctx, heartbeatRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterDomain provides a mock function with given fields: ctx, registerRequest
func (_m *TChanWorkflowService) RegisterDomain(ctx thrift.Context, registerRequest *shared.RegisterDomainRequest) error {
	ret := _m.Called(ctx, registerRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.RegisterDomainRequest) error); ok {
		r0 = rf(ctx, registerRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RespondActivityTaskCanceled provides a mock function with given fields: ctx, canceledRequest
func (_m *TChanWorkflowService) RespondActivityTaskCanceled(ctx thrift.Context, canceledRequest *shared.RespondActivityTaskCanceledRequest) error {
	ret := _m.Called(ctx, canceledRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.RespondActivityTaskCanceledRequest) error); ok {
		r0 = rf(ctx, canceledRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RespondActivityTaskCompleted provides a mock function with given fields: ctx, completeRequest
func (_m *TChanWorkflowService) RespondActivityTaskCompleted(ctx thrift.Context, completeRequest *shared.RespondActivityTaskCompletedRequest) error {
	ret := _m.Called(ctx, completeRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.RespondActivityTaskCompletedRequest) error); ok {
		r0 = rf(ctx, completeRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RespondActivityTaskFailed provides a mock function with given fields: ctx, failRequest
func (_m *TChanWorkflowService) RespondActivityTaskFailed(ctx thrift.Context, failRequest *shared.RespondActivityTaskFailedRequest) error {
	ret := _m.Called(ctx, failRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.RespondActivityTaskFailedRequest) error); ok {
		r0 = rf(ctx, failRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RespondDecisionTaskCompleted provides a mock function with given fields: ctx, completeRequest
func (_m *TChanWorkflowService) RespondDecisionTaskCompleted(ctx thrift.Context, completeRequest *shared.RespondDecisionTaskCompletedRequest) error {
	ret := _m.Called(ctx, completeRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.RespondDecisionTaskCompletedRequest) error); ok {
		r0 = rf(ctx, completeRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartWorkflowExecution provides a mock function with given fields: ctx, startRequest
func (_m *TChanWorkflowService) StartWorkflowExecution(ctx thrift.Context, startRequest *shared.StartWorkflowExecutionRequest) (*shared.StartWorkflowExecutionResponse, error) {
	ret := _m.Called(ctx, startRequest)

	var r0 *shared.StartWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.StartWorkflowExecutionRequest) *shared.StartWorkflowExecutionResponse); ok {
		r0 = rf(ctx, startRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.StartWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.StartWorkflowExecutionRequest) error); ok {
		r1 = rf(ctx, startRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TerminateWorkflowExecution provides a mock function with given fields: ctx, terminateRequest
func (_m *TChanWorkflowService) TerminateWorkflowExecution(ctx thrift.Context, terminateRequest *shared.TerminateWorkflowExecutionRequest) error {
	ret := _m.Called(ctx, terminateRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.TerminateWorkflowExecutionRequest) error); ok {
		r0 = rf(ctx, terminateRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequestCancelWorkflowExecution provides a mock function with given fields: ctx, cancelRequest
func (_m *TChanWorkflowService) RequestCancelWorkflowExecution(ctx thrift.Context, cancelRequest *shared.RequestCancelWorkflowExecutionRequest) error {
	ret := _m.Called(ctx, cancelRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.RequestCancelWorkflowExecutionRequest) error); ok {
		r0 = rf(ctx, cancelRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignalWorkflowExecution provides a mock function with given fields: ctx, signalRequest
func (_m *TChanWorkflowService) SignalWorkflowExecution(ctx thrift.Context, signalRequest *shared.SignalWorkflowExecutionRequest) error {
	ret := _m.Called(ctx, signalRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.SignalWorkflowExecutionRequest) error); ok {
		r0 = rf(ctx, signalRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDomain provides a mock function with given fields: ctx, updateRequest
func (_m *TChanWorkflowService) UpdateDomain(ctx thrift.Context, updateRequest *shared.UpdateDomainRequest) (*shared.UpdateDomainResponse, error) {
	ret := _m.Called(ctx, updateRequest)

	var r0 *shared.UpdateDomainResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.UpdateDomainRequest) *shared.UpdateDomainResponse); ok {
		r0 = rf(ctx, updateRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.UpdateDomainResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.UpdateDomainRequest) error); ok {
		r1 = rf(ctx, updateRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ cadence.TChanWorkflowService = (*TChanWorkflowService)(nil)