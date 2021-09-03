package carrier

import (
	"context"
	"errors"
	"parcel-service/internal/app/model"
	"parcel-service/internal/app/service/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_NewCarrierRequest(t *testing.T) {
	addReqPayload := model.CarrierRequest{
		CarrierID: 1,
		ParcelID:  1,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc     string
		payload  model.CarrierRequest
		mockRepo func() *mocks.MockCarrierRepository
		expErr   error
	}{
		{
			desc:    "should return success",
			payload: addReqPayload,
			mockRepo: func() *mocks.MockCarrierRepository {
				r := mocks.NewMockCarrierRepository(ctrl)
				r.EXPECT().InsertCarrierRequest(gomock.Any(), gomock.Any()).Return(nil)
				return r
			},
			expErr: nil,
		},

		{
			desc:    "should return db error",
			payload: addReqPayload,
			mockRepo: func() *mocks.MockCarrierRepository {
				r := mocks.NewMockCarrierRepository(ctrl)
				r.EXPECT().InsertCarrierRequest(gomock.Any(), gomock.Any()).Return(errors.New("db-error"))
				return r
			},
			expErr: errors.New("db-error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewService(tc.mockRepo())
			err := s.NewCarrierRequest(context.Background(), tc.payload)
			assert.Equal(t, tc.expErr, err)
		})
	}
}
