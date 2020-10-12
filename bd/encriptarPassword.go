package bd

import "golang.org/x/crypto/bcrypt"

// EncriptarPassword ...
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	var sliceBytes = []byte(pass)                                // convirtiendo el string en bytes usando slice
	bytes, err := bcrypt.GenerateFromPassword(sliceBytes, costo) //paramatros: el texto a encriptar en formato de bytes y el costo
	return string(bytes), err                                    // retornando la respuesta de la funcion de encriptacion convertida en string
}
