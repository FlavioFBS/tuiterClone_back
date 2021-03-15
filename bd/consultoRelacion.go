package bd

import (
	"context"
	"os"
	"time"
	"fmt"
	"github.com/FlavioFBS/tuiterClone_back/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ConsultoRelacion
func ConsultoRelacion (t models.Relacion) (bool, error) {
	ctx,cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DB_MONGO")) // selecionar db
	col := db.Collection("relacion")

	// condicion de busqueda
	condicion := bson.M{
		"usuarioid": t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}
	var resultado models.Relacion
	fmt.Println(resultado)
	// se busca, si hay resultado se va a resultado
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
