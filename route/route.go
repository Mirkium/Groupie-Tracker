package route

import (
	controller "CocktailMan/controller"
	"fmt"
	"html/template"
	"net/http"
)

const (
	Red   = "\033[91m"
	Reset = "\033[0m"
)

var Filter_NonAlcohol bool = false
var FilterCategory bool = false
var FilterAlcool bool = false
var FilterDrink bool = false
var FilterGlass bool = false

var Cocktail CocktailReturn

func Handle() {

	//==========================================================ACCUEIL=======================================================================

	http.HandleFunc("/groupie_tracker/home", func(w http.ResponseWriter, r *http.Request) {
		var Menu ReturnMenu
		temp, err := template.ParseFiles("./templates/Accueil.html")
		if err != nil {
			fmt.Println("Tu as une erreur de type : ", Red, err, Reset)
		}

		controller.FiltreAlcool()
		controller.Filter_NonAlcohol()

		Menu.ListCarrouselAlcool, Menu.ListCarrouselNonAlcool = List_img()

		for _, K := range controller.CocktailAlcohol {
			Menu.ListCocktailAlcool = append(Menu.ListCocktailAlcool, CocktailReturn{
				Cocktail_Name: K.Name,
				Cocktail_img:  K.Img,
				Cocktail_Id:   K.Id,
			})
		}
		for _, K := range controller.CocktailNonAlcohol {
			Menu.ListCocktailNonAlcool = append(Menu.ListCocktailNonAlcool, CocktailReturn{
				Cocktail_Name: K.Name,
				Cocktail_img:  K.Img,
				Cocktail_Id:   K.Id,
			})
		}

		temp.Execute(w, Menu)
	})

	//========================================================================================================================================

	http.HandleFunc("/groupie_tracker/cocktail_man/connect", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("./templates/Connect.html")
		if err != nil {
			fmt.Println("Tu as une erreur de type : ", Red, err, Reset)
			return
		}

		temp.Execute(w, nil)
	})

	http.HandleFunc("/Connect", func(w http.ResponseWriter, r *http.Request) {

	})

	//=====================================================================================================================================

	//========================================================COCKTAIL=====================================================================

	http.HandleFunc("/groupie_tracker/cocktail", func(w http.ResponseWriter, r *http.Request) {
		RecupName(w, r)
		controller.SearchCocktail(Cocktail.Cocktail_Name)
		fmt.Print(Cocktail.Cocktail_Name)
		Cocktail.Cocktail_img = controller.DataCocktail.Drinks[0].StrDrinkThumb

		temp, err := template.ParseFiles("./templates/Cocktail.html")
		if err != nil {
			fmt.Print("Tu as une erreur de type : ", Red, err, Reset)
			return
		}

		temp.Execute(w, Cocktail)
	})

	//======================================================================================================================================

	//========================================================REGISTER======================================================================

	http.HandleFunc("/groupie_tracker/cocktail_man/register", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("./templates/Register.html")
		if err != nil {
			fmt.Println("Tu as une erreur de type : ", Red, err, Reset)
			return
		}

		temp.Execute(w, nil)
	})

	http.HandleFunc("/AddUsser", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/groupie_tracker/cocktail_man/profil", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/not/found", func(w http.ResponseWriter, r *http.Request) {

	})
}
