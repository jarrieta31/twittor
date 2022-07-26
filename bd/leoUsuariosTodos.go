package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/jarrieta31/twittor/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoUsuariosTodos Lee los usuarios registrados en el sistema, si se recibe "R" en quienes */
func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoCN.Database("twittor")
	col := bd.Collection("usuarios")

	var results []*models.Usuario

	/*Creamos la definicion de los parámetros para realizar la busqueda.
	Nota: es muy importate primero setear SetSkip y depues SetLimit*/
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	/*Nota: Cuando la consulta no es un findOne lo que devuelve es un cursor*/
	cur, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var encontrado, incluir bool

	/*Hay recorrer con el cursor todos los resultados*/
	for cur.Next(ctx) {
		//Cada resultado hay que decodificarlo y lo guardamos en la variable s
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		/*Ahora chequeamos la relación y para ello vamos guardando en r */
		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false

		encontrado, err = ConsultoRelacion(r)
		/*filtra los usuarios que no seguimos */
		if tipo == "new" && encontrado == false {
			incluir = true
		}
		/*filtra los usuarios que sí seguimos */
		if tipo == "follow" && encontrado == true {
			incluir = true
		}
		/*filtra que no sea nuestro primero usuario */
		if r.UsuarioRelacionID == ID {
			incluir = false
		}
		//Limpiamos los datos, para no mostrar información susceptible
		if incluir == true {
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			//Lo agrega a la lista de resultados
			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	//Si todo sale bien
	cur.Close(ctx)
	return results, true
}
