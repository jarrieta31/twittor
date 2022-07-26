package routers

import (
	"net/http"

	"github.com/jarrieta31/twittor/bd"
	"github.com/jarrieta31/twittor/models"
)

/*BajaRelacion realiza el borrado de la relación en usuarios*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {

	//obtenemos el ID de la url
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario  //id de usuario actual
	t.UsuarioRelacionID = ID //ID del usuario al que dejamos de seguir

	status, err := bd.BorroRelacion(t)

	//Su hubo un erro con la BD
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	//Si no se puedo eliminar la relación
	if status == false {
		http.Error(w, "No se ha logrado borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
