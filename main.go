package main

import (
	"log"

	"github.com/jarrieta31/twittor/bd"
	"github.com/jarrieta31/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD.")
		return
	}
	handlers.Manejadores()
}
