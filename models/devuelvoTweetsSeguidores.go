package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DevuelvoTweetsSeguidores es la structura con la que devolveremos los tweets*/
type DevuelvoTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID         string             `bson:"usuarioid" json:"userId,omitempty"`
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"userRelacionId,omitempty"`
	Tweet             struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}