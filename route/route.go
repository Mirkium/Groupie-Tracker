package route

import (
	controller "CocktailMan/controller"
	"fmt"
	"html/template"
	"net/http"
)



var Filter_NonAlcohol bool = false
var FilterCategory bool = false
var FilterAlcool bool = false
var FilterDrink bool = false
var FilterGlass bool = false

func Handle() {

	http.HandleFunc("/groupie_tracker/cocktail_man/accueil", func(w http.ResponseWriter, r *http.Request) {
		var Menu []ReturnMenu
		temp, err := template.ParseFiles("../templates/Accueil.html")
		if err != nil {
			fmt.Println("Tu as un problème de type : ", err)
		}
		controller.FiltreAlcool()
		controller.Filter_NonAlcohol()

		if !Filter_NonAlcohol && !FilterAlcool {
			for _, Element := range controller.CocktailAlcohol {
				Menu = append(Menu, ReturnMenu{Element.Name, Element.Img})
			}
			for _, Element := range controller.CocktailNonAlcohol {
				Menu = append(Menu, ReturnMenu{Element.Name, Element.Img})
			}

			temp.Execute(w, Menu)
		} else if !Filter_NonAlcohol && FilterAlcool{
			
			for _, Element := range controller.CocktailAlcohol {
				Menu = append(Menu, ReturnMenu{Element.Name, Element.Img})
			}
			temp.Execute(w, Menu)
		} else if Filter_NonAlcohol && Filter_NonAlcohol {
			for _, Element := range controller.CocktailNonAlcohol {
				Menu = append(Menu, ReturnMenu{Element.Name, Element.Img})
			}

			temp.Execute(w, Menu)
		} else {
			temp.Execute(w, nil)
		}

	})

	http.HandleFunc("/groupie_tracker/cocktail_man/connect", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("../templates/Connect.html")
		if err != nil {
			fmt.Println("Tu as une erreur de type :", err)
			return
		}

		temp.Execute(w, nil)
	})

	http.HandleFunc("/Connect", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/groupie_tracker/cocktail_man/coktail", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/groupie_tracker/cocktail_man/register", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("../templates/Register.html")
		if err != nil {
			fmt.Println("Tu as une erreur de type :", err)
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
