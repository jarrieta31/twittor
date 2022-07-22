package routers

/* ProcesoToken es llamado por casi todos los EndPoints*/

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
	//clave necesaria para desencriptar el token
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")

	/*La variable claims es un modelo para almacenar el contenido del token*/
	claims := &models.Claim{} //claim tiene que ser un puntero

	/* divide el texto del token y crea un array con dos elementos,
	la posición 0 es la palabra Bearer, y la posicion 1 es el token, luego chequeamos
	que el largo este correcto.
	*/
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		// Importante: los mensajes de error no pueden tener mayusculas ni signos de puntuación.
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	/*Aquí obtengo el token como string de la posición 1, quito los espacios y sustituyo
	el valor de la variable tk */
	tk = strings.TrimSpace(splitToken[1])
	//fmt.Println(tk)

	/*En este paso se decoficia el token y se obtiene un objeto json con los datos contenidos
	en el, chequenado que el token es valido. */
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

	// Si hubo cualquier otro error
	return claims, false, string(""), err
}
