// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory/fosite/handler/openid (interfaces: OpenIDConnectRequestStorage)

package internal

import (
	context "context"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
)

// Mock of OpenIDConnectRequestStorage interface
type MockOpenIDConnectRequestStorage struct {
	ctrl     *gomock.Controller
	recorder *_MockOpenIDConnectRequestStorageRecorder
}

// Recorder for MockOpenIDConnectRequestStorage (not exported)
type _MockOpenIDConnectRequestStorageRecorder struct {
	mock *MockOpenIDConnectRequestStorage
}

func NewMockOpenIDConnectRequestStorage(ctrl *gomock.Controller) *MockOpenIDConnectRequestStorage {
	mock := &MockOpenIDConnectRequestStorage{ctrl: ctrl}
	mock.recorder = &_MockOpenIDConnectRequestStorageRecorder{mock}
	return mock
}

func (_m *MockOpenIDConnectRequestStorage) EXPECT() *_MockOpenIDConnectRequestStorageRecorder {
	return _m.recorder
}

func (_m *MockOpenIDConnectRequestStorage) CreateOpenIDConnectSession(_param0 context.Context, _param1 string, _param2 fosite.Requester) error {
	ret := _m.ctrl.Call(_m, "CreateOpenIDConnectSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockOpenIDConnectRequestStorageRecorder) CreateOpenIDConnectSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateOpenIDConnectSession", arg0, arg1, arg2)
}

func (_m *MockOpenIDConnectRequestStorage) DeleteOpenIDConnectSession(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteOpenIDConnectSession", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockOpenIDConnectRequestStorageRecorder) DeleteOpenIDConnectSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteOpenIDConnectSession", arg0, arg1)
}

func (_m *MockOpenIDConnectRequestStorage) GetOpenIDConnectSession(_param0 context.Context, _param1 string, _param2 fosite.Requester) (fosite.Requester, error) {
	ret := _m.ctrl.Call(_m, "GetOpenIDConnectSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockOpenIDConnectRequestStorageRecorder) GetOpenIDConnectSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOpenIDConnectSession", arg0, arg1, arg2)
}
