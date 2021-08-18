package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"parcel-service/internal/app/model"
	"parcel-service/internal/app/service/mocks"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreatePercel(t *testing.T) {
	payload := `{ "user_id":1, "source_address":"Dhaka Bangladesh", "destination_address":"Pabna Shadar", "source_time":"2021-10-10 10:10:12", "type":"Document" }`
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc          string
		payload       string
		mockSvc       func() *mocks.MockParcelService
		expStatusCode int
		expResponse   string
	}{
		{
			desc:    "should success",
			payload: payload,
			mockSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().CreateParcel(gomock.Any(), gomock.Any()).Return(nil)
				return s
			},
			expStatusCode: http.StatusCreated,
			expResponse:   `{"success":true,"errors":null,"data":"successful"}`,
		},
		{
			desc:    "should return decode error",
			payload: `------------`,
			mockSvc: func() *mocks.MockParcelService {
				return mocks.NewMockParcelService(ctrl)
			},
			expStatusCode: http.StatusUnprocessableEntity,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid character '-' in numeric literal","message_title":"Decode Error","severity":"error"}],"data":null}`,
		},
		{
			desc:    "should return invalid parcel error",
			payload: payload,
			mockSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().CreateParcel(gomock.Any(), gomock.Any()).Return(model.ErrInvalid)
				return s
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid","message_title":"invalid parcel","severity":"error"}],"data":null}`,
		},
		{
			desc:    "should return internal server error",
			payload: payload,
			mockSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().CreateParcel(gomock.Any(), gomock.Any()).Return(errors.New("server-error"))
				return s
			},
			expStatusCode: http.StatusInternalServerError,
			expResponse:   `{"success":false,"errors":[{"code":"SERVER_ERROR","message":"server-error","message_title":"failed to create parcel","severity":"error"}],"data":null}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewServer(":8080", tc.mockSvc())

			w := httptest.NewRecorder()
			body := strings.NewReader(tc.payload)
			r := httptest.NewRequest(http.MethodPost, "/api/v1/parcel", body)

			router := mux.NewRouter()
			router.Methods(http.MethodPost).Path("/api/v1/parcel").HandlerFunc(s.createParcel)
			router.ServeHTTP(w, r)
			assert.Equal(t, tc.expStatusCode, w.Code)
			assert.Equal(t, tc.expResponse, w.Body.String())
		})
	}
}
