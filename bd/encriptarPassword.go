package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es la rutina que me permite encriptar la password*/
func EncriptarPassword(pass string) (string, error) {
	/*El costo el agoritmo lo va a elevar al cuadrado y esa es la
	cantidad de veces o pasadas que encripta la password
	*/
	costo := 8 //Encripta la password 2 a la 8 veces

	//esta función retorna un slice de bytes
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	//para retornarlo es necesario convertirlo a string
	return string(bytes), err
}
