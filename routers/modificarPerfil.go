package routers

import (
	"encoding/json"

	"net/http"

	"github.com/jarrieta31/twittor/bd"
	"github.com/jarrieta31/twittor/models"
)

/*ModificarPerfil modifica el perfil de usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	//modelo para guardar los datos que el usuario envió
	var t models.Usuario

	/*Obtenemos los datos que vienen en el body en formato json*/
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos Incorrectos, "+err.Error(), 400)
		return
	}

	var status bool
	/*Modificamos el usuario en la BD utilizando la variable global obtenida del token IDUsuario*/
	status, err = bd.ModificoRegistro(t, IDUsuario)

	//si hubo un error
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro. Reintente nuevamente. "+err.Error(), 400)
		return
	}
	// si no hubo error pero no modificó ningún registro porque no lo encontro
	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario.", 400)
		return
	}

	//si todo salió bien
	w.WriteHeader(http.StatusCreated)
}
