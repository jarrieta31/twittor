package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jarrieta31/twittor/bd"
)

/*VerPerfil permite extraer los valores del Perfil*/
func VerPerfil(w http.ResponseWriter, r *http.Request) {

	//obtemos el ID que viene por la url de la petición Get
	ID := r.URL.Query().Get("id")
	//Chequea que el id no venga vacío
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	// Si no encontro el perfil
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar el registro. "+err.Error(), 400)
		return
	}

	/*Si encontró el perfil, seteamos el header para avisar que lo que vamos a enviar es un json y
	enviamos el perfil encontrado*/
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	//codificamos el contenido del perfil a json
	json.NewEncoder(w).Encode(perfil)

}
