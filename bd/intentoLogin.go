package bd

import (
	"github.com/jarrieta31/twittor/models"

	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el chequeo de login a la BD*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	//verfica si existe el usuario, returna un usuario, un boleano y el ID
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	// Si no lo encuentra se termina aca
	if encontrado == false {
		return usu, false
	}
	/*creo una slice de bytes para guardar la password sin encriptar que ingresó el usuario*/
	passwordBytes := []byte(password)
	/*Creo una slice de bytes para guardar la password encryptada que me vino de la base de datos*/
	passwordBD := []byte(usu.Password)
	/*Ahora para comparar la pasword encyptada con la que no está encriptada debemos usar una
	función de bcrypt que recibe primaro la encriptada y luego la que está sin encriptar.
	Si falla es porque no coinciden */
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
