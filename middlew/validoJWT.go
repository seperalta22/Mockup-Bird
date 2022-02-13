package middlew

import (
	"net/http"

	"github.com/SergioPeralta22/Mockup-Bird/routers"
)

func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoTokens(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token !"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
