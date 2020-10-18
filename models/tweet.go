package models

// Tweet captura el mensaje que llega del Body
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
