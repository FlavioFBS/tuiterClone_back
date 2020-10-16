package bd

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/FlavioFBS/tuiterClone_back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BuscoPerfil busca un perfio en la BD
func BuscoPerfil(ID string) (models.Usuario, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database(os.Getenv("DB_MONGO"))
	col := db.Collection("usuarios")
	var perfil models.Usuario

	// convertir ID(string) a objectID:
	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{
		"_id": objID,
	}

	// buscar y guardar en variable perfil
	err := col.FindOne(ctx, condicion).Decode(&perfil)

	// limpiar el valor
	perfil.Password = ""

	if err != nil {
		log.Fatal("Registro no encontrado: " + err.Error())
		return perfil, err
	}
	return perfil, nil
}
