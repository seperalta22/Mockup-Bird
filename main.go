package main

import (
	"log"

	"github.com/SergioPeralta22/Mockup-Bird/bd"
	"github.com/SergioPeralta22/Mockup-Bird/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}
	handlers.Manejadores()
}
