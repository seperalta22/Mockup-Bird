package middlew

import (
	"net/http"

	"github.com/SergioPeralta22/Mockup-Bird/bd"
)

//* ChequeoBD es el middlew que me permite conocer el estado de la BD.
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
