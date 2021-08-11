package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"parcel-service/internal/app/model"
	"github.com/gorilla/mux"
	"strconv"
)

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

	if err := s.carrierParcelAcceptService.AssignCarrierToParcel(r.Context(), data); err != nil {
		if errors.Is(err, model.ErrInvalid) {
		ErrInvalidEntityResponse(w, "invalid user", err)
		return
	}
		log.Error().Err(err).Msgf("[parcel/{id}/accept] failed to assign carrier to parcel: %v", err)
		ErrInternalServerResponse(w, "failed to assign carrier to parcel", err)
		return
	}
	SuccessResponse(w, http.StatusCreated, key)
}