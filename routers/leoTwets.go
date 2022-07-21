package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jarrieta31/twittor/bd"
)

/*LeoTweets leo los tweets*/
func LeoTweets(w http.ResponseWriter, r *http.Request) {

	//extraigo el id de la url
	ID := r.URL.Query().Get("id")

	//Si el id viene vació
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	//otra manera de verificar es el siguiente
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}

	//
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "Debe enviar el parámetro página, con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	//parceamos el valor de pagina que es un int comun y necesitamos que sea int64
	pag := int64(pagina)
	respuesta, correcto := bd.LeoTweets(ID, pag)

	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
