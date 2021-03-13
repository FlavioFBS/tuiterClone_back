package bd

import (
	"context"
	"os"
	"time"
	"github.com/FlavioFBS/tuiterClone_back/models"
)

// BorroRelacion
func BorroRelacion (t models.Relacion) (bool, error) {
	ctx,cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DB_MONGO")) // selecionar db
	col := db.Collection("relacion")

	_,err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
