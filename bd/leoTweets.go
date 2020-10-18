package bd

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/FlavioFBS/tuiterClone_back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// LeoTweets devuelve in slice de tipo DevuelvoTweet y un boolean
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // cancelar el contexto al finalizar

	db := MongoCN.Database(os.Getenv("DB_MONGO")) // selecionar db
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	var limite int64
	limite = 20
	opciones := options.Find()
	opciones.SetLimit(limite)                           // cantidad por pagina
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) // ordenado por fecha descendente(-1)
	opciones.SetSkip((pagina - 1) * limite)             // saltos de registros por el empaginado

	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	// recorrer cada documento del cursos:
	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}
