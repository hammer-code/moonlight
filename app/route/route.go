package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hammer-code/moonlight/app/certificates/controller"
)

type route struct {
	certController *controller.Controller
}

func NewRoute(c *controller.Controller) *mux.Router {
	r := mux.NewRouter()
	v1 := r.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	v1.HandleFunc("/certificates/{id}", c.GetCertificatByIdAndEvent)

	return v1
}
