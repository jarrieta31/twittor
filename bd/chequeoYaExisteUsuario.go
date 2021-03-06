package bd

import (
	"context"
	"time"

	"github.com/jarrieta31/twittor/models"

	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un email de parámetro y chequea si ya está en la BD.*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	/*Al realizar la busqueda, si hay un error lo guarda en err, y si no, lo guarda en resultado.
	FindOne recibe el contexto y la condición de busqueda y Decode convierte la salida
	de la busqueda a josn y lo asigna a resultado.
	*/
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex() //convierte el ObjID a string
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
