// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zeebe-io/zeebe/clients/go/pkg/pb (interfaces: GatewayClient,Gateway_ActivateJobsClient)

package mock_pb

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	pb "github.com/zeebe-io/zeebe/clients/go/pkg/pb"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockGatewayClient is a mock of GatewayClient interface
type MockGatewayClient struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayClientMockRecorder
}

// MockGatewayClientMockRecorder is the mock recorder for MockGatewayClient
type MockGatewayClientMockRecorder struct {
	mock *MockGatewayClient
}

// NewMockGatewayClient creates a new mock instance
func NewMockGatewayClient(ctrl *gomock.Controller) *MockGatewayClient {
	mock := &MockGatewayClient{ctrl: ctrl}
	mock.recorder = &MockGatewayClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockGatewayClient) EXPECT() *MockGatewayClientMockRecorder {
	return _m.recorder
}

// ActivateJobs mocks base method
func (_m *MockGatewayClient) ActivateJobs(_param0 context.Context, _param1 *pb.ActivateJobsRequest, _param2 ...grpc.CallOption) (pb.Gateway_ActivateJobsClient, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ActivateJobs", _s...)
	ret0, _ := ret[0].(pb.Gateway_ActivateJobsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActivateJobs indicates an expected call of ActivateJobs
func (_mr *MockGatewayClientMockRecorder) ActivateJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ActivateJobs", reflect.TypeOf((*MockGatewayClient)(nil).ActivateJobs), _s...)
}

// CancelWorkflowInstance mocks base method
func (_m *MockGatewayClient) CancelWorkflowInstance(_param0 context.Context, _param1 *pb.CancelWorkflowInstanceRequest, _param2 ...grpc.CallOption) (*pb.CancelWorkflowInstanceResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CancelWorkflowInstance", _s...)
	ret0, _ := ret[0].(*pb.CancelWorkflowInstanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelWorkflowInstance indicates an expected call of CancelWorkflowInstance
func (_mr *MockGatewayClientMockRecorder) CancelWorkflowInstance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CancelWorkflowInstance", reflect.TypeOf((*MockGatewayClient)(nil).CancelWorkflowInstance), _s...)
}

// CompleteJob mocks base method
func (_m *MockGatewayClient) CompleteJob(_param0 context.Context, _param1 *pb.CompleteJobRequest, _param2 ...grpc.CallOption) (*pb.CompleteJobResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CompleteJob", _s...)
	ret0, _ := ret[0].(*pb.CompleteJobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompleteJob indicates an expected call of CompleteJob
func (_mr *MockGatewayClientMockRecorder) CompleteJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CompleteJob", reflect.TypeOf((*MockGatewayClient)(nil).CompleteJob), _s...)
}

// CreateWorkflowInstance mocks base method
func (_m *MockGatewayClient) CreateWorkflowInstance(_param0 context.Context, _param1 *pb.CreateWorkflowInstanceRequest, _param2 ...grpc.CallOption) (*pb.CreateWorkflowInstanceResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CreateWorkflowInstance", _s...)
	ret0, _ := ret[0].(*pb.CreateWorkflowInstanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWorkflowInstance indicates an expected call of CreateWorkflowInstance
func (_mr *MockGatewayClientMockRecorder) CreateWorkflowInstance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateWorkflowInstance", reflect.TypeOf((*MockGatewayClient)(nil).CreateWorkflowInstance), _s...)
}

// CreateWorkflowInstanceWithResult mocks base method
func (_m *MockGatewayClient) CreateWorkflowInstanceWithResult(_param0 context.Context, _param1 *pb.CreateWorkflowInstanceWithResultRequest, _param2 ...grpc.CallOption) (*pb.CreateWorkflowInstanceWithResultResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CreateWorkflowInstanceWithResult", _s...)
	ret0, _ := ret[0].(*pb.CreateWorkflowInstanceWithResultResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWorkflowInstanceWithResult indicates an expected call of CreateWorkflowInstanceWithResult
func (_mr *MockGatewayClientMockRecorder) CreateWorkflowInstanceWithResult(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateWorkflowInstanceWithResult", reflect.TypeOf((*MockGatewayClient)(nil).CreateWorkflowInstanceWithResult), _s...)
}

// DeployWorkflow mocks base method
func (_m *MockGatewayClient) DeployWorkflow(_param0 context.Context, _param1 *pb.DeployWorkflowRequest, _param2 ...grpc.CallOption) (*pb.DeployWorkflowResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DeployWorkflow", _s...)
	ret0, _ := ret[0].(*pb.DeployWorkflowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeployWorkflow indicates an expected call of DeployWorkflow
func (_mr *MockGatewayClientMockRecorder) DeployWorkflow(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeployWorkflow", reflect.TypeOf((*MockGatewayClient)(nil).DeployWorkflow), _s...)
}

// FailJob mocks base method
func (_m *MockGatewayClient) FailJob(_param0 context.Context, _param1 *pb.FailJobRequest, _param2 ...grpc.CallOption) (*pb.FailJobResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "FailJob", _s...)
	ret0, _ := ret[0].(*pb.FailJobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FailJob indicates an expected call of FailJob
func (_mr *MockGatewayClientMockRecorder) FailJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FailJob", reflect.TypeOf((*MockGatewayClient)(nil).FailJob), _s...)
}

// PublishMessage mocks base method
func (_m *MockGatewayClient) PublishMessage(_param0 context.Context, _param1 *pb.PublishMessageRequest, _param2 ...grpc.CallOption) (*pb.PublishMessageResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "PublishMessage", _s...)
	ret0, _ := ret[0].(*pb.PublishMessageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishMessage indicates an expected call of PublishMessage
func (_mr *MockGatewayClientMockRecorder) PublishMessage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "PublishMessage", reflect.TypeOf((*MockGatewayClient)(nil).PublishMessage), _s...)
}

// ResolveIncident mocks base method
func (_m *MockGatewayClient) ResolveIncident(_param0 context.Context, _param1 *pb.ResolveIncidentRequest, _param2 ...grpc.CallOption) (*pb.ResolveIncidentResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ResolveIncident", _s...)
	ret0, _ := ret[0].(*pb.ResolveIncidentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveIncident indicates an expected call of ResolveIncident
func (_mr *MockGatewayClientMockRecorder) ResolveIncident(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ResolveIncident", reflect.TypeOf((*MockGatewayClient)(nil).ResolveIncident), _s...)
}

// SetVariables mocks base method
func (_m *MockGatewayClient) SetVariables(_param0 context.Context, _param1 *pb.SetVariablesRequest, _param2 ...grpc.CallOption) (*pb.SetVariablesResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "SetVariables", _s...)
	ret0, _ := ret[0].(*pb.SetVariablesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetVariables indicates an expected call of SetVariables
func (_mr *MockGatewayClientMockRecorder) SetVariables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetVariables", reflect.TypeOf((*MockGatewayClient)(nil).SetVariables), _s...)
}

// ThrowError mocks base method
func (_m *MockGatewayClient) ThrowError(_param0 context.Context, _param1 *pb.ThrowErrorRequest, _param2 ...grpc.CallOption) (*pb.ThrowErrorResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ThrowError", _s...)
	ret0, _ := ret[0].(*pb.ThrowErrorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ThrowError indicates an expected call of ThrowError
func (_mr *MockGatewayClientMockRecorder) ThrowError(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ThrowError", reflect.TypeOf((*MockGatewayClient)(nil).ThrowError), _s...)
}

// Topology mocks base method
func (_m *MockGatewayClient) Topology(_param0 context.Context, _param1 *pb.TopologyRequest, _param2 ...grpc.CallOption) (*pb.TopologyResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Topology", _s...)
	ret0, _ := ret[0].(*pb.TopologyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Topology indicates an expected call of Topology
func (_mr *MockGatewayClientMockRecorder) Topology(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Topology", reflect.TypeOf((*MockGatewayClient)(nil).Topology), _s...)
}

// UpdateJobRetries mocks base method
func (_m *MockGatewayClient) UpdateJobRetries(_param0 context.Context, _param1 *pb.UpdateJobRetriesRequest, _param2 ...grpc.CallOption) (*pb.UpdateJobRetriesResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "UpdateJobRetries", _s...)
	ret0, _ := ret[0].(*pb.UpdateJobRetriesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateJobRetries indicates an expected call of UpdateJobRetries
func (_mr *MockGatewayClientMockRecorder) UpdateJobRetries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UpdateJobRetries", reflect.TypeOf((*MockGatewayClient)(nil).UpdateJobRetries), _s...)
}

// MockGateway_ActivateJobsClient is a mock of Gateway_ActivateJobsClient interface
type MockGateway_ActivateJobsClient struct {
	ctrl     *gomock.Controller
	recorder *MockGateway_ActivateJobsClientMockRecorder
}

// MockGateway_ActivateJobsClientMockRecorder is the mock recorder for MockGateway_ActivateJobsClient
type MockGateway_ActivateJobsClientMockRecorder struct {
	mock *MockGateway_ActivateJobsClient
}

// NewMockGateway_ActivateJobsClient creates a new mock instance
func NewMockGateway_ActivateJobsClient(ctrl *gomock.Controller) *MockGateway_ActivateJobsClient {
	mock := &MockGateway_ActivateJobsClient{ctrl: ctrl}
	mock.recorder = &MockGateway_ActivateJobsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockGateway_ActivateJobsClient) EXPECT() *MockGateway_ActivateJobsClientMockRecorder {
	return _m.recorder
}

// CloseSend mocks base method
func (_m *MockGateway_ActivateJobsClient) CloseSend() error {
	ret := _m.ctrl.Call(_m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (_mr *MockGateway_ActivateJobsClientMockRecorder) CloseSend() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CloseSend", reflect.TypeOf((*MockGateway_ActivateJobsClient)(nil).CloseSend))
}

// Context mocks base method
func (_m *MockGateway_ActivateJobsClient) Context() context.Context {
	ret := _m.ctrl.Call(_m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (_mr *MockGateway_ActivateJobsClientMockRecorder) Context() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Context", reflect.TypeOf((*MockGateway_ActivateJobsClient)(nil).Context))
}

// Header mocks base method
func (_m *MockGateway_ActivateJobsClient) Header() (metadata.MD, error) {
	ret := _m.ctrl.Call(_m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (_mr *MockGateway_ActivateJobsClientMockRecorder) Header() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Header", reflect.TypeOf((*MockGateway_ActivateJobsClient)(nil).Header))
}

// Recv mocks base method
func (_m *MockGateway_ActivateJobsClient) Recv() (*pb.ActivateJobsResponse, error) {
	ret := _m.ctrl.Call(_m, "Recv")
	ret0, _ := ret[0].(*pb.ActivateJobsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (_mr *MockGateway_ActivateJobsClientMockRecorder) Recv() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Recv", reflect.TypeOf((*MockGateway_ActivateJobsClient)(nil).Recv))
}

// RecvMsg mocks base method
func (_m *MockGateway_ActivateJobsClient) RecvMsg(_param0 interface{}) error {
	ret := _m.ctrl.Call(_m, "RecvMsg", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (_mr *MockGateway_ActivateJobsClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RecvMsg", reflect.TypeOf((*MockGateway_ActivateJobsClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (_m *MockGateway_ActivateJobsClient) SendMsg(_param0 interface{}) error {
	ret := _m.ctrl.Call(_m, "SendMsg", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (_mr *MockGateway_ActivateJobsClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SendMsg", reflect.TypeOf((*MockGateway_ActivateJobsClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (_m *MockGateway_ActivateJobsClient) Trailer() metadata.MD {
	ret := _m.ctrl.Call(_m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (_mr *MockGateway_ActivateJobsClientMockRecorder) Trailer() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Trailer", reflect.TypeOf((*MockGateway_ActivateJobsClient)(nil).Trailer))
}
