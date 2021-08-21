package parcel

import (
	"context"
	"errors"
	"parcel-service/internal/app/model"
	"parcel-service/internal/app/service/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	payload = `{ "user_id":1, "source_address":"Dhaka Bangladesh", "destination_address":"Pabna Shadar", "source_time":"2021-10-10 10:10:12", "type":"Document" }`
)

func TestService_InsertParcel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc     string
		payload  model.Parcel
		mockRepo func() *mocks.MockParcelRepository
		expErr   error
	}{
		{
			desc: "should return success",
			payload: model.Parcel{
				UserID:             1,
				SourceAddress:      "Dhaka Bangladesh",
				DestinationAddress: "Pabna Shadar",
				SourceTime:         "2021-10-10 10: 10: 12",
				ParcelType:         "Document",
				Price:              200,
				CarrierFee:         180,
				CompanyFee:         20,
			},
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().InsertParcel(gomock.Any(), gomock.Any()).Return(nil)
				return r
			},
			expErr: nil,
		},

		{
			desc: "should return db error",
			payload: model.Parcel{
				UserID:             1,
				SourceAddress:      "Dhaka Bangladesh",
				DestinationAddress: "Pabna Shadar",
				SourceTime:         "2021-10-10 10: 10: 12",
				ParcelType:         "Document",
				Price:              200,
				CarrierFee:         180,
				CompanyFee:         20,
			},
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
			err := s.CreateParcel(context.Background(), tc.payload)
			assert.Equal(t, tc.expErr, err)
		})
	}
}
func TestService_GetParcelByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	parcel := model.Parcel{
		ID:                 1,
		UserID:             1,
		SourceAddress:      "Dhaka Bangladesh",
		DestinationAddress: "Pabna Shadar",
		SourceTime:         "2021-10-10 10: 10: 12",
		ParcelType:         "Document",
		Price:              200,
		CarrierFee:         180,
		CompanyFee:         20,
	}

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
				r.EXPECT().FetchParcelByID(gomock.Any(), 1).Return(parcel, nil)
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
				r.EXPECT().FetchParcelByID(gomock.Any(), 1).Return(model.Parcel{}, model.ErrNotFound)
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
				r.EXPECT().FetchParcelByID(gomock.Any(), 1).Return(model.Parcel{}, errors.New("db-error"))
				return r
			},
			expErr:    errors.New("db-error"),
			expParcel: model.Parcel{},
		},
		{
			desc:     "should return empty parcel",
			parcelID: 1,
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().FetchParcelByID(gomock.Any(), 1).Return(model.Parcel{}, nil)
				return r
			},
			expErr:    nil,
			expParcel: model.Parcel{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewService(tc.mockRepo())
			user, err := s.GetParcelByID(context.Background(), 1)
			assert.Equal(t, tc.expErr, err)
			assert.EqualValues(t, tc.expParcel, user)
		})
	}
}
