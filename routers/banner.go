package routers

import (
	"github.com/FlavioFBS/tuiterClone_back/bd"
	"github.com/FlavioFBS/tuiterClone_back/models"
	"io"
	"net/http"
	"os"
	"strings"
)

// SubirBanner function to upload image
func SubirBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banners")
	var arrayPoint = strings.Split(handler.Filename, ".")
	var extension = arrayPoint[len(arrayPoint)-1]
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666) // saving storage on server
	if err != nil {
		http.Error(w, "Error to upload image! "+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file) // copy file get in reserved address

	if err != nil {
		http.Error(w, "Error to copy image! "+err.Error(), http.StatusBadRequest)
		return
	}

	// registrar avatars en la BD:
	var usuario models.Usuario
	var status bool
	usuario.Banner = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)

	if err != nil || status == false {
		http.Error(w, "Error to record banner on BD! "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
