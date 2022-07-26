package routers

import (
	"net/http"

	"encoding/json"

	"github.com/jarrieta31/twittor/models"

	"github.com/jarrieta31/twittor/bd"
)

/*ConsultaRelacion chequea si hay relación entre 2 usuarios*/
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(t)

	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
