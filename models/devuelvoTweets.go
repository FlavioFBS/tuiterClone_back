package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DevuelvoTweets es la estructura con la que se devuelve los tweets
type DevuelvoTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userid" json:"userId,omitemty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitemty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitemty"`
}
