package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/jarrieta31/twittor/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPerfil busca un perfil en la BBD*/
func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	/*La variable objID almacena el ID transformado a hexadecimal dado que mongo
	necesita el id en ese formato*/
	objID, _ := primitive.ObjectIDFromHex(ID)

	/*En condicion declaramos justamente la condicion de busqueda*/
	condicion := bson.M{
		"_id": objID,
	}

	//Hacemos la consulta y guardamos el usuario obtendio en perifl
	err := col.FindOne(ctx, condicion).Decode(&perfil)
	/*Antes devolver el perfil ocultamos la password que tiene el usuario para no mostrarla */
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return perfil, err
	}
	//Si todo salió bien
	return perfil, nil
}
