package routers

import (
	"encoding/json"
	"net/http"

	"github.com/FlavioFBS/tuiterClone_back/bd"
)

// VerPerfil permite extraer los valores del perfil
func VerPerfil(w http.ResponseWriter, r *http.Request) {

	// parametros del req.params:
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametros ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "No se encontro el registro. "+err.Error(), 400)
		return
	}

	// indicar que se enviara un json
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// devolviendo los datos
	json.NewEncoder(w).Encode(perfil)
}
