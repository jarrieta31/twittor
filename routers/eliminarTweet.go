package routers

import (
	"net/http"

	"github.com/jarrieta31/twittor/bd"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	//obtengo el id de la url
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el tweet", http.StatusBadRequest)
		return

	}

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)

}
