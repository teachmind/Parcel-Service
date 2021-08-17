package server

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"parcel-service/internal/app/service/mocks"
	"strings"
	"testing"
)

func TestParcelCarrierAccept(t *testing.T)  {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc 			string
		payload			string
		mockSvc			func() *mocks.MockCarrierParcelAcceptService
		expStatusCode 	int
		expResponse 	string
	} {
		{
			desc:    "should success",
			payload: `{ "carrier_id": 2}`,
			mockSvc: func() *mocks.MockCarrierParcelAcceptService {
				s := mocks.NewMockCarrierParcelAcceptService(ctrl)
				s.EXPECT().AssignCarrierToParcel(gomock.Any(), gomock.Any()).Return(nil)
				return s
			},
			expStatusCode: http.StatusCreated,
			expResponse:   `{"success":true,"errors":null,"data":"Successful"}`,
		},
		{
			desc:    "should return decode error",
			payload: `------------`,
			mockSvc: func() *mocks.MockCarrierParcelAcceptService {
				return mocks.NewMockCarrierParcelAcceptService(ctrl)
			},
			expStatusCode: http.StatusUnprocessableEntity,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid character '-' in numeric literal","message_title":"Decode Error","severity":"error"}],"data":null}`,
		},
		{
			desc:    "should return invalid carrier id",
			payload: `{}`,
			mockSvc: func() *mocks.MockCarrierParcelAcceptService {
				s := mocks.NewMockCarrierParcelAcceptService(ctrl)
				return s
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"Carrier ID is required :empty","message_title":"Invalid Input","severity":"error"}],"data":null}`,
		},
		{
			desc:    "should return internal server error",
			payload: `{ "carrier_id": 2 }`,
			mockSvc: func() *mocks.MockCarrierParcelAcceptService {
				s := mocks.NewMockCarrierParcelAcceptService(ctrl)
				s.EXPECT().AssignCarrierToParcel(gomock.Any(), gomock.Any()).Return(errors.New("server-error"))
				return s
			},
			expStatusCode: http.StatusInternalServerError,
			expResponse:   `{"success":false,"errors":[{"code":"SERVER_ERROR","message":"server-error","message_title":"failed to assign carrier to parcel","severity":"error"}],"data":null}`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewServer(":8080", tc.mockSvc())

			w := httptest.NewRecorder()
			body := strings.NewReader(tc.payload)
			r := httptest.NewRequest(http.MethodPost, "/api/v1/parcel/1/accept", body)
			r = mux.SetURLVars(r, map[string]string{"id": "1"}) //to get the id from route
			router := mux.NewRouter()
			router.Methods(http.MethodPost).Path("/api/v1/parcel/{id}/accept").HandlerFunc(s.parcelCarrierAccept)
			router.ServeHTTP(w, r)
			assert.Equal(t, tc.expStatusCode, w.Code)
			assert.Equal(t, tc.expResponse, w.Body.String())
		})
	}
}