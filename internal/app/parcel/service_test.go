package parcel

import (
	"context"
	"errors"
	"parcel-service/internal/app/model"
	"parcel-service/internal/app/service/mocks"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var parcel = model.Parcel{
	UserID:             1,
	SourceAddress:      "Dhaka Bangladesh",
	DestinationAddress: "Pabna Shadar",
	SourceTime:         time.Now(),
	ParcelType:         "Document",
	Status:             2,
	Price:              200.0,
	CarrierFee:         180.0,
	CompanyFee:         20.0,
}

func TestService_GetParcels(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	parcels := []model.Parcel{
		{
			ID:                 1,
			UserID:             1,
			CarrierID:          0,
			Status:             0,
			SourceAddress:      "Dhaka Bangladesh",
			DestinationAddress: "Pabna Shadar",
			ParcelType:         "Document",
			Price:              200,
			CarrierFee:         180,
			CompanyFee:         20,
		}, {
			ID:                 1,
			UserID:             1,
			SourceAddress:      "Dhaka Bangladesh",
			DestinationAddress: "Pabna Shadar",
			ParcelType:         "Document",
			Price:              200,
			CarrierFee:         180,
			CompanyFee:         20,
		}}

	testCases := []struct {
		desc      string
		status    int
		offset    int
		limit     int
		mockRepo  func() *mocks.MockParcelRepository
		expErr    error
		expParcel []model.Parcel
	}{
		{
			desc:   "should return success",
			status: 1,
			offset: 2,
			limit:  4,
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().GetParcelsList(gomock.Any(), 1, 2, 4).Return(parcels, nil)
				return r
			},
			expErr:    nil,
			expParcel: parcels,
		},

		{
			desc:   "Should return not found",
			status: 1,
			offset: 2,
			limit:  4,
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().GetParcelsList(gomock.Any(), 1, 2, 4).Return([]model.Parcel{}, model.ErrNotFound)
				return r
			},
			expErr:    model.ErrNotFound,
			expParcel: []model.Parcel(nil),
		},
		{
			desc:   "should return DB error",
			status: 1,
			offset: 2,
			limit:  4,
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().GetParcelsList(gomock.Any(), 1, 2, 4).Return([]model.Parcel{}, errors.New("db-error"))
				return r
			},
			expErr:    errors.New("db-error"),
			expParcel: []model.Parcel(nil),
		},
		{
			desc:   "should return empty parcel",
			status: 1,
			offset: 2,
			limit:  4,
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().GetParcelsList(gomock.Any(), 1, 2, 4).Return([]model.Parcel{}, nil)
				return r
			},
			expErr:    nil,
			expParcel: []model.Parcel{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewService(tc.mockRepo())
			parcels, err := s.GetParcels(context.Background(), 1, 2, 4)
			assert.Equal(t, tc.expErr, err)
			assert.EqualValues(t, tc.expParcel, parcels)
		})
	}
}

func TestService_CreateParcel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc     string
		mockRepo func() *mocks.MockParcelRepository
		expErr   error
	}{
		{
			desc: "should return success",
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().InsertParcel(gomock.Any(), gomock.Any()).Return(nil)
				return r
			},
			expErr: nil,
		},

		{
			desc: "should return db error",
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().InsertParcel(gomock.Any(), gomock.Any()).Return(errors.New("db-error"))
				return r
			},
			expErr: errors.New("db-error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewService(tc.mockRepo())
			err := s.CreateParcel(context.Background(), parcel)
			assert.Equal(t, tc.expErr, err)
		})
	}
}

func TestService_GetParcelByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc      string
		parcelID  int
		mockRepo  func() *mocks.MockParcelRepository
		expErr    error
		expParcel model.Parcel
	}{
		{
			desc:     "should return success",
			parcelID: 1,
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().FetchParcelByID(gomock.Any(), parcel.ID).Return(parcel, nil)
				return r
			},
			expErr:    nil,
			expParcel: parcel,
		},
		{
			desc:     "Should return not found",
			parcelID: 1,
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().FetchParcelByID(gomock.Any(), parcel.ID).Return(model.Parcel{}, model.ErrNotFound)
				return r
			},
			expErr:    model.ErrNotFound,
			expParcel: model.Parcel{},
		},
		{
			desc:     "should return DB error",
			parcelID: 1,
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().FetchParcelByID(gomock.Any(), parcel.ID).Return(model.Parcel{}, errors.New("db-error"))
				return r
			},
			expErr:    errors.New("db-error"),
			expParcel: model.Parcel{},
		},
		{
			desc:     "should return empty parcel list",
			parcelID: 1,
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().FetchParcelByID(gomock.Any(), parcel.ID).Return(model.Parcel{}, nil)
				return r
			},
			expErr:    nil,
			expParcel: model.Parcel{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewService(tc.mockRepo())
			parcel, err := s.GetParcelByID(context.Background(), parcel.ID)
			assert.Equal(t, tc.expErr, err)
			assert.EqualValues(t, tc.expParcel, parcel)
		})
	}
}

func TestService_EditParcel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc     string
		mockRepo func() *mocks.MockParcelRepository
		expErr   error
	}{
		{
			desc: "should return success",
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().UpdateParcel(gomock.Any(), gomock.Any()).Return(nil)
				return r
			},
			expErr: nil,
		},

		{
			desc: "should return db error",
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().UpdateParcel(gomock.Any(), gomock.Any()).Return(errors.New("db-error"))
				return r
			},
			expErr: errors.New("db-error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewService(tc.mockRepo())
			err := s.EditParcel(context.Background(), parcel)
			assert.Equal(t, tc.expErr, err)
		})
	}
}
