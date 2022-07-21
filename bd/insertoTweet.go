package bd

import (
	"context"
	"time"

	"github.com/jarrieta31/twittor/models"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoTweet graba el tweet en la BD*/
func InsertoTweet(t models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	/*Para guardar el registro tenemos que convertirlo al formato bson*/
	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)
	//si dió error
	if err != nil {
		return "", false, err
	}

	/*Obtiene el ID en formato hexadecimal del último campo insertado*/
	objID, _ := result.InsertedID.(primitive.ObjectID)

	//si todod salió bien
	return objID.String(), true, nil

}
