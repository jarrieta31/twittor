package routers

import (
	"encoding/json"

	"net/http"

	"time"

	"github.com/jarrieta31/twittor/bd"

	"github.com/jarrieta31/twittor/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	//capturamos el mensaje del body y lo decodificamos en la variable mensaje
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el Tweet", 400)
		return
	}

	//Si todo sale bien
	w.WriteHeader(http.StatusCreated)
}
