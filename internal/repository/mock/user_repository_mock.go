// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	model "qd-authentication-api/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepositoryer is a mock of UserRepositoryer interface.
type MockUserRepositoryer struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryerMockRecorder
}

// MockUserRepositoryerMockRecorder is the mock recorder for MockUserRepositoryer.
type MockUserRepositoryerMockRecorder struct {
	mock *MockUserRepositoryer
}

// NewMockUserRepositoryer creates a new mock instance.
func NewMockUserRepositoryer(ctrl *gomock.Controller) *MockUserRepositoryer {
	mock := &MockUserRepositoryer{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepositoryer) EXPECT() *MockUserRepositoryerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepositoryer) Create(ctx context.Context, user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryerMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepositoryer)(nil).Create), ctx, user)
}

// GetByEmail mocks base method.
func (m *MockUserRepositoryer) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUserRepositoryerMockRecorder) GetByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUserRepositoryer)(nil).GetByEmail), ctx, email)
}

// GetByVerificationToken mocks base method.
func (m *MockUserRepositoryer) GetByVerificationToken(ctx context.Context, token string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByVerificationToken", ctx, token)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByVerificationToken indicates an expected call of GetByVerificationToken.
func (mr *MockUserRepositoryerMockRecorder) GetByVerificationToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByVerificationToken", reflect.TypeOf((*MockUserRepositoryer)(nil).GetByVerificationToken), ctx, token)
}

// Update mocks base method.
func (m *MockUserRepositoryer) Update(ctx context.Context, user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryerMockRecorder) Update(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepositoryer)(nil).Update), ctx, user)
}
