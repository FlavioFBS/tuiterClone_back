package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/FlavioFBS/tuiterClone_back/middlew"
	"github.com/FlavioFBS/tuiterClone_back/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores setear puerto, el handler y servidor se pone a escuchar
func Manejadores() {
	router := mux.NewRouter()

	// rutas
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil)).Methods("GET")
	

	// variable de entorno
	PORT := os.Getenv("PORT_TUITER")
	if PORT == "" {
		PORT = "8081"
	}

	// creando handler
	handler := cors.AllowAll().Handler(router)
	// escuchar en puerto y con el manejador que establecio un router
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
