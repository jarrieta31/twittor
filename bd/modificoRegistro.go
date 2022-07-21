package bd

import (
	"context"
	"time"

	"github.com/jarrieta31/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModificoRegistro permite modificar el perfil del usuario*/
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	registro := make(map[string]interface{})
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	registro["fechaNacimiento"] = u.FechaNacimiento
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}

	updtString := bson.M{
		"$set": registro,
	}

	//convertimos el ID recibido de string a ObjectIDFromHex
	objID, _ := primitive.ObjectIDFromHex(ID)

	/*Ahora crearmos un filtro para utilizar como condición para la actualización.
	El $eq significa que el campo _id de la BD sea igual al ID recibido*/
	filtro := bson.M{
		"_id": bson.M{"$eq": objID},
	}

	_, err := col.UpdateOne(ctx, filtro, updtString)

	//si hubo un error al en la actualización
	if err != nil {
		return false, err
	}

	//si estuvo todo bien
	return true, nil

}
