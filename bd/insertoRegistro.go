package bd

import (
	"context"
	"os"
	"time"

	"github.com/FlavioFBS/tuiterClone_back/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertoRegistro es la parada final con la BD para insertar los datos del usuario
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	// creacion de contexto
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // cancelar el contexto al finalizar

	db := MongoCN.Database(os.Getenv("DB_MONGO")) // selecionar db
	col := db.Collection("usuarios")              // seleccionar coleccion

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
