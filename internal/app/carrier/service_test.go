package carrier

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"parcel-service/internal/app/model"
	"parcel-service/internal/app/service/mocks"
	"testing"
)

func TestService_AssignCarrierToParcel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc     string
		payload  model.CarrierRequest
		mockRepo func() *mocks.MockCarrierParcelAcceptRepository
		expErr   error
	}{
		{
			desc: "should return success",
			payload: model.CarrierRequest{
				ParcelID: 1,
				CarrierID:    2,
				Status:  1,
			},
			mockRepo: func() *mocks.MockCarrierParcelAcceptRepository {
				r := mocks.NewMockCarrierParcelAcceptRepository(ctrl)
				r.EXPECT().UpdateCarrierRequest(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				return r
			},
			expErr: nil,
		},{
			desc: "should return db-error",
			payload: model.CarrierRequest{
				ParcelID: 1,
				CarrierID:    2,
				Status:  1,
			},
			mockRepo: func() *mocks.MockCarrierParcelAcceptRepository {
				r := mocks.NewMockCarrierParcelAcceptRepository(ctrl)
				r.EXPECT().UpdateCarrierRequest(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("db-error"))
				return r
			},
			expErr: errors.New("db-error"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewService(tc.mockRepo())
			err := s.AssignCarrierToParcel(context.Background(), tc.payload)
			assert.Equal(t, tc.expErr, err)
		})
	}
}
