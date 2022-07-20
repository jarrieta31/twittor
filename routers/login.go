package routers

/*Endpoint para el login de Usuario	*/

import (
	"encoding/json"

	"net/http"
	"time"

	"github.com/jarrieta31/twittor/bd"
	"github.com/jarrieta31/twittor/jwt"
	"github.com/jarrieta31/twittor/models"
)

/*Login realiza el login*/
func Login(w http.ResponseWriter, r *http.Request) {
	/* Seteamos el Header para decirle al navegador que el contenido que vamos a devolver es de
	tipo json */
	w.Header().Add("content-type", "application/json")
	// declaramos una variable t de tipo usuario.
	var t models.Usuario
	/*Procesamos el Body del Request para ver lo que tiene adentro, si trae el email y la Password
	del usuario lo guardamos en t. */
	err := json.NewDecoder(r.Body).Decode(&t)
	//chequeamos si no hubo error
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña inválidos! Error: "+err.Error(), 400)
		return
	}
	//chequeamos que el email no venga vacío
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido ", 400)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	/* La función IntentoLogin retorna false en existe si no encuntra al usuario en la BD*/
	if existe == false {
		http.Error(w, "Usuario y/o Contraseña inválidos.", 400)
		return
	}
	/* Si logró loguearse construimos un token en formato string con los datos del usuario */
	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el Token correspondiente "+err.Error(), 400)
		return
	}

	//modelo para devolver el token al navegador
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	//Seteamos el Header para application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //setea un status 200 o 201
	//asigna el valor de modelo del token a la respueta
	json.NewEncoder(w).Encode(resp)

	// Ejemplo para crear una cookie, esto no lo vamos a usar, es para saber como se hace.
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
