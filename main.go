package main

import (
	"log"
	"net/http"
	"GroupieTracker/routes"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	routes.InitRoutes()
	log.Println("Le serveur écoute sur le port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erreur du serveur : ", err)
	}
}
