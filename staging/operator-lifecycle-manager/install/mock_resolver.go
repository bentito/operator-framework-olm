// Code generated by MockGen. DO NOT EDIT.
// Source: install/resolver.go

// Package install is a generated GoMock package.
package install

import (
	v1alpha1 "github.com/coreos-inc/alm/apis/clusterserviceversion/v1alpha1"
	client "github.com/coreos-inc/operator-client/pkg/client"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	reflect "reflect"
)

// MockResolver is a mock of Resolver interface
type MockResolver struct {
	ctrl     *gomock.Controller
	recorder *MockResolverMockRecorder
}

// MockResolverMockRecorder is the mock recorder for MockResolver
type MockResolverMockRecorder struct {
	mock *MockResolver
}

// NewMockResolver creates a new mock instance
func NewMockResolver(ctrl *gomock.Controller) *MockResolver {
	mock := &MockResolver{ctrl: ctrl}
	mock.recorder = &MockResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResolver) EXPECT() *MockResolverMockRecorder {
	return m.recorder
}

// CheckInstalled mocks base method
func (m *MockResolver) CheckInstalled(s v1alpha1.NamedInstallStrategy, owner v1.ObjectMeta, ownerType v1.TypeMeta) (bool, error) {
	ret := m.ctrl.Call(m, "CheckInstalled", s, owner, ownerType)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckInstalled indicates an expected call of CheckInstalled
func (mr *MockResolverMockRecorder) CheckInstalled(s, owner, ownerType interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckInstalled", reflect.TypeOf((*MockResolver)(nil).CheckInstalled), s, owner, ownerType)
}

// ApplyStrategy mocks base method
func (m *MockResolver) ApplyStrategy(s v1alpha1.NamedInstallStrategy, owner v1.ObjectMeta, ownerType v1.TypeMeta) error {
	ret := m.ctrl.Call(m, "ApplyStrategy", s, owner, ownerType)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyStrategy indicates an expected call of ApplyStrategy
func (mr *MockResolverMockRecorder) ApplyStrategy(s, owner, ownerType interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyStrategy", reflect.TypeOf((*MockResolver)(nil).ApplyStrategy), s, owner, ownerType)
}

// UnmarshalStrategy mocks base method
func (m *MockResolver) UnmarshalStrategy(s v1alpha1.NamedInstallStrategy) (Strategy, error) {
	ret := m.ctrl.Call(m, "UnmarshalStrategy", s)
	ret0, _ := ret[0].(Strategy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnmarshalStrategy indicates an expected call of UnmarshalStrategy
func (mr *MockResolverMockRecorder) UnmarshalStrategy(s interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalStrategy", reflect.TypeOf((*MockResolver)(nil).UnmarshalStrategy), s)
}

// MockStrategy is a mock of Strategy interface
type MockStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockStrategyMockRecorder
}

// MockStrategyMockRecorder is the mock recorder for MockStrategy
type MockStrategyMockRecorder struct {
	mock *MockStrategy
}

// NewMockStrategy creates a new mock instance
func NewMockStrategy(ctrl *gomock.Controller) *MockStrategy {
	mock := &MockStrategy{ctrl: ctrl}
	mock.recorder = &MockStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStrategy) EXPECT() *MockStrategyMockRecorder {
	return m.recorder
}

// Install mocks base method
func (m *MockStrategy) Install(client client.Interface, owner v1.ObjectMeta, ownerType v1.TypeMeta) error {
	ret := m.ctrl.Call(m, "Install", client, owner, ownerType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install
func (mr *MockStrategyMockRecorder) Install(client, owner, ownerType interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockStrategy)(nil).Install), client, owner, ownerType)
}

// CheckInstalled mocks base method
func (m *MockStrategy) CheckInstalled(client client.Interface, owner v1.ObjectMeta) (bool, error) {
	ret := m.ctrl.Call(m, "CheckInstalled", client, owner)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckInstalled indicates an expected call of CheckInstalled
func (mr *MockStrategyMockRecorder) CheckInstalled(client, owner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckInstalled", reflect.TypeOf((*MockStrategy)(nil).CheckInstalled), client, owner)
}
