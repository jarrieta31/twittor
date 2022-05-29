package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el modelo de usuario de la base de MongoDB*/
type Usuario struct {
	ID              primitive.ObjectID `bson: "_id,omitemty" json:"id"`
	Nombre          string             `bson: "nombre" json:"nombre, omitemty"`
	Apellido        string             `bson: "apellidos" json:"apellido, omitemty"`
	FechaNacimiento time.Time          `bson: "fechaNacimiento" json:"fechaNacimeinto, omitemty"`
	Email           string             `bson: "email" json:"email"`
	Password        string             `bson: "password" json:"password, omitemty"`
	Avatar          string             `bson: "avatar" json:"avatar, omitemty"`
	Banner          string             `bson: "banner" json:"banner, omitemty"`
	Biografia       string             `bson: "biografia" json:"biografia, omitemty"`
	Ubicacion       string             `bson: "ubicacion" json:"ubicacion, omitemty"`
	SitioWeb        string             `bson: "sitioWeb" json:"sitioWeb, omitemty"`
}
