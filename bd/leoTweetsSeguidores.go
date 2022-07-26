package bd

import (
	"context"
	"time"

	"github.com/jarrieta31/twittor/models"

	"go.mongodb.org/mongo-driver/bson"
)

/*LeoTweetsSeguidores lee los tweets de mis seguidores*/
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	//variable para el salto de pagina de la paginación
	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)

	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	//ahora unimos las tablas usando lookup y necesita 4 parametros para poder hace eso
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid", //es el campo local por el que se unen las tablas
			"foreignField": "userid",            //es el campo en la tabla tweet
			"as":           "tweet",             // nos pide un alias
		}})

	/* El unwind nos sirve para que todos los documentos nos vengan iguales */
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	/* Condición para ordenar los datos con fecha más reciente primero (fecha descendente)*/
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	//condicion del salto para paginar
	condiciones = append(condiciones, bson.M{"$skip": skip})
	//condicion del límite para la paginación
	condiciones = append(condiciones, bson.M{"$limit": 20})

	// La función Aggregate nos devuelve un cursor que no es necesario recorrer
	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvoTweetsSeguidores
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	//todo salió bien
	return result, true

}
