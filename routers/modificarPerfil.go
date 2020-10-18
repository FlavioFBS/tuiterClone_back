package routers

import (
	"encoding/json"
	"net/http"

	"github.com/FlavioFBS/tuiterClone_back/bd"
	"github.com/FlavioFBS/tuiterClone_back/models"
)

// ModificarPerfil modificar el perfil de un usuario
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos: "+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar actualizar el registro. "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se logró modificar el registro. "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
