package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"parcel-service/internal/app/model"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func (s *server) getParcelList(w http.ResponseWriter, r *http.Request) {
	status, err := strconv.Atoi(r.URL.Query().Get("status"))
	if err != nil {
		ErrInvalidEntityResponse(w, "Invalid status value", err)
		return
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		ErrInvalidEntityResponse(w, "Invalid offset value", err)
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		ErrInvalidEntityResponse(w, "Invalid limit value", err)
		return
	}

	parcels, err := s.parcelService.GetParcels(r.Context(), status, limit, offset)
	if err != nil {
		if errors.Is(err, model.ErrInvalid) || errors.Is(err, model.ErrNotFound) {
			ErrInvalidEntityResponse(w, "No data exist for these query parmas", err)
			return
		}
		log.Error().Err(err).Msgf("[getParcelList] failed to get parcels for '%d', '%d', '%d': %v", status, limit, offset, err)
		ErrInternalServerResponse(w, "Failed to fetch parcel list for given query params", err)
		return
	}
	SuccessResponse(w, http.StatusOK, parcels)
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

	parcel, err := s.parcelService.CreateParcel(r.Context(), data)

	if err != nil {
		if errors.Is(err, model.ErrInvalid) {
			ErrInvalidEntityResponse(w, "invalid parcel", err)
			return
		}
		log.Error().Err(err).Msgf("[parcel] failed to create parcel Error: %v", err)
		ErrInternalServerResponse(w, "failed to create parcel", err)
		return
	}

	SuccessResponse(w, http.StatusCreated, parcel)
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
		if errors.Is(err, model.ErrInvalid) {
			ErrInvalidEntityResponse(w, "This ID does not exist.", err)
			return
		}
		if errors.Is(err, model.ErrNotFound) {
			ErrNotFoundResponse(w, "This ID does not exist.", err)
			return
		}
		log.Error().Err(err).Msgf("[getParcel] failed to parcel '%d': %v", data.ID, err)
		ErrInternalServerResponse(w, "Failed to fetch parcel "+strconv.Itoa(data.ID), err)
		return
	}

	SuccessResponse(w, http.StatusOK, parcel)
}

func (s *server) editParcel(w http.ResponseWriter, r *http.Request) {
	var data model.Parcel
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

	data.ID = parcelID

	if err := s.parcelService.EditParcel(r.Context(), data); err != nil {
		if errors.Is(err, model.ErrInvalid) || errors.Is(err, model.ErrNotFound) {
			ErrInvalidEntityResponse(w, "invalid Request", err)
			return
		}

		log.Error().Err(err).Msgf("[parcel/{id}/request] failed to update parcel: %v", err)
		ErrInternalServerResponse(w, "failed to update parcel", err)
		return
	}

	SuccessResponse(w, http.StatusCreated, "Success")
}

func (s *server) parcelCarrierAccept(w http.ResponseWriter, r *http.Request) {
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

	// validating input credentials for parcel request
	if err := data.ValidateCarrierId(); err != nil {
		ErrInvalidEntityResponse(w, "Invalid Input", err)
		return
	}

	if err := s.carrierService.AssignCarrierToParcel(r.Context(), data); err != nil {
		log.Error().Err(err).Msgf("[parcelCarrierAccept] failed to assign carrier to parcel: %v", err)
		ErrInternalServerResponse(w, "failed to assign carrier to parcel", err)
		return
	}
	SuccessResponse(w, http.StatusNoContent, "Successful")
}
