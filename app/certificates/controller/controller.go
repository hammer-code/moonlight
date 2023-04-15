package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hammer-code/moonlight/domain/certificates"
)

type Controller struct {
	service certificates.CertificateService
}

func NewController(service certificates.CertificateService) *Controller {
	return &Controller{
		service: service,
	}
}

func (c Controller) GetCertificatByIdAndEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// event := vars["event"]
	id := vars["id"]

	// if event == "" {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write([]byte("err: event nil"))
	// 	return
	// }

	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("err: id nil"))
		return
	}

	cer, err := c.service.GetByExternalIDAndEvent(r.Context(), id)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("retrn err"))
		return
	}

	dByte, _ := json.Marshal(&cer)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dByte)
}
