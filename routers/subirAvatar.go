package routers

import (
	"fmt"
	"github.com/FlavioFBS/tuiterClone_back/bd"
	"github.com/FlavioFBS/tuiterClone_back/models"
	"io"
	"net/http"
	"os"
	"strings"
)

// SubirAvatar funcion para subir imagen
func SubirAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatars")
	var arrayPoint = strings.Split(handler.Filename, ".")
	var extension = arrayPoint[len(arrayPoint)-1]
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

	fmt.Println("antes de abrir archivo......")
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666) // guardando almacenamiento en servidor
	if err != nil {
		http.Error(w, "Error al subir la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("antes de copiar......")
	_, err = io.Copy(f, file) // copiar el archivo recibo en la direccion reservarda

	if err != nil {
		http.Error(w, "Error al copiar la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	// registrar avatars en la BD:
	var usuario models.Usuario
	var status bool
	usuario.Avatar = IDUsuario + "." + extension
	fmt.Println("antes de modificar registro......")
	status, err = bd.ModificoRegistro(usuario, IDUsuario)

	if err != nil || status == false {
		http.Error(w, "Error al grabar el avatars en la BD! "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("antes de set-header......")
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
