package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es la rutina que me permite encriptar la password*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8 //Encripta la password 2 a la 8 veces
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
