package middlew

import (
	"net/http"

	"github.com/FlavioFBS/tuiterClone_back/bd"
)

// ChequeoBD ...
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			statusCode := 500
			http.Error(w, "Conexion perdida con la Base de Datos", statusCode)
			return
		}
		next.ServeHTTP(w, r)
	}
}
