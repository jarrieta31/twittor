package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/jarrieta31/twittor/models"
)

func GeneroJWT(t models.Usuario) (string, error) {

	//Jwt trabaja con un array de bytes y no con string
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")

	//Esta es la carga útil del token
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	/*Para crear el token se necesitan 2 parámetros, el algoritmo de encriptación y la carga util*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	//Por último es necesrio firmar el  token
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
