package bd

import (
	"context"
	"os"
	"time"

	"github.com/FlavioFBS/tuiterClone_back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertoTweet registro de tuit en la BD
func InsertoTweet(t models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // cancelar el contexto al finalizar

	db := MongoCN.Database(os.Getenv("DB_MONGO")) // selecionar db
	col := db.Collection("tweet")                 // seleccionar coleccion

	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	// insertando
	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID) // extraer el objID
	return objID.String(), true, nil
}
