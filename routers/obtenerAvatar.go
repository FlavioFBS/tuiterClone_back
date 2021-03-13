package routers

import (
	"github.com/FlavioFBS/tuiterClone_back/bd"
	"io"
	"net/http"
	"os"
)

// ObtenerAvatar send avatars to HTTP
func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	// find perfil by id
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatars/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}
	// enviar el archivo binario al response
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}

}
