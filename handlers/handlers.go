package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores setear puerto, el handler y servidor se pone a escuchar
func Manejadores () {
	router := mux.NewRouter()

	// variable de entorno
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}

	// creando handler
	handler := cors.AllowAll().Handler(router)
	// escuchar en puerto y con el manejador que establecio un router
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}