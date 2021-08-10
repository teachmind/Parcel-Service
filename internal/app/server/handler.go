package server

import (
	"encoding/json"
	"net/http"
	"parcel-service/internal/app/model"


)

func (s *server) parcelCarrierAccept(w http.ResponseWriter, r *http.Request) {
	var data model.CarrierRequest
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		ErrUnprocessableEntityResponse(w, "Decode Error", err)
		return
	}
	SuccessResponse(w, http.StatusCreated, data)
}