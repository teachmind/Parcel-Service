// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/service/service.go

// Package mock_service is a generated GoMock package.
package mocks

import (
	context "context"
	model "parcel-service/internal/app/model"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockParcelRepository is a mock of ParcelRepository interface.
type MockParcelRepository struct {
	ctrl     *gomock.Controller
	recorder *MockParcelRepositoryMockRecorder
}

// MockParcelRepositoryMockRecorder is the mock recorder for MockParcelRepository.
type MockParcelRepositoryMockRecorder struct {
	mock *MockParcelRepository
}

// NewMockParcelRepository creates a new mock instance.
func NewMockParcelRepository(ctrl *gomock.Controller) *MockParcelRepository {
	mock := &MockParcelRepository{ctrl: ctrl}
	mock.recorder = &MockParcelRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParcelRepository) EXPECT() *MockParcelRepositoryMockRecorder {
	return m.recorder
}

// FetchParcelByID mocks base method.
func (m *MockParcelRepository) FetchParcelByID(ctx context.Context, parcelID int) (model.Parcel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchParcelByID", ctx, parcelID)
	ret0, _ := ret[0].(model.Parcel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchParcelByID indicates an expected call of FetchParcelByID.
func (mr *MockParcelRepositoryMockRecorder) FetchParcelByID(ctx, parcelID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchParcelByID", reflect.TypeOf((*MockParcelRepository)(nil).FetchParcelByID), ctx, parcelID)
}

// GetParcelsList mocks base method.
func (m *MockParcelRepository) GetParcelsList(ctx context.Context, status, limit, offset int) ([]model.Parcel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParcelsList", ctx, status, limit, offset)
	ret0, _ := ret[0].([]model.Parcel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParcelsList indicates an expected call of GetParcelsList.
func (mr *MockParcelRepositoryMockRecorder) GetParcelsList(ctx, status, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParcelsList", reflect.TypeOf((*MockParcelRepository)(nil).GetParcelsList), ctx, status, limit, offset)
}

// InsertParcel mocks base method.
func (m *MockParcelRepository) InsertParcel(ctx context.Context, parcel model.Parcel) (model.Parcel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertParcel", ctx, parcel)
	ret0, _ := ret[0].(model.Parcel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertParcel indicates an expected call of InsertParcel.
func (mr *MockParcelRepositoryMockRecorder) InsertParcel(ctx, parcel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertParcel", reflect.TypeOf((*MockParcelRepository)(nil).InsertParcel), ctx, parcel)
}

// UpdateParcel mocks base method.
func (m *MockParcelRepository) UpdateParcel(ctx context.Context, parcel model.Parcel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateParcel", ctx, parcel)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateParcel indicates an expected call of UpdateParcel.
func (mr *MockParcelRepositoryMockRecorder) UpdateParcel(ctx, parcel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateParcel", reflect.TypeOf((*MockParcelRepository)(nil).UpdateParcel), ctx, parcel)
}

// MockParcelService is a mock of ParcelService interface.
type MockParcelService struct {
	ctrl     *gomock.Controller
	recorder *MockParcelServiceMockRecorder
}

// MockParcelServiceMockRecorder is the mock recorder for MockParcelService.
type MockParcelServiceMockRecorder struct {
	mock *MockParcelService
}

// NewMockParcelService creates a new mock instance.
func NewMockParcelService(ctrl *gomock.Controller) *MockParcelService {
	mock := &MockParcelService{ctrl: ctrl}
	mock.recorder = &MockParcelServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParcelService) EXPECT() *MockParcelServiceMockRecorder {
	return m.recorder
}

// CreateParcel mocks base method.
func (m *MockParcelService) CreateParcel(ctx context.Context, parcel model.Parcel) (model.Parcel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateParcel", ctx, parcel)
	ret0, _ := ret[0].(model.Parcel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateParcel indicates an expected call of CreateParcel.
func (mr *MockParcelServiceMockRecorder) CreateParcel(ctx, parcel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateParcel", reflect.TypeOf((*MockParcelService)(nil).CreateParcel), ctx, parcel)
}

// EditParcel mocks base method.
func (m *MockParcelService) EditParcel(ctx context.Context, parcel model.Parcel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditParcel", ctx, parcel)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditParcel indicates an expected call of EditParcel.
func (mr *MockParcelServiceMockRecorder) EditParcel(ctx, parcel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditParcel", reflect.TypeOf((*MockParcelService)(nil).EditParcel), ctx, parcel)
}

// GetParcelByID mocks base method.
func (m *MockParcelService) GetParcelByID(ctx context.Context, parcelID int) (model.Parcel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParcelByID", ctx, parcelID)
	ret0, _ := ret[0].(model.Parcel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParcelByID indicates an expected call of GetParcelByID.
func (mr *MockParcelServiceMockRecorder) GetParcelByID(ctx, parcelID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParcelByID", reflect.TypeOf((*MockParcelService)(nil).GetParcelByID), ctx, parcelID)
}

// GetParcels mocks base method.
func (m *MockParcelService) GetParcels(ctx context.Context, status, limit, offset int) ([]model.Parcel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParcels", ctx, status, limit, offset)
	ret0, _ := ret[0].([]model.Parcel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParcels indicates an expected call of GetParcels.
func (mr *MockParcelServiceMockRecorder) GetParcels(ctx, status, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParcels", reflect.TypeOf((*MockParcelService)(nil).GetParcels), ctx, status, limit, offset)
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCarrierRepository) EXPECT() *MockCarrierRepositoryMockRecorder {
	return m.recorder
}

// MockCarrierRepository is a mock of CarrierRepository interface.
type MockCarrierRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCarrierRepositoryMockRecorder
}

// MockCarrierRepositoryMockRecorder is the mock recorder for MockCarrierRepository.
type MockCarrierRepositoryMockRecorder struct {
	mock *MockCarrierRepository
}

// NewMockCarrierRepository creates a new mock instance.
func NewMockCarrierRepository(ctrl *gomock.Controller) *MockCarrierRepository {
	mock := &MockCarrierRepository{ctrl: ctrl}
	mock.recorder = &MockCarrierRepositoryMockRecorder{mock}
	return mock
}

// InsertCarrierRequest mocks base method.
func (m *MockCarrierRepository) InsertCarrierRequest(ctx context.Context, carrierReq model.CarrierRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCarrierRequest", ctx, carrierReq)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertCarrierRequest indicates an expected call of InsertCarrierRequest.
func (mr *MockCarrierRepositoryMockRecorder) InsertCarrierRequest(ctx, carrierReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCarrierRequest", reflect.TypeOf((*MockCarrierRepository)(nil).InsertCarrierRequest), ctx, carrierReq)
}

// UpdateCarrierRequest mocks base method.
func (m *MockCarrierRepository) UpdateCarrierRequest(ctx context.Context, parcel model.CarrierRequest, acceptStatus, rejectStatus, parcelStatus int, sourceTime time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCarrierRequest", ctx, parcel, acceptStatus, rejectStatus, parcelStatus, sourceTime)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCarrierRequest indicates an expected call of UpdateCarrierRequest.
func (mr *MockCarrierRepositoryMockRecorder) UpdateCarrierRequest(ctx, parcel, acceptStatus, rejectStatus, parcelStatus, sourceTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCarrierRequest", reflect.TypeOf((*MockCarrierRepository)(nil).UpdateCarrierRequest), ctx, parcel, acceptStatus, rejectStatus, parcelStatus, sourceTime)
}

// MockCarrierService is a mock of CarrierService interface.
type MockCarrierService struct {
	ctrl     *gomock.Controller
	recorder *MockCarrierServiceMockRecorder
}

// MockCarrierServiceMockRecorder is the mock recorder for MockCarrierService.
type MockCarrierServiceMockRecorder struct {
	mock *MockCarrierService
}

// NewMockCarrierService creates a new mock instance.
func NewMockCarrierService(ctrl *gomock.Controller) *MockCarrierService {
	mock := &MockCarrierService{ctrl: ctrl}
	mock.recorder = &MockCarrierServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCarrierService) EXPECT() *MockCarrierServiceMockRecorder {
	return m.recorder
}

// AssignCarrierToParcel mocks base method.
func (m *MockCarrierService) AssignCarrierToParcel(ctx context.Context, parcel model.CarrierRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignCarrierToParcel", ctx, parcel)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignCarrierToParcel indicates an expected call of AssignCarrierToParcel.
func (mr *MockCarrierServiceMockRecorder) AssignCarrierToParcel(ctx, parcel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignCarrierToParcel", reflect.TypeOf((*MockCarrierService)(nil).AssignCarrierToParcel), ctx, parcel)
}

// NewCarrierRequest mocks base method.
func (m *MockCarrierService) NewCarrierRequest(ctx context.Context, carrierReq model.CarrierRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCarrierRequest", ctx, carrierReq)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewCarrierRequest indicates an expected call of NewCarrierRequest.
func (mr *MockCarrierServiceMockRecorder) NewCarrierRequest(ctx, carrierReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCarrierRequest", reflect.TypeOf((*MockCarrierService)(nil).NewCarrierRequest), ctx, carrierReq)
}
