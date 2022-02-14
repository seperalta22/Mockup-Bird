package routers

import (
	"net/http"

	"github.com/SergioPeralta22/Mockup-Bird/bd"
)

/* EliminarTweet me permite eliminar un tweet determinado  */
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar borrar el tweet"+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
