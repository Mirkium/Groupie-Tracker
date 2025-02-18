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

var Cocktail CocktailReturn

func Handle() {

	//==========================================================ACCUEIL=======================================================================

	http.HandleFunc("/groupie_tracker/cocktail_man/accueil", func(w http.ResponseWriter, r *http.Request) {
		var Menu ReturnMenu
		temp, err := template.ParseFiles("./templates/Accueil.html")
		if err != nil {
			fmt.Println("Tu as un problème de type : ", err)
		}
		controller.FiltreAlcool()
		controller.Filter_NonAlcohol()

		Menu.ListCarrouselAlcool, Menu.ListCarrouselNonAlcool = List_img()

		for _, K := range controller.CocktailAlcohol {
			Menu.ListCocktail = append(Menu.ListCocktail, CocktailReturn{
				Cocktail_Name: K.Name,
				Cocktail_img:  K.Img,
				Cocktail_Id:   K.Id,
			})
		}
		temp.Execute(w, Menu)

	})

	//========================================================================================================================================

	//==========================================================ACCUEIL FILTER================================================================

	http.HandleFunc("/groupie_tracker/cocktail_man/accueil/filter_alcohol", func(w http.ResponseWriter, r *http.Request) {
		var Menu ReturnMenu
		temp, err := template.ParseFiles("./templates/Accueil.html")
		if err != nil {
			fmt.Println("Tu as un problème de type : ", err)
		}
		controller.FiltreAlcool()

		Menu.ListCarrouselAlcool, Menu.ListCarrouselNonAlcool = List_img()
		if !Filter_NonAlcohol && !FilterAlcool {
			for _, Element := range controller.CocktailAlcohol {
				Menu.ListCocktail = append(Menu.ListCocktail, CocktailReturn{Element.Name, Element.Img, Element.Id})
			}
			for _, Element := range controller.CocktailNonAlcohol {
				Menu.ListCocktail = append(Menu.ListCocktail, CocktailReturn{Element.Name, Element.Img, Element.Id})
			}

			temp.Execute(w, Menu)
		} else if !Filter_NonAlcohol && FilterAlcool {

			for _, Element := range controller.CocktailAlcohol {
				Menu.ListCocktail = append(Menu.ListCocktail, CocktailReturn{Element.Name, Element.Img, Element.Id})
			}
			temp.Execute(w, Menu)
		} else if Filter_NonAlcohol && Filter_NonAlcohol {
			for _, Element := range controller.CocktailNonAlcohol {
				Menu.ListCocktail = append(Menu.ListCocktail, CocktailReturn{Element.Name, Element.Img, Element.Id})
			}

			temp.Execute(w, Menu)
		} else {
			temp.Execute(w, nil)
		}
	})

	//========================================================================================================================================

	//==========================================================ACCUEIL NON FILTER============================================================

	http.HandleFunc("/groupie_tracker/cocktail_man/accueil/!filter_alcohol", func(w http.ResponseWriter, r *http.Request) {
		var Menu ReturnMenu
		temp, err := template.ParseFiles("./templates/Accueil.html")
		if err != nil {
			fmt.Println("Tu as un problème de type : ", err)
		}

		controller.Filter_NonAlcohol()

		Menu.ListCarrouselAlcool, Menu.ListCarrouselNonAlcool = List_img()
		if !Filter_NonAlcohol && !FilterAlcool {
			for _, Element := range controller.CocktailAlcohol {
				Menu.ListCocktail = append(Menu.ListCocktail, CocktailReturn{Element.Name, Element.Img, Element.Id})
			}
			for _, Element := range controller.CocktailNonAlcohol {
				Menu.ListCocktail = append(Menu.ListCocktail, CocktailReturn{Element.Name, Element.Img, Element.Id})
			}

			temp.Execute(w, Menu)
		} else if !Filter_NonAlcohol && FilterAlcool {

			for _, Element := range controller.CocktailAlcohol {
				Menu.ListCocktail = append(Menu.ListCocktail, CocktailReturn{Element.Name, Element.Img, Element.Id})
			}
			temp.Execute(w, Menu)
		} else if Filter_NonAlcohol && Filter_NonAlcohol {
			for _, Element := range controller.CocktailNonAlcohol {
				Menu.ListCocktail = append(Menu.ListCocktail, CocktailReturn{Element.Name, Element.Img, Element.Id})
			}

			temp.Execute(w, Menu)
		} else {
			temp.Execute(w, nil)
		}
	})

	//========================================================================================================================================

	http.HandleFunc("/groupie_tracker/cocktail_man/connect", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("./templates/Connect.html")
		if err != nil {
			fmt.Println("Tu as une erreur de type :", err)
			return
		}

		temp.Execute(w, nil)
	})

	http.HandleFunc("/Connect", func(w http.ResponseWriter, r *http.Request) {

	})

	//=====================================================================================================================================

	//========================================================COCKTAIL=====================================================================

	http.HandleFunc("/groupie_tracker/cocktail_man/cocktail", func(w http.ResponseWriter, r *http.Request) {
		RecupName(w, r)
		controller.SearchCocktail(Cocktail.Cocktail_Name)
		fmt.Print(Cocktail.Cocktail_Name)
		Cocktail.Cocktail_img = controller.DataCocktail.Drinks[0].StrDrinkThumb

		temp, err := template.ParseFiles("./templates/Cocktail.html")
		if err != nil {
			fmt.Print("Tu as une erreur de type :", err)
			return
		}

		temp.Execute(w, Cocktail)
	})

	//======================================================================================================================================

	//========================================================REGISTER======================================================================

	http.HandleFunc("/groupie_tracker/cocktail_man/register", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("./templates/Register.html")
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
