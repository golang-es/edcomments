package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-es/edcomments/commons"
	"github.com/golang-es/edcomments/migration"
	"github.com/golang-es/edcomments/routes"
	"github.com/urfave/negroni"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
	flag.IntVar(&commons.Port, "port", 8080, "Puerto para el servidor web")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzó la migración...")
		migration.Migrate()
		log.Println("Finalizó la migración.")
	}

	// Inicia las rutas
	router := routes.InitRoutes()

	// Inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// Inicia el servidor
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", commons.Port),
		Handler: n,
	}

	log.Printf("Iniciado en http://localhost:%d", commons.Port)
	log.Println(server.ListenAndServe())
	log.Println("Finalizó la ejecución programa")
}
