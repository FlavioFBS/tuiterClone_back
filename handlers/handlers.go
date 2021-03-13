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
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")

	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")


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
