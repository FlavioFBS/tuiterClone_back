package bd

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/FlavioFBS/tuiterClone_back/models"
)

// ModificoRegistro permite modificar el perfil de usuario
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // cancelar el contexto al finalizar

	db := MongoCN.Database(os.Getenv("DB_MONGO")) // selecionar db
	col := db.Collection("usuarios")              // seleccionar coleccion

	// creando mapa con clave-string, valor-interface
	registro := make(map[string]interface{})

	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}
	registro["fechaNacimiento"] = u.FechaNacimiento
	if len(u.Avatar) > 0 {
		registro["avatars"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banners"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}
	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}

	updateString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{"$eq": objID}
	filtro := bson.M{
		"_id": condicion,
	}

	// instruccion de mongo
	_, err := col.UpdateOne(ctx, filtro, updateString)

	if err != nil {
		return false, err
	}
	return true, nil
}
