package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FlavioFBS/tuiterClone_back/bd"
	"github.com/FlavioFBS/tuiterClone_back/models"
)

// GraboTweet regitra el tweet en la BD
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar el tweet. "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se logró insertar el Tweet.", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
