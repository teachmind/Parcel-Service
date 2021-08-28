package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"parcel-service/internal/app/model"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func (s *server) getParcelList(w http.ResponseWriter, r *http.Request) {
	status, err := strconv.Atoi(r.URL.Query().Get("status"))
	if err == nil {
		fmt.Printf("%d", status)
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err == nil {
		fmt.Printf("%d", offset)
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err == nil {
		fmt.Printf("%d", limit)
	}

	parcel, err := s.parcelService.GetParcels(r.Context(), status, limit, offset)
	if err != nil {
		// if errors.Is(err, model.ErrInvalid) {
		// 	ErrInvalidEntityResponse(w, "invalid parcel", err)
		// 	return
		// }
		fmt.Print(err)
		log.Error().Err(err).Msgf("[parcel/{id}] failed to parcel '%d': %v", nil, err)
		// ErrInternalServerResponse(w, "failed to create parcel", err)
		return
	}
	SuccessResponse(w, http.StatusCreated, parcel)
}

func (s *server) newParcel(w http.ResponseWriter, r *http.Request) {
	var data model.Parcel

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		ErrUnprocessableEntityResponse(w, "Decode Error", err)
		return
	}

	if err := data.ValidateParcelInput(); err != nil {
		ErrInvalidEntityResponse(w, "Invalid Input", err)
		return
	}

	if err := s.parcelService.CreateParcel(r.Context(), data); err != nil {
		if errors.Is(err, model.ErrInvalid) {
			ErrInvalidEntityResponse(w, "invalid parcel", err)
			return
		}
		log.Error().Err(err).Msgf("[parcel] failed to create parcel Error: %v", err)
		ErrInternalServerResponse(w, "failed to create parcel", err)
		return
	}
	SuccessResponse(w, http.StatusCreated, "successful")
}

func (s *server) addCarrierRequest(w http.ResponseWriter, r *http.Request) {
	var data model.CarrierRequest
	vars := mux.Vars(r)

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		ErrUnprocessableEntityResponse(w, "Decode Error", err)
		return
	}
	parcelID, err := strconv.Atoi(vars["id"])
	if err != nil {
		ErrInvalidEntityResponse(w, "Invalid Parcel ID", err)
		return
	}
	data.ParcelID = parcelID

	if err := data.ValidateCarrierId(); err != nil {
		ErrInvalidEntityResponse(w, "Invalid Input", err)
		return
	}

	if err := s.carrierService.NewCarrierRequest(r.Context(), data); err != nil {
		if errors.Is(err, model.ErrInvalid) {
			ErrInvalidEntityResponse(w, "invalid Request", err)
			return
		}
		log.Error().Err(err).Msgf("[addCarrierRequest] failed to add new carrier request: %v", err)
		ErrInternalServerResponse(w, "failed to add new carrier request", err)
		return
	}
	SuccessResponse(w, http.StatusCreated, "Success")
}

func (s *server) getParcel(w http.ResponseWriter, r *http.Request) {
	var data model.Parcel

	vars := mux.Vars(r)
	parcelID, err := strconv.Atoi(vars["id"])

	if err != nil {
		ErrInvalidEntityResponse(w, "Invalid Parcel ID", err)
		return
	}

	data.ID = parcelID

	parcel, err := s.parcelService.GetParcelByID(r.Context(), data.ID)

	if err != nil {
		if errors.Is(err, model.ErrInvalid) || errors.Is(err, model.ErrNotFound) {
			ErrInvalidEntityResponse(w, "This ID does not exist.", err)
			return
		}
		log.Error().Err(err).Msgf("[getParcel] failed to parcel '%d': %v", data.ID, err)
		ErrInternalServerResponse(w, "Failed to fetch parcel "+strconv.Itoa(data.ID), err)
		return
	}

	SuccessResponse(w, http.StatusOK, parcel)
}
