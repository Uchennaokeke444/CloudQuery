// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/authentication/v1beta1 (interfaces: AuthenticationV1beta1Interface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	rest "k8s.io/client-go/rest"
)

// MockAuthenticationV1beta1Interface is a mock of AuthenticationV1beta1Interface interface.
type MockAuthenticationV1beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationV1beta1InterfaceMockRecorder
}

// MockAuthenticationV1beta1InterfaceMockRecorder is the mock recorder for MockAuthenticationV1beta1Interface.
type MockAuthenticationV1beta1InterfaceMockRecorder struct {
	mock *MockAuthenticationV1beta1Interface
}

// NewMockAuthenticationV1beta1Interface creates a new mock instance.
func NewMockAuthenticationV1beta1Interface(ctrl *gomock.Controller) *MockAuthenticationV1beta1Interface {
	mock := &MockAuthenticationV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockAuthenticationV1beta1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthenticationV1beta1Interface) EXPECT() *MockAuthenticationV1beta1InterfaceMockRecorder {
	return m.recorder
}

// RESTClient mocks base method.
func (m *MockAuthenticationV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockAuthenticationV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockAuthenticationV1beta1Interface)(nil).RESTClient))
}

// SelfSubjectReviews mocks base method.
func (m *MockAuthenticationV1beta1Interface) SelfSubjectReviews() v1beta1.SelfSubjectReviewInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelfSubjectReviews")
	ret0, _ := ret[0].(v1beta1.SelfSubjectReviewInterface)
	return ret0
}

// SelfSubjectReviews indicates an expected call of SelfSubjectReviews.
func (mr *MockAuthenticationV1beta1InterfaceMockRecorder) SelfSubjectReviews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelfSubjectReviews", reflect.TypeOf((*MockAuthenticationV1beta1Interface)(nil).SelfSubjectReviews))
}

// TokenReviews mocks base method.
func (m *MockAuthenticationV1beta1Interface) TokenReviews() v1beta1.TokenReviewInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenReviews")
	ret0, _ := ret[0].(v1beta1.TokenReviewInterface)
	return ret0
}

// TokenReviews indicates an expected call of TokenReviews.
func (mr *MockAuthenticationV1beta1InterfaceMockRecorder) TokenReviews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenReviews", reflect.TypeOf((*MockAuthenticationV1beta1Interface)(nil).TokenReviews))
}
