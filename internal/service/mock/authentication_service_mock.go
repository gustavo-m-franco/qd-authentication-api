// Code generated by MockGen. DO NOT EDIT.
// Source: ./authentication_service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	model "qd-authentication-api/internal/model"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthenticationServicer is a mock of AuthenticationServicer interface.
type MockAuthenticationServicer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationServicerMockRecorder
}

// MockAuthenticationServicerMockRecorder is the mock recorder for MockAuthenticationServicer.
type MockAuthenticationServicerMockRecorder struct {
	mock *MockAuthenticationServicer
}

// NewMockAuthenticationServicer creates a new mock instance.
func NewMockAuthenticationServicer(ctrl *gomock.Controller) *MockAuthenticationServicer {
	mock := &MockAuthenticationServicer{ctrl: ctrl}
	mock.recorder = &MockAuthenticationServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthenticationServicer) EXPECT() *MockAuthenticationServicerMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockAuthenticationServicer) Authenticate(ctx context.Context, email, password string) (*model.AuthTokensResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx, email, password)
	ret0, _ := ret[0].(*model.AuthTokensResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockAuthenticationServicerMockRecorder) Authenticate(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthenticationServicer)(nil).Authenticate), ctx, email, password)
}

// GetPublicKey mocks base method.
func (m *MockAuthenticationServicer) GetPublicKey(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKey", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicKey indicates an expected call of GetPublicKey.
func (mr *MockAuthenticationServicerMockRecorder) GetPublicKey(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKey", reflect.TypeOf((*MockAuthenticationServicer)(nil).GetPublicKey), ctx)
}

// Register mocks base method.
func (m *MockAuthenticationServicer) Register(ctx context.Context, email, password, firstName, lastName string, dateOfBirth *time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, email, password, firstName, lastName, dateOfBirth)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockAuthenticationServicerMockRecorder) Register(ctx, email, password, firstName, lastName, dateOfBirth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthenticationServicer)(nil).Register), ctx, email, password, firstName, lastName, dateOfBirth)
}

// ResendEmailVerification mocks base method.
func (m *MockAuthenticationServicer) ResendEmailVerification(ctx context.Context, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResendEmailVerification", ctx, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResendEmailVerification indicates an expected call of ResendEmailVerification.
func (mr *MockAuthenticationServicerMockRecorder) ResendEmailVerification(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResendEmailVerification", reflect.TypeOf((*MockAuthenticationServicer)(nil).ResendEmailVerification), ctx, email)
}

// VerifyEmail mocks base method.
func (m *MockAuthenticationServicer) VerifyEmail(ctx context.Context, verificationToken string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyEmail", ctx, verificationToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyEmail indicates an expected call of VerifyEmail.
func (mr *MockAuthenticationServicerMockRecorder) VerifyEmail(ctx, verificationToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyEmail", reflect.TypeOf((*MockAuthenticationServicer)(nil).VerifyEmail), ctx, verificationToken)
}

// VerifyTokenAndDecodeEmail mocks base method.
func (m *MockAuthenticationServicer) VerifyTokenAndDecodeEmail(ctx context.Context, token string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyTokenAndDecodeEmail", ctx, token)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyTokenAndDecodeEmail indicates an expected call of VerifyTokenAndDecodeEmail.
func (mr *MockAuthenticationServicerMockRecorder) VerifyTokenAndDecodeEmail(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyTokenAndDecodeEmail", reflect.TypeOf((*MockAuthenticationServicer)(nil).VerifyTokenAndDecodeEmail), ctx, token)
}
