package routers

import (
	"encoding/json"
	"net/http"

	"github.com/FlavioFBS/tuiterClone_back/bd"
	"github.com/FlavioFBS/tuiterClone_back/models"
)

// Registro es la funcion para crear en la BD el registro del usuario
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	// guardar el dato del body de la peticion a la variable 't'
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos: "+err.Error(), 400)
		return
	}
	// validaciones
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario: "+err.Error(), 500)
		return
	}
	if status == false {
		http.Error(w, "No se a logrado registrar al usuario", 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
