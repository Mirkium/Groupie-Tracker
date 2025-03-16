package controllers

import (
	"GroupieTracker/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

var apiBaseURL = "https://www.thecocktaildb.com/api/json/v1/1/"

//============================================ Récupérer tous les cocktails (alcoolisé ou non) ========================================

func getAllCocktails() []models.CocktailReturn {

	var listDefault []models.CocktailReturn

	alcool := FiltreAlcool()
	nonAlcool := Filter_NonAlcohol()

	for k := 0; k < len(alcool); k++ {
		listDefault = append(listDefault, alcool[k])
	}
	for L := 0; L < len(nonAlcool); L++ {
		listDefault = append(listDefault, nonAlcool[L])
	}
	return listDefault
}

//=====================================================================================================================================

//============================================ Récupérer tous les cocktails alcoolisés ================================================

func FiltreAlcool() []models.CocktailReturn {

	var listReturn []models.CocktailReturn

	UrlReq := "https://www.thecocktaildb.com/api/json/v1/1/filter.php?a=Alcoholic"

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodGet, UrlReq, nil)
	if errReq != nil {
		fmt.Println("Erreur lors de la création de la requête :", errReq)
	}
	req.Header.Set("User-Agent", "CocktailMan")

	res, errRes := httpClient.Do(req)
	if errRes != nil {
		fmt.Println("Erreur lors de l'envoi de la requête :", errRes)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		var dataCocktail models.CocktailList
		errDecode := json.NewDecoder(res.Body).Decode(&dataCocktail)
		if errDecode != nil {
			fmt.Println("Erreur de décodage JSON :", errDecode)
		}
		for _, drink := range dataCocktail.Drinks {
			listReturn = append(listReturn, models.CocktailReturn{
				NameCocktail: drink.StrDrink,
				ImgCocktail:  drink.StrDrinkThumb,
				IdCocktail:   drink.IDDrink,
				Alcool:       true,
				Like:         false,
			})
		}
	} else {
		fmt.Printf("Erreur HTTP - Code : %d, Message : %s\n", res.StatusCode, res.Status)
	}

	return listReturn
}

//=====================================================================================================================================

//============================================ Récupérer tous les cocktails non alcoolisés ============================================

func Filter_NonAlcohol() []models.CocktailReturn {

	var listReturn []models.CocktailReturn

	UrlReq := "https://www.thecocktaildb.com/api/json/v1/1/filter.php?a=Non_Alcoholic"

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodGet, UrlReq, nil)
	if errReq != nil {
		fmt.Println("Erreur lors de la création de la requête :", Red, errReq, Reset)
	}

	req.Header.Set("User-Agent", "CocktailMan")

	res, errRes := httpClient.Do(req)
	if errRes != nil {
		fmt.Println("Erreur lors de l'envoi de la requête :", Red, errRes, Reset)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		var dataCocktail models.CocktailList
		errDecode := json.NewDecoder(res.Body).Decode(&dataCocktail)
		if errDecode != nil {
			fmt.Println("Erreur de décodage JSON :", Red, errDecode, Reset)
		}
		for _, drink := range dataCocktail.Drinks {
			listReturn = append(listReturn, models.CocktailReturn{
				NameCocktail: drink.StrDrink,
				ImgCocktail:  drink.StrDrinkThumb,
				IdCocktail:   drink.IDDrink,
				Alcool:       false,
				Like:         false,
			})
		}
	} else {
		fmt.Printf("Erreur HTTP - Code : %d, Message : %s\n", res.StatusCode, res.Status)
	}
	return listReturn
}

//=====================================================================================================================================

//====================================== Récupérer un cocktail spécifique à partir de son nom ==========================================

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println(Red, "Erreur d'input", Reset)
		return
	}

	CheckValueName, _ := regexp.MatchString("[a-zA-Z-]{1,64}$", r.FormValue("search"))

	if !CheckValueName {
		ValueError.IfError = true
		ValueError.Error = "Your name are misspelled!"
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	} else {
		ValueError.IfError = false

		url := "https://www.thecocktaildb.com/api/json/v1/1/search.php?s=" + r.FormValue("search")

		response, _ := http.Get(url)
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)

		var cocktailDetail models.CocktailDetail
		json.Unmarshal(body, &cocktailDetail)

		if len(cocktailDetail.Drinks) == 0 {
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
		Verre := models.CreateCocktail{
			Name:        cocktailDetail.Drinks[0].StrDrink,
			Instruction: cocktailDetail.Drinks[0].StrInstructions,
			Img:         cocktailDetail.Drinks[0].StrDrinkThumb,
		}

		http.Redirect(w, r, "/cocktail?name="+Verre.Name, http.StatusFound)
	}

}

//=====================================================================================================================================

// Fonction générique pour récupérer les cocktails
func getCocktailsFromAPI(url string) ([]models.Cocktail, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var cocktailList models.CocktailList
	if err := json.Unmarshal(body, &cocktailList); err != nil {
		return nil, err
	}

	return cocktailList.Drinks, nil
}

//==============================================================================================================

//====================================Récupe un cocktail avec son nom===========================================

func RecupName(name string) models.Cocktail {
	URLReq := "https://www.thecocktaildb.com/api/json/v1/1/search.php?s=" + name

	var retourCocktail models.Cocktail

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodGet, URLReq, nil)
	if errReq != nil {
		fmt.Println("Erreur lors de la création de la requête :", errReq)
	}

	req.Header.Set("User-Agent", "CocktailMan")

	res, errRes := httpClient.Do(req)
	if errRes != nil {
		fmt.Println("Erreur lors de l'envoi de la requête :", errRes)
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		var DataCocktail models.CocktailList
		errDecode := json.NewDecoder(res.Body).Decode(&DataCocktail)

		if errDecode != nil {
			fmt.Println("Erreur de décodage JSON :", errDecode)
		}
		fmt.Print(DataCocktail.Drinks[0].StrDrink)
		retourCocktail.IDDrink = DataCocktail.Drinks[0].IDDrink
		retourCocktail.StrDrink = DataCocktail.Drinks[0].StrDrink
		retourCocktail.StrDrinkThumb = DataCocktail.Drinks[0].StrDrinkThumb
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}
	return retourCocktail
}

//==============================================================================================================
