package server

import (
	"fmt"
	"net/http"
	"strconv"

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
