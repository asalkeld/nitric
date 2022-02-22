// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/azure-sdk-for-go/services/eventgrid/2018-01-01/eventgrid/eventgridapi (interfaces: BaseClientAPI)

// Package mock_eventgridapi is a generated GoMock package.
package mock_eventgridapi

import (
	context "context"
	reflect "reflect"

	eventgrid "github.com/Azure/azure-sdk-for-go/services/eventgrid/2018-01-01/eventgrid"
	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
)

// MockBaseClientAPI is a mock of BaseClientAPI interface.
type MockBaseClientAPI struct {
	ctrl     *gomock.Controller
	recorder *MockBaseClientAPIMockRecorder
}

// MockBaseClientAPIMockRecorder is the mock recorder for MockBaseClientAPI.
type MockBaseClientAPIMockRecorder struct {
	mock *MockBaseClientAPI
}

// NewMockBaseClientAPI creates a new mock instance.
func NewMockBaseClientAPI(ctrl *gomock.Controller) *MockBaseClientAPI {
	mock := &MockBaseClientAPI{ctrl: ctrl}
	mock.recorder = &MockBaseClientAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBaseClientAPI) EXPECT() *MockBaseClientAPIMockRecorder {
	return m.recorder
}

// PublishCloudEventEvents mocks base method.
func (m *MockBaseClientAPI) PublishCloudEventEvents(arg0 context.Context, arg1 string, arg2 []eventgrid.CloudEventEvent) (autorest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishCloudEventEvents", arg0, arg1, arg2)
	ret0, _ := ret[0].(autorest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishCloudEventEvents indicates an expected call of PublishCloudEventEvents.
func (mr *MockBaseClientAPIMockRecorder) PublishCloudEventEvents(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishCloudEventEvents", reflect.TypeOf((*MockBaseClientAPI)(nil).PublishCloudEventEvents), arg0, arg1, arg2)
}

// PublishCustomEventEvents mocks base method.
func (m *MockBaseClientAPI) PublishCustomEventEvents(arg0 context.Context, arg1 string, arg2 []interface{}) (autorest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishCustomEventEvents", arg0, arg1, arg2)
	ret0, _ := ret[0].(autorest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishCustomEventEvents indicates an expected call of PublishCustomEventEvents.
func (mr *MockBaseClientAPIMockRecorder) PublishCustomEventEvents(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishCustomEventEvents", reflect.TypeOf((*MockBaseClientAPI)(nil).PublishCustomEventEvents), arg0, arg1, arg2)
}

// PublishEvents mocks base method.
func (m *MockBaseClientAPI) PublishEvents(arg0 context.Context, arg1 string, arg2 []eventgrid.Event) (autorest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishEvents", arg0, arg1, arg2)
	ret0, _ := ret[0].(autorest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishEvents indicates an expected call of PublishEvents.
func (mr *MockBaseClientAPIMockRecorder) PublishEvents(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishEvents", reflect.TypeOf((*MockBaseClientAPI)(nil).PublishEvents), arg0, arg1, arg2)
}
