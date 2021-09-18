package server

import (
	"errors"
	"fmt"
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
				s.EXPECT().CreateParcel(gomock.Any(), gomock.Any()).Return(parcel, nil)
				return s
			},
			expStatusCode: http.StatusCreated,
			expResponse:   `{"success":true,"errors":null,"data":{"id":0,"user_id":1,"carrier_id":0,"status":0,"source_address":"Dhaka Bangladesh","destination_address":"Pabna Shadar","source_time":"2020-04-11T21:34:01Z","type":"Document","price":200,"carrier_fee":180,"company_fee":20,"created_at":"2020-04-11T21:34:01Z","updated_at":"2020-04-11T21:34:01Z"}}`,
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
			desc:    "should return invalid input",
			payload: `{ "user_id":1 }`,
			mockParcelSvc: func() *mocks.MockParcelService {
				return mocks.NewMockParcelService(ctrl)
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"source Address is required :empty","message_title":"Invalid Input","severity":"error"}],"data":null}`,
		},
		{
			desc:    "should return invalid parcel error",
			payload: payload,
			mockParcelSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().CreateParcel(gomock.Any(), gomock.Any()).Return(model.Parcel{}, model.ErrInvalid)
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
				s.EXPECT().CreateParcel(gomock.Any(), gomock.Any()).Return(model.Parcel{}, errors.New("server-error"))
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

func TestAddCarrierRequest(t *testing.T) {
	payload := `{ "carrier_id":1 }`
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	parcelId := map[string]string{"valid": "1", "invalid": "invalid"}

	testCases := []struct {
		desc           string
		payload        string
		mockCarrierSvc func() *mocks.MockCarrierService
		parcelId       string
		expStatusCode  int
		expResponse    string
	}{
		{
			desc:     "should success",
			payload:  payload,
			parcelId: parcelId["valid"],
			mockCarrierSvc: func() *mocks.MockCarrierService {
				s := mocks.NewMockCarrierService(ctrl)
				s.EXPECT().NewCarrierRequest(gomock.Any(), gomock.Any()).Return(nil)
				return s
			},
			expStatusCode: http.StatusCreated,
			expResponse:   `{"success":true,"errors":null,"data":"Success"}`,
		},
		{
			desc:     "should return decode error",
			payload:  `------------`,
			parcelId: parcelId["valid"],
			mockCarrierSvc: func() *mocks.MockCarrierService {
				return mocks.NewMockCarrierService(ctrl)
			},
			expStatusCode: http.StatusUnprocessableEntity,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid character '-' in numeric literal","message_title":"Decode Error","severity":"error"}],"data":null}`,
		},
		{
			desc:     "should return invalid input",
			payload:  `{ "carrier_id":0 }`,
			parcelId: parcelId["valid"],
			mockCarrierSvc: func() *mocks.MockCarrierService {
				return mocks.NewMockCarrierService(ctrl)
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"Carrier ID is required :empty","message_title":"Invalid Input","severity":"error"}],"data":null}`,
		},
		{
			desc:     "should return invalid request error",
			payload:  payload,
			parcelId: parcelId["valid"],
			mockCarrierSvc: func() *mocks.MockCarrierService {
				s := mocks.NewMockCarrierService(ctrl)
				s.EXPECT().NewCarrierRequest(gomock.Any(), gomock.Any()).Return(model.ErrInvalid)
				return s
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid","message_title":"invalid Request","severity":"error"}],"data":null}`,
		},
		{
			desc:     "should return internal server error",
			payload:  payload,
			parcelId: parcelId["valid"],
			mockCarrierSvc: func() *mocks.MockCarrierService {
				s := mocks.NewMockCarrierService(ctrl)
				s.EXPECT().NewCarrierRequest(gomock.Any(), gomock.Any()).Return(errors.New("server-error"))
				return s
			},
			expStatusCode: http.StatusInternalServerError,
			expResponse:   `{"success":false,"errors":[{"code":"SERVER_ERROR","message":"server-error","message_title":"failed to add new carrier request","severity":"error"}],"data":null}`,
		},
		{
			desc:     "should return invalid parcel ID",
			payload:  payload,
			parcelId: parcelId["invalid"],
			mockCarrierSvc: func() *mocks.MockCarrierService {
				return mocks.NewMockCarrierService(ctrl)
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"strconv.Atoi: parsing \"invalid\": invalid syntax","message_title":"Invalid Parcel ID","severity":"error"}],"data":null}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewServer(":8080", nil, tc.mockCarrierSvc())

			w := httptest.NewRecorder()
			body := strings.NewReader(tc.payload)
			r := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/parcel/%s/request", tc.parcelId), body)
			r = mux.SetURLVars(r, map[string]string{"id": tc.parcelId})

			router := mux.NewRouter()
			router.Methods(http.MethodPost).Path("/api/v1/parcel/{id}/request").HandlerFunc(s.addCarrierRequest)
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
func TestEditParcel(t *testing.T) {
	payload := `{ "status":1 }`
	parcelId := map[string]string{"valid": "1", "invalid": "invalid"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc          string
		payload       string
		parcelId      string
		mockParcelSvc func() *mocks.MockParcelService
		expStatusCode int
		expResponse   string
	}{
		{
			desc:     "should success",
			parcelId: parcelId["valid"],
			payload:  payload,
			mockParcelSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().EditParcel(gomock.Any(), gomock.Any()).Return(nil)
				return s
			},
			expStatusCode: http.StatusCreated,
			expResponse:   `{"success":true,"errors":null,"data":"Success"}`,
		},
		{
			desc:     "should return decode error",
			parcelId: parcelId["valid"],
			payload:  `------------`,
			mockParcelSvc: func() *mocks.MockParcelService {
				return mocks.NewMockParcelService(ctrl)
			},
			expStatusCode: http.StatusUnprocessableEntity,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid character '-' in numeric literal","message_title":"Decode Error","severity":"error"}],"data":null}`,
		},
		{
			desc:     "should return invalid request error",
			parcelId: parcelId["valid"],
			payload:  payload,
			mockParcelSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().EditParcel(gomock.Any(), gomock.Any()).Return(model.ErrInvalid)
				return s
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid","message_title":"invalid Request","severity":"error"}],"data":null}`,
		},
		{
			desc:     "should return internal server error",
			parcelId: parcelId["valid"],
			payload:  payload,
			mockParcelSvc: func() *mocks.MockParcelService {
				s := mocks.NewMockParcelService(ctrl)
				s.EXPECT().EditParcel(gomock.Any(), gomock.Any()).Return(errors.New("server-error"))
				return s
			},
			expStatusCode: http.StatusInternalServerError,
			expResponse:   `{"success":false,"errors":[{"code":"SERVER_ERROR","message":"server-error","message_title":"failed to update parcel","severity":"error"}],"data":null}`,
		},
		{
			desc:     "should return invalid parcel ID",
			payload:  payload,
			parcelId: parcelId["invalid"],
			mockParcelSvc: func() *mocks.MockParcelService {
				return mocks.NewMockParcelService(ctrl)
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"strconv.Atoi: parsing \"invalid\": invalid syntax","message_title":"Invalid Parcel ID","severity":"error"}],"data":null}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewServer(":8080", tc.mockParcelSvc(), nil)

			w := httptest.NewRecorder()
			body := strings.NewReader(tc.payload)

			r := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/parcel/%s", tc.parcelId), body)
			r = mux.SetURLVars(r, map[string]string{"id": tc.parcelId})

			router := mux.NewRouter()
			router.Methods(http.MethodPut).Path("/api/v1/parcel/{id}").HandlerFunc(s.editParcel)
			router.ServeHTTP(w, r)
			assert.Equal(t, tc.expStatusCode, w.Code)
			assert.Equal(t, tc.expResponse, w.Body.String())
		})
	}
}

func TestParcelCarrierAccept(t *testing.T) {
	parcelId := map[string]string{"valid": "1", "invalid": "invalid"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc          string
		payload       string
		parcelId      string
		mockSvc       func() *mocks.MockCarrierService
		expStatusCode int
		expResponse   string
	}{
		{
			desc:     "should success",
			parcelId: parcelId["valid"],
			payload:  `{ "carrier_id": 2}`,
			mockSvc: func() *mocks.MockCarrierService {
				s := mocks.NewMockCarrierService(ctrl)
				s.EXPECT().AssignCarrierToParcel(gomock.Any(), gomock.Any()).Return(nil)
				return s
			},
			expStatusCode: http.StatusCreated,
			expResponse:   `{"success":true,"errors":null,"data":"Successful"}`,
		},
		{
			desc:     "should return decode error",
			parcelId: parcelId["valid"],
			payload:  `------------`,
			mockSvc: func() *mocks.MockCarrierService {
				return mocks.NewMockCarrierService(ctrl)
			},
			expStatusCode: http.StatusUnprocessableEntity,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid character '-' in numeric literal","message_title":"Decode Error","severity":"error"}],"data":null}`,
		},
		{
			desc:     "should return invalid carrier id",
			parcelId: parcelId["valid"],
			payload:  `{}`,
			mockSvc: func() *mocks.MockCarrierService {
				s := mocks.NewMockCarrierService(ctrl)
				return s
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"Carrier ID is required :empty","message_title":"Invalid Input","severity":"error"}],"data":null}`,
		},
		{
			desc:     "should return internal server error",
			parcelId: parcelId["valid"],
			payload:  `{ "carrier_id": 2 }`,
			mockSvc: func() *mocks.MockCarrierService {
				s := mocks.NewMockCarrierService(ctrl)
				s.EXPECT().AssignCarrierToParcel(gomock.Any(), gomock.Any()).Return(errors.New("server-error"))
				return s
			},
			expStatusCode: http.StatusInternalServerError,
			expResponse:   `{"success":false,"errors":[{"code":"SERVER_ERROR","message":"server-error","message_title":"failed to assign carrier to parcel","severity":"error"}],"data":null}`,
		},
		{
			desc:     "should return invalid parcel ID",
			payload:  `{ "carrier_id": 2 }`,
			parcelId: parcelId["invalid"],
			mockSvc: func() *mocks.MockCarrierService {
				return mocks.NewMockCarrierService(ctrl)
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"strconv.Atoi: parsing \"invalid\": invalid syntax","message_title":"Invalid Parcel ID","severity":"error"}],"data":null}`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewServer(":8080", nil, tc.mockSvc())

			w := httptest.NewRecorder()
			body := strings.NewReader(tc.payload)
			r := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/parcel/%s/accept", tc.parcelId), body)
			r = mux.SetURLVars(r, map[string]string{"id": tc.parcelId}) //to get the id from route
			router := mux.NewRouter()
			router.Methods(http.MethodPost).Path("/api/v1/parcel/{id}/accept").HandlerFunc(s.parcelCarrierAccept)
			router.ServeHTTP(w, r)
			assert.Equal(t, tc.expStatusCode, w.Code)
			assert.Equal(t, tc.expResponse, w.Body.String())
		})
	}
}
