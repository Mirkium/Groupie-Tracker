package controllers

import (
	"GroupieTracker/models"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

const (
	Red   = "\033[91m"
	Reset = "\033[0m"
)

var apiURL = "https://www.thecocktaildb.com/api/json/v1/1/filter.php?a="

var Menu models.Menu
var Profil models.Profil

func HomePage(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Print("Error : ", Red, err, Reset)
		return
	}

	Menu.Cocktail = getAllCocktails()
	Menu.CocktailAlcool = FiltreAlcool()
	Menu.NonCocktailAlcool = Filter_NonAlcohol()

	cocktails := getAllCocktails()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des cocktails:", Red, err, Reset)
	}
	Menu.Cocktail = cocktails
	for drink := range Menu.Cocktail {
		for L := range Menu.Profil.Like {
			if Menu.Cocktail[drink].NameCocktail == Menu.Profil.Like[L].NameCocktail {
				Menu.Cocktail[drink].Like = true
			}
		}
	}
	for drink := range Menu.CocktailAlcool {
		for L := range Menu.Profil.Like {
			if Menu.CocktailAlcool[drink].NameCocktail == Menu.Profil.Like[L].NameCocktail {
				Menu.CocktailAlcool[drink].Like = true
			}
		}
	}
	for drink := range Menu.NonCocktailAlcool {
		for L := range Menu.Profil.Like {
			if Menu.NonCocktailAlcool[drink].NameCocktail == Menu.Profil.Like[L].NameCocktail {
				Menu.NonCocktailAlcool[drink].Like = true
			}
		}
	}

	temp.Execute(w, Menu)
}

func CocktailPage(w http.ResponseWriter, r *http.Request) {
	cocktailID := r.URL.Query().Get("name")
	url := "https://www.thecocktaildb.com/api/json/v1/1/search.php?s=" + cocktailID
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(cocktailID)

	var cocktailDetail models.CocktailDetail
	json.Unmarshal(body, &cocktailDetail)

	temp, err := template.ParseFiles("./templates/cocktail.html")
	if err != nil {
		fmt.Print("Error : ", Red, err, Reset)
		return
	}
	Verre := models.CreateCocktail{
		Name:        cocktailDetail.Drinks[0].StrDrink,
		Instruction: cocktailDetail.Drinks[0].StrInstructions,
		Img:         cocktailDetail.Drinks[0].StrDrinkThumb,
	}

	temp.Execute(w, Verre)
}

func getCocktails(alcohol string) []models.Cocktail {
	url := apiURL + alcohol
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var cocktailList models.CocktailList
	json.Unmarshal(body, &cocktailList)

	return cocktailList.Drinks
}
