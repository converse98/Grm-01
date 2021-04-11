package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/converse98/Grm-01/middlew"
	"github.com/converse98/Grm-01/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Funciona como seteo del Puerto, el handler y pongo a escuchar al servidor */
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
