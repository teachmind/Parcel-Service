// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	model "parcel-service/internal/app/model"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockCarrierParcelAcceptRepository is a mock of CarrierAcceptRepository interface.
type MockCarrierParcelAcceptRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCarrierParcelAcceptRepositoryMockRecorder
}

// MockCarrierParcelAcceptRepositoryMockRecorder is the mock recorder for MockCarrierParcelAcceptRepository.
type MockCarrierParcelAcceptRepositoryMockRecorder struct {
	mock *MockCarrierParcelAcceptRepository
}

// NewMockCarrierParcelAcceptRepository creates a new mock instance.
func NewMockCarrierParcelAcceptRepository(ctrl *gomock.Controller) *MockCarrierParcelAcceptRepository {
	mock := &MockCarrierParcelAcceptRepository{ctrl: ctrl}
	mock.recorder = &MockCarrierParcelAcceptRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCarrierParcelAcceptRepository) EXPECT() *MockCarrierParcelAcceptRepositoryMockRecorder {
	return m.recorder
}

// UpdateCarrierRequest mocks base method.
func (m *MockCarrierParcelAcceptRepository) UpdateCarrierRequest(ctx context.Context, parcel model.CarrierRequest, acceptStatus, rejectStatus, parcelStatus int, sourceTime time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCarrierRequest", ctx, parcel, acceptStatus, rejectStatus, parcelStatus, sourceTime)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCarrierRequest indicates an expected call of UpdateCarrierRequest.
func (mr *MockCarrierParcelAcceptRepositoryMockRecorder) UpdateCarrierRequest(ctx, parcel, acceptStatus, rejectStatus, parcelStatus, sourceTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCarrierRequest", reflect.TypeOf((*MockCarrierParcelAcceptRepository)(nil).UpdateCarrierRequest), ctx, parcel, acceptStatus, rejectStatus, parcelStatus, sourceTime)
}

// MockCarrierParcelAcceptService is a mock of CarrierParcelAcceptService interface.
type MockCarrierParcelAcceptService struct {
	ctrl     *gomock.Controller
	recorder *MockCarrierParcelAcceptServiceMockRecorder
}

// MockCarrierParcelAcceptServiceMockRecorder is the mock recorder for MockCarrierParcelAcceptService.
type MockCarrierParcelAcceptServiceMockRecorder struct {
	mock *MockCarrierParcelAcceptService
}

// NewMockCarrierParcelAcceptService creates a new mock instance.
func NewMockCarrierParcelAcceptService(ctrl *gomock.Controller) *MockCarrierParcelAcceptService {
	mock := &MockCarrierParcelAcceptService{ctrl: ctrl}
	mock.recorder = &MockCarrierParcelAcceptServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCarrierParcelAcceptService) EXPECT() *MockCarrierParcelAcceptServiceMockRecorder {
	return m.recorder
}

// AssignCarrierToParcel mocks base method.
func (m *MockCarrierParcelAcceptService) AssignCarrierToParcel(ctx context.Context, parcel model.CarrierRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignCarrierToParcel", ctx, parcel)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignCarrierToParcel indicates an expected call of AssignCarrierToParcel.
func (mr *MockCarrierParcelAcceptServiceMockRecorder) AssignCarrierToParcel(ctx, parcel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignCarrierToParcel", reflect.TypeOf((*MockCarrierParcelAcceptService)(nil).AssignCarrierToParcel), ctx, parcel)
}
