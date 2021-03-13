package main

import (
	"github.com/FlavioFBS/tuiterClone_back/bd"
	"github.com/FlavioFBS/tuiterClone_back/handlers"
	"log"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
