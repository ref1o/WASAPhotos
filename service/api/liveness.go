package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// liveness is an HTTP handler that checks the API server status. If the server cannot serve requests (e.g., some
// resources are not ready), this should reply with HTTP Status 500. Otherwise, with HTTP Status 200
func (rt *_router) liveness(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Esempio di controllo di vitalità con il ping del database
	if err := rt.DB.Ping(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Se tutto è a posto, rispondiamo con uno stato OK (HTTP 200)
	w.WriteHeader(http.StatusOK)
}
