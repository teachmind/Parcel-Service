package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"parcel-service/internal/app/model"
	"parcel-service/internal/app/service/mocks"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestNewParcel(t *testing.T) {
	payload := `{ "user_id":1, "source_address":"Dhaka Bangladesh", "destination_address":"Pabna Shadar", "source_time":"3021-10-10T10:10:12Z", "type":"Document" }`
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc          string
		payload       string
		mockParcelSvc func() *mocks.MockParcelService
		expStatusCode int
		expResponse   string
	}{
		{
			desc:    "should success",
			payload: payload,
			mockParcelSvc: func() *mocks.MockParcelService {
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
			mockParcelSvc: func() *mocks.MockParcelService {
				return mocks.NewMockParcelService(ctrl)
			},
			expStatusCode: http.StatusUnprocessableEntity,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid character '-' in numeric literal","message_title":"Decode Error","severity":"error"}],"data":null}`,
		},
		{
			desc:    "should return invalid parcel error",
			payload: payload,
			mockParcelSvc: func() *mocks.MockParcelService {
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
			mockParcelSvc: func() *mocks.MockParcelService {
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
			s := NewServer(":8080", tc.mockParcelSvc(), nil)

			w := httptest.NewRecorder()
			body := strings.NewReader(tc.payload)
			r := httptest.NewRequest(http.MethodPost, "/api/v1/parcel", body)

			router := mux.NewRouter()
			router.Methods(http.MethodPost).Path("/api/v1/parcel").HandlerFunc(s.newParcel)
			router.ServeHTTP(w, r)
			assert.Equal(t, tc.expStatusCode, w.Code)
			assert.Equal(t, tc.expResponse, w.Body.String())
		})
	}
}

func TestGetParcel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	parcel := model.Parcel{
		UserID:             1,
		SourceAddress:      "Dhaka Bangladesh",
		DestinationAddress: "Pabna Shadar",
		SourceTime:         time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
		ParcelType:         "Document",
		Price:              200.0,
		CarrierFee:         180.0,
		CompanyFee:         20.0,
		CreatedAt:          time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
		UpdatedAt:          time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
	}

	testCases := []struct {
		desc          string
		mockParcelSvc func() *mocks.MockParcelService
		parcelID      string
		expStatusCode int
		expResponse   string
	}{
		{
			desc: "should success",
			mockParcelSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().GetParcelByID(gomock.Any(), 1).Return(parcel, nil)
				return s
			},
			parcelID:      "1",
			expStatusCode: http.StatusOK,
			expResponse:   `{"success":true,"errors":null,"data":{"id":0,"user_id":1,"carrier_id":0,"status":0,"source_address":"Dhaka Bangladesh","destination_address":"Pabna Shadar","source_time":"2020-04-11T21:34:01Z","type":"Document","price":200,"carrier_fee":180,"company_fee":20,"created_at":"2020-04-11T21:34:01Z","updated_at":"2020-04-11T21:34:01Z"}}`,
		},
		{
			desc: "should return ID not exist",
			mockParcelSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().GetParcelByID(gomock.Any(), 1).Return(model.Parcel{}, model.ErrInvalid)
				return s
			},
			parcelID:      "1",
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid","message_title":"This ID does not exist.","severity":"error"}],"data":null}`,
		},
		{
			desc: "should return internal server error",
			mockParcelSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().GetParcelByID(gomock.Any(), 1).
					Return(model.Parcel{}, errors.New("server-error"))
				return s
			},
			parcelID:      "1",
			expStatusCode: http.StatusInternalServerError,
			expResponse:   `{"success":false,"errors":[{"code":"SERVER_ERROR","message":"server-error","message_title":"Failed to fetch parcel 1","severity":"error"}],"data":null}`,
		},
		{
			desc: "should return internal server error",
			mockParcelSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().GetParcelByID(gomock.Any(), 1).
					Return(model.Parcel{}, model.ErrNotFound)
				return s
			},
			parcelID:      "1",
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"not found","message_title":"This ID does not exist.","severity":"error"}],"data":null}`,
		},
		{
			desc: "should return invalid parcel ID",
			mockParcelSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				return s
			},
			parcelID:      "__",
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"strconv.Atoi: parsing \"__\": invalid syntax","message_title":"Invalid Parcel ID","severity":"error"}],"data":null}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewServer(":8080", tc.mockParcelSvc(), nil)

			w := httptest.NewRecorder()
			body := strings.NewReader("")
			r := httptest.NewRequest(http.MethodGet, "/api/v1/parcel/"+tc.parcelID, body)
			r = mux.SetURLVars(r, map[string]string{"id": tc.parcelID})

			router := mux.NewRouter()
			router.Methods(http.MethodGet).Path("/api/v1/parcel/{id}").HandlerFunc(s.getParcel)
			router.ServeHTTP(w, r)
			assert.Equal(t, tc.expStatusCode, w.Code)
			assert.Equal(t, tc.expResponse, w.Body.String())
		})
	}
}

func TestAddCarrierRequest(t *testing.T) {
	payload := `{ "carrier_id":1, "parcel_id":1 }`
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc           string
		payload        string
		mockCarrierSvc func() *mocks.MockCarrierService
		expStatusCode  int
		expResponse    string
	}{
		{
			desc:    "should success",
			payload: payload,
			mockCarrierSvc: func() *mocks.MockCarrierService {
				s := mocks.NewMockCarrierService(ctrl)
				s.EXPECT().NewCarrierRequest(gomock.Any(), gomock.Any()).Return(nil)
				return s
			},
			expStatusCode: http.StatusCreated,
			expResponse:   `{"success":true,"errors":null,"data":"Success"}`,
		},
		{
			desc:    "should return decode error",
			payload: `------------`,
			mockCarrierSvc: func() *mocks.MockCarrierService {
				return mocks.NewMockCarrierService(ctrl)
			},
			expStatusCode: http.StatusUnprocessableEntity,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid character '-' in numeric literal","message_title":"Decode Error","severity":"error"}],"data":null}`,
		},
		{
			desc:    "should return invalid request error",
			payload: payload,
			mockCarrierSvc: func() *mocks.MockCarrierService {
				s := mocks.NewMockCarrierService(ctrl)
				s.EXPECT().NewCarrierRequest(gomock.Any(), gomock.Any()).Return(model.ErrInvalid)
				return s
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid","message_title":"invalid Request","severity":"error"}],"data":null}`,
		},
		{
			desc:    "should return internal server error",
			payload: payload,
			mockCarrierSvc: func() *mocks.MockCarrierService {
				s := mocks.NewMockCarrierService(ctrl)
				s.EXPECT().NewCarrierRequest(gomock.Any(), gomock.Any()).Return(errors.New("server-error"))
				return s
			},
			expStatusCode: http.StatusInternalServerError,
			expResponse:   `{"success":false,"errors":[{"code":"SERVER_ERROR","message":"server-error","message_title":"failed to add new carrier request","severity":"error"}],"data":null}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewServer(":8080", nil, tc.mockCarrierSvc())

			w := httptest.NewRecorder()
			body := strings.NewReader(tc.payload)
			r := httptest.NewRequest(http.MethodPost, "/api/v1/parcel/1/request", body)
			r = mux.SetURLVars(r, map[string]string{"id": "1"})

			router := mux.NewRouter()
			router.Methods(http.MethodPost).Path("/api/v1/parcel/{id}/request").HandlerFunc(s.addCarrierRequest)
			router.ServeHTTP(w, r)
			assert.Equal(t, tc.expStatusCode, w.Code)
			assert.Equal(t, tc.expResponse, w.Body.String())
		})
	}
}

func TestGetPercels(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	parcels := []model.Parcel{
		{
			UserID:             1,
			Status:             1,
			SourceAddress:      "Dhaka Bangladesh",
			DestinationAddress: "Pabna Shadar",
			SourceTime:         time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
			ParcelType:         "Document",
			Price:              200.0,
			CarrierFee:         180.0,
			CompanyFee:         20.0,
			CreatedAt:          time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
			UpdatedAt:          time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
		}, {
			UserID:             2,
			Status:             1,
			SourceAddress:      "Dhaka Bangladesh",
			DestinationAddress: "Pabna Shadar",
			SourceTime:         time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
			ParcelType:         "Document",
			Price:              200.0,
			CarrierFee:         180.0,
			CompanyFee:         20.0,
			CreatedAt:          time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
			UpdatedAt:          time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
		},
	}

	testCases := []struct {
		desc          string
		mockParcelSvc func() *mocks.MockParcelService
		status        string
		offset        string
		limit         string
		expStatusCode int
		expResponse   string
	}{
		{
			desc: "should success",
			mockParcelSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().GetParcels(gomock.Any(), 1, 2, 0).Return(parcels, nil)
				return s
			},
			status:        "1",
			offset:        "0",
			limit:         "2",
			expStatusCode: http.StatusOK,
			expResponse:   `{"success":true,"errors":null,"data":[{"id":0,"user_id":1,"carrier_id":0,"status":1,"source_address":"Dhaka Bangladesh","destination_address":"Pabna Shadar","source_time":"2020-04-11T21:34:01Z","type":"Document","price":200,"carrier_fee":180,"company_fee":20,"created_at":"2020-04-11T21:34:01Z","updated_at":"2020-04-11T21:34:01Z"},{"id":0,"user_id":2,"carrier_id":0,"status":1,"source_address":"Dhaka Bangladesh","destination_address":"Pabna Shadar","source_time":"2020-04-11T21:34:01Z","type":"Document","price":200,"carrier_fee":180,"company_fee":20,"created_at":"2020-04-11T21:34:01Z","updated_at":"2020-04-11T21:34:01Z"}]}`,
		},
		{
			desc: "should return empty parcel list",
			mockParcelSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().GetParcels(gomock.Any(), 0, 2, 0).Return(nil, nil)
				return s
			},
			status:        "0",
			offset:        "0",
			limit:         "2",
			expStatusCode: http.StatusOK,
			expResponse:   `{"success":true,"errors":null,"data":null}`,
		},
		{
			desc: "should return internal server error",
			mockParcelSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().GetParcels(gomock.Any(), 1, 2, -1).Return([]model.Parcel{}, errors.New("pq: OFFSET must not be negative"))
				return s
			},
			status:        "1",
			offset:        "-1",
			limit:         "2",
			expStatusCode: http.StatusInternalServerError,
			expResponse:   `{"success":false,"errors":[{"code":"SERVER_ERROR","message":"pq: OFFSET must not be negative","message_title":"Failed to fetch parcel list for given query params","severity":"error"}],"data":null}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewServer(":8080", tc.mockParcelSvc(), nil)

			w := httptest.NewRecorder()
			body := strings.NewReader("")
			r := httptest.NewRequest(http.MethodGet, "/api/v1/parcel?status="+tc.status+"&offset="+tc.offset+"&limit="+tc.limit, body)

			router := mux.NewRouter()
			router.Methods(http.MethodGet).Path("/api/v1/parcel").Queries("status", tc.status, "offset", tc.offset, "limit", tc.limit).HandlerFunc(s.getParcelList)
			router.ServeHTTP(w, r)
			assert.Equal(t, tc.expStatusCode, w.Code)
			assert.Equal(t, tc.expResponse, w.Body.String())
		})
	}
}
func TestEditParcel(t *testing.T) {
	payload := `{ "status":1 }`
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc          string
		payload       string
		mockParcelSvc func() *mocks.MockParcelService
		expStatusCode int
		expResponse   string
	}{
		{
			desc:    "should success",
			payload: payload,
			mockParcelSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().EditParcel(gomock.Any(), gomock.Any()).Return(nil)
				return s
			},
			expStatusCode: http.StatusCreated,
			expResponse:   `{"success":true,"errors":null,"data":"Success"}`,
		},
		{
			desc:    "should return decode error",
			payload: `------------`,
			mockParcelSvc: func() *mocks.MockParcelService {
				return mocks.NewMockParcelService(ctrl)
			},
			expStatusCode: http.StatusUnprocessableEntity,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid character '-' in numeric literal","message_title":"Decode Error","severity":"error"}],"data":null}`,
		},
		{
			desc:    "should return invalid request error",
			payload: payload,
			mockParcelSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().EditParcel(gomock.Any(), gomock.Any()).Return(model.ErrInvalid)
				return s
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid","message_title":"invalid Request","severity":"error"}],"data":null}`,
		},
		{
			desc:    "should return internal server error",
			payload: payload,
			mockParcelSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().EditParcel(gomock.Any(), gomock.Any()).Return(errors.New("server-error"))
				return s
			},
			expStatusCode: http.StatusInternalServerError,
			expResponse:   `{"success":false,"errors":[{"code":"SERVER_ERROR","message":"server-error","message_title":"failed to update parcel","severity":"error"}],"data":null}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewServer(":8080", tc.mockParcelSvc(), nil)

			w := httptest.NewRecorder()
			body := strings.NewReader(tc.payload)
			r := httptest.NewRequest(http.MethodPut, "/api/v1/parcel/1", body)
			r = mux.SetURLVars(r, map[string]string{"id": "1"})

			router := mux.NewRouter()
			router.Methods(http.MethodPut).Path("/api/v1/parcel/{id}").HandlerFunc(s.editParcel)
			router.ServeHTTP(w, r)
			assert.Equal(t, tc.expStatusCode, w.Code)
			assert.Equal(t, tc.expResponse, w.Body.String())
		})
	}
}
