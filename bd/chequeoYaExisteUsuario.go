package bd

import (
	"context"
	"time"

	"github.com/FlavioFBS/tuiterClone_back/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ChequeoYaExisteUsuario verifica si el email ya se registró en la BD
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// dbName := os.Getenv("DB_MONGO")
	db := MongoCN.Database("TuiterClone") // selecionar db
	col := db.Collection("usuarios")      // seleccionar coleccion

	condicion := bson.M{"email": email}
	var resultado models.Usuario
	err := col.FindOne(ctx, condicion).Decode(&resultado) // buscar en la BD, codificar a json y pasarlo a resultado
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
