package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"parcel-service/internal/app/model"
	mock_service "parcel-service/internal/app/service/mocks"
	"strconv"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetPercels(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc          string
		status        int
		offset        int
		limit         int
		mockParcelSvc func() *mock_service.MockParcelService
		expStatusCode int
		expResponse   string
	}{
		{
			desc:   "should success",
			status: 1,
			offset: 2,
			limit:  4,
			mockParcelSvc: func() *mock_service.MockParcelService {
				s := mock_service.NewMockParcelService(ctrl)
				s.EXPECT().GetParcels(gomock.Any(), 1, 2, 4).Return(nil)
				return s
			},
			expStatusCode: http.StatusCreated,
			expResponse:   `{"success":true,"errors":null,"data":"successful"}`,
		},
		{
			desc:   "should return decode error",
			status: -1,
			offset: 2,
			limit:  4,
			mockParcelSvc: func() *mock_service.MockParcelService {
				return mock_service.NewMockParcelService(ctrl)
			},
			expStatusCode: http.StatusUnprocessableEntity,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid character '-' in numeric literal","message_title":"Decode Error","severity":"error"}],"data":null}`,
		},
		{
			desc:   "should return invalid parcel error",
			status: 1,
			offset: 2,
			limit:  4,
			mockParcelSvc: func() *mock_service.MockParcelService {
				s := mock_service.NewMockParcelService(ctrl)
				s.EXPECT().GetParcels(gomock.Any(), 1, 2, 4).Return(model.ErrInvalid)
				return s
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid","message_title":"invalid parcel","severity":"error"}],"data":null}`,
		},
		{
			desc:   "should return internal server error",
			status: 1,
			offset: 2,
			limit:  4,
			mockParcelSvc: func() *mock_service.MockParcelService {
				s := mock_service.NewMockParcelService(ctrl)
				s.EXPECT().GetParcels(gomock.Any(), 1, 2, 4).Return(errors.New("server-error"))
				return s
			},
			expStatusCode: http.StatusInternalServerError,
			expResponse:   `{"success":false,"errors":[{"code":"SERVER_ERROR","message":"server-error","message_title":"failed to create parcel","severity":"error"}],"data":null}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewServer(":8080", tc.mockParcelSvc())

			w := httptest.NewRecorder()
			body := strings.NewReader("")
			r := httptest.NewRequest(http.MethodGet, "/api/v1/parcel/"+strconv.Itoa(tc.status)+strconv.Itoa(tc.offset)+strconv.Itoa(tc.limit), body)

			router := mux.NewRouter()
			router.Methods(http.MethodGet).Path("/api/v1/parcel/").Queries("status", "offset", "limit").HandlerFunc(s.getParcelList)
			router.ServeHTTP(w, r)
			assert.Equal(t, tc.expStatusCode, w.Code)
			assert.Equal(t, tc.expResponse, w.Body.String())
		})
	}
}
