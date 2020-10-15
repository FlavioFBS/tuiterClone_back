package jwt

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/FlavioFBS/tuiterClone_back/models"
)

// GeneroJWT genera el jwt para el login
func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte(os.Getenv("GO_PRIVATE_KEY"))

	// lista de privilegios que se grabaran en el payload
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	// agregar header y contenido al token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	//firmar token
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
