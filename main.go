package main

import (
	"log"
	"github.com/FlavioFBS/tuiterClone_back/handlers"
	"github.com/FlavioFBS/tuiterClone_back/bd"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
