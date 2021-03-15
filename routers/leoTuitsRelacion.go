package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FlavioFBS/tuiterClone_back/bd"
)

// LeoTuitsSeguidores lee los tweets de los seguidores
func LeoTuitsSeguidores(w http.ResponseWriter, r *http.Request) {
	var pagina_url string = r.URL.Query().Get("pagina")
	if len(pagina_url) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(pagina_url)

	if err != nil || pagina == 0 {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}
	respuesta, correcto := bd.LeoTuitSeguidores(IDUsuario, pagina)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
