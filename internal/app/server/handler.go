package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"parcel-service/internal/app/model"

	"github.com/rs/zerolog/log"
)

func (s *server) createParcel(w http.ResponseWriter, r *http.Request) {
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
