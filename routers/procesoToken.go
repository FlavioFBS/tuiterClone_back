package routers

import (
	"errors"
	"os"
	"strings"

	"github.com/FlavioFBS/tuiterClone_back/bd"
	"github.com/FlavioFBS/tuiterClone_back/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Email valor de email usado en todos los endpoints
var Email string

// IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints
var IDUsuario string

// ProcesoToken proceso para extraer los valores del token
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte(os.Getenv("GO_PRIVATE_KEY"))
	claims := &models.Claim{}

	// en caso el token incluya el prefijo 'Bearer'
	tk = strings.TrimSpace(strings.ReplaceAll(tk, "Bearer", ""))

	// validando el token
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalido")
	}
	return claims, false, string(""), err
}
