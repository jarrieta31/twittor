package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jarrieta31/twittor/bd"
	"github.com/jarrieta31/twittor/models"
)

/*Email valor de Email usado en todos los EndPoints*/
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints*/
var IDUsuario string

/*ProcesoToken proceso el token para extraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{} //claim tiene que ser un puntero

	/* divie el texto del token y crea un array con dos elementos,
	la posición 0 es la palabra Bearer, y la posicion 1 es el token
	*/
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		// Importante: los mensajes de error no pueden tener mayusculas ni signos de puntuación.
		return claims, false, string(""), errors.New("formato de token invalida")
	}

	//aquí obtengo el token como string de la posición 1
	tk = strings.TrimSpace(splitToken[1])

	/*
		*En este paso se decoficia el token y se obtiene un objeto json con los datos contenidos
		en el, chequenado que el token es valido.
	*/
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	/* Una vez tenemos los datos debemos ver si el usuario existe*/
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	// si hubo error en el token
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err

}
