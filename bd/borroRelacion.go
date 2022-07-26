package bd

import (
	"context"
	"time"

	"github.com/jarrieta31/twittor/models"
)

func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	//Ahora simplemente eliminamos la relacion de la BD
	_, err := col.DeleteOne(ctx, t)
	// si hubo un error
	if err != nil {
		return false, err
	}

	//todo salió bién
	return true, nil

}
