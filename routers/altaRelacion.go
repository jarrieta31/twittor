package routers

import (
	"net/http"

	"github.com/jarrieta31/twittor/bd"
	"github.com/jarrieta31/twittor/models"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {

	//Capturamos de la url el parámetro id y hacemos el chequeo
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario  // id del usuario actaul
	t.UsuarioRelacionID = ID // id del usuario que vamos a seguir y viene por la url

	status, err := bd.InsertoRelacion(t)
	//Si hubo un problema con la BD
	if err != nil {
		http.Error(w, "Ocurrió un error al insertar relación!. "+err.Error(), http.StatusBadRequest)
		return
	}
	//Si no se puedo insertar el retistro
	if status == false {
		http.Error(w, "No se ha logrado insertar la relación!. "+err.Error(), http.StatusBadRequest)
		return
	}

	//todo salió bien
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
