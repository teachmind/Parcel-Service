package parcel

import (
	"context"
	"errors"
	"fmt"
	"parcel-service/internal/app/model"
	mock_service "parcel-service/internal/app/service/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

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
		mockRepo  func() *mock_service.MockParcelRepository
		expErr    error
		expParcel []model.Parcel
	}{
		{
			desc:   "should return success",
			status: 1,
			offset: 2,
			limit:  4,
			mockRepo: func() *mock_service.MockParcelRepository {
				r := mock_service.NewMockParcelRepository(ctrl)
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
			mockRepo: func() *mock_service.MockParcelRepository {
				r := mock_service.NewMockParcelRepository(ctrl)
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
			mockRepo: func() *mock_service.MockParcelRepository {
				r := mock_service.NewMockParcelRepository(ctrl)
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
			mockRepo: func() *mock_service.MockParcelRepository {
				r := mock_service.NewMockParcelRepository(ctrl)
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
			fmt.Println(parcels)
			fmt.Println("heloooo")
			fmt.Println(tc.expParcel)
			fmt.Println("heloooo    2")
			assert.Equal(t, tc.expErr, err)
			assert.EqualValues(t, tc.expParcel, parcels)
		})
	}
}
