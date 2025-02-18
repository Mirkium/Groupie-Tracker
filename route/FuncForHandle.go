package route

import (
	"fmt"
	"net/http"
	"os"
)

func List_img() ([]string, []string) {
	Folder_Alcool := "./assets/img/Alcool"
	Folder_NonAlcool := "./assets/img/Non_Alcool"

	var ListReturn []string
	var ListReturnAlcool []string

	AlcoolDos, err := os.ReadDir(Folder_Alcool)
	NonAlcoolDos, err := os.ReadDir(Folder_NonAlcool)

	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du dossier: %v", err)
	}

	for _, element := range AlcoolDos {
		Fichier := element.Name()
		ListReturnAlcool = append(ListReturnAlcool, Fichier)
	}

	for _, element := range NonAlcoolDos {
		Fichier := element.Name()
		ListReturn = append(ListReturn, Fichier)
	}

	return ListReturnAlcool, ListReturn
}

//Fonction qui récupère l'id du cocktail

func RecupName(w http.ResponseWriter, r *http.Request) {
	Name := r.URL.Query().Get("name")
	if Name == "" {
		http.Error(w, "Missing Name parameter", http.StatusBadRequest)
		return
	}

	Cocktail.Cocktail_Name = Name
}
