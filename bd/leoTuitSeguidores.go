package bd

import (
	"context"
	"github.com/FlavioFBS/tuiterClone_back/models"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"time"
)

// LeoTuitSeguidores lee los tuits de mis seguidores
func LeoTuitSeguidores(ID string, pagina int) ([]models.DevuelvoTuitsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DB_MONGO")) // selecionar db
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20
	condiciones := make([]bson.M, 0)
	// usando framework agregate de mongo)
	// $match para buscar usuarioId de la relacion
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet", // con qué tabla se quiere unir el campo relacion
			"localField":   "usuariorelcionid",
			"foreignField": "userid", // el campo de la tabla tweet donde está el idusuario
			"as":           "tweet",
		}})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})                // para que el json se entregue sin la estructura maestro-detalle
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}}) // ordenar por fecha descendente
	condiciones = append(condiciones, bson.M{"$skip": skip})                      // empaginado
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvoTuitsSeguidores // formatear slice con la estructura armada
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
