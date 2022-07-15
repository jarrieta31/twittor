package routers

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
	w.Header().Add("content-type", "application/json")

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
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
	//chequeamos que el email no venga vacío
	if existe == false {
		http.Error(w, "Usuario y/o Contraseña inválidos.", 400)
		return
	}
	//Construimos un token
	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el Token correspondiente "+err.Error(), 400)
		return
	}

	//modelo para devolver al navegador
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //setea un status 200 o 201
	json.NewEncoder(w).Encode(resp)

	//Ejemplo para crear una cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
