package middlew

import (
	"net/http"

	"github.com/jarrieta31/twittor/routers"
)

/*ValidoJWT permite validar el JWT que nos viene en la peticion */
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*Obtenemos el token recibido desde el request y lo validamos. Para ello
		leemos la variable 'Authorization' que viene en e Header y la validamos
		con la funcion ProcesoToken*/
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token ! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}

}
