package parcel

import (
	"context"
	"errors"
	"fmt"
	"parcel-service/internal/app/model"
	"parcel-service/internal/app/service/mocks"
	"parcel-service/internal/app/util"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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
				PhoneNumber: "01738799349",
				Password:    "123456",
				CategoryID:  1,
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
				PhoneNumber: "01738799349",
				Password:    "12345",
				CategoryID:  1,
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

func TestService_GetParcelByPhoneAndPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	password, _ := util.HashPassword("123456")
	parcel := model.Parcel{
		ID:         1,
		CarrierID:  "01738799349",
		CarrierFee: password,
		CompanyFee: 1,
	}

	testCases := []struct {
		desc      string
		phone     string
		password  string
		mockRepo  func() *mocks.MockParcelRepository
		expErr    error
		expParcel model.Parcel
	}{
		{
			desc:     "should return success",
			phone:    "01738799349",
			password: "123456",
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().GetParcelByPhone(gomock.Any(), "01738799349").Return(parcel, nil)
				return r
			},
			expErr:    nil,
			expParcel: parcel,
		},
		{
			desc:     "should return invalid request error",
			phone:    "",
			password: "",
			mockRepo: func() *mocks.MockParcelRepository {
				return mocks.NewMockParcelRepository(ctrl)
			},
			expErr:    fmt.Errorf("invalid login request :%w", model.ErrInvalid),
			expParcel: model.Parcel{},
		},
		{
			desc:     "should return DB error",
			phone:    "01738799349",
			password: "123456",
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().GetParcelByPhone(gomock.Any(), "01738799349").Return(model.Parcel{}, errors.New("db-error"))
				return r
			},
			expErr:    errors.New("db-error"),
			expParcel: model.Parcel{},
		},
		{
			desc:     "should return wrong password error",
			phone:    "01738799349",
			password: "wrong-password",
			mockRepo: func() *mocks.MockParcelRepository {
				r := mocks.NewMockParcelRepository(ctrl)
				r.EXPECT().GetParcelByPhone(gomock.Any(), "01738799349").Return(parcel, nil)
				return r
			},
			expErr:    fmt.Errorf("wrong password :%w", model.ErrInvalid),
			expParcel: model.Parcel{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewService(tc.mockRepo())
			parcel, err := s.GetParcelByPhoneAndPassword(context.Background(), tc.phone, tc.password)
			assert.Equal(t, tc.expErr, err)
			assert.EqualValues(t, tc.expParcel, parcel)
		})
	}
}
