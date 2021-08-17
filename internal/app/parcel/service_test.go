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

func TestService_CreateParcel(t *testing.T) {
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
