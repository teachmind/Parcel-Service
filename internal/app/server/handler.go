package server

import (
	"net/http"
)

func (s *server) getParcelList(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	offset := r.URL.Query().Get("status")
	limit := r.URL.Query().Get("status")
	//fmt.Printf("%+v", idUser)

}
