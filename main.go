package main

import (
	"log"

	"github.com/converse98/Grm-01/bd"
	"github.com/converse98/Grm-01/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
