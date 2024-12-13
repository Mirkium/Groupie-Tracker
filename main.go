package main

import (
	Route "CocktailMan/route"
	"net/http"
)

func main() {

	Route.Handle()

	fileServer := http.FileServer(http.Dir("./assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	http.ListenAndServe("localhost:8080", nil)
}
