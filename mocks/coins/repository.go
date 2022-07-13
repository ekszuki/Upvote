// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/coins/repository.go

// Package coins is a generated GoMock package.
package coins

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "klever.io/interview/pkg/domain"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockRepository) Delete(ctx context.Context, coinID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, coinID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(ctx, coinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, coinID)
}

// FindByID mocks base method.
func (m *MockRepository) FindByID(ctx context.Context, coinID uint) (*domain.Coin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, coinID)
	ret0, _ := ret[0].(*domain.Coin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockRepositoryMockRecorder) FindByID(ctx, coinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockRepository)(nil).FindByID), ctx, coinID)
}

// ListActive mocks base method.
func (m *MockRepository) ListActive(ctx context.Context) ([]domain.Coin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListActive", ctx)
	ret0, _ := ret[0].([]domain.Coin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListActive indicates an expected call of ListActive.
func (mr *MockRepositoryMockRecorder) ListActive(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListActive", reflect.TypeOf((*MockRepository)(nil).ListActive), ctx)
}

// Save mocks base method.
func (m *MockRepository) Save(ctx context.Context, coin *domain.Coin) (*domain.Coin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, coin)
	ret0, _ := ret[0].(*domain.Coin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockRepositoryMockRecorder) Save(ctx, coin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRepository)(nil).Save), ctx, coin)
}

// Update mocks base method.
func (m *MockRepository) Update(ctx context.Context, coin *domain.Coin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, coin)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(ctx, coin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), ctx, coin)
}

// VoteDown mocks base method.
func (m *MockRepository) VoteDown(ctx context.Context, coinID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VoteDown", ctx, coinID)
	ret0, _ := ret[0].(error)
	return ret0
}

// VoteDown indicates an expected call of VoteDown.
func (mr *MockRepositoryMockRecorder) VoteDown(ctx, coinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VoteDown", reflect.TypeOf((*MockRepository)(nil).VoteDown), ctx, coinID)
}

// VoteUP mocks base method.
func (m *MockRepository) VoteUP(ctx context.Context, coinID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VoteUP", ctx, coinID)
	ret0, _ := ret[0].(error)
	return ret0
}

// VoteUP indicates an expected call of VoteUP.
func (mr *MockRepositoryMockRecorder) VoteUP(ctx, coinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VoteUP", reflect.TypeOf((*MockRepository)(nil).VoteUP), ctx, coinID)
}