package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var CocktailAlcohol []Drink
var CocktailNonAlcohol []Drink

var TypeCategory []ReturnCategory

var TypeGlass []ReturnGlass

var DataCocktail Cocktail

//================================================================FILTER ALCOOL=============================================================

func FiltreAlcool() {

	UrlReq := "https://www.thecocktaildb.com/api/json/v1/1/filter.php?a=Alcoholic"

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodGet, UrlReq, nil)
	if errReq != nil {
		fmt.Println("Erreur lors de la création de la requête :", errReq)
		return
	}

	req.Header.Set("User-Agent", "CocktailMan")

	res, errRes := httpClient.Do(req)
	if errRes != nil {
		fmt.Println("Erreur lors de l'envoi de la requête :", errRes)
		return
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		var dataCocktail Cocktail
		errDecode := json.NewDecoder(res.Body).Decode(&dataCocktail)
		if errDecode != nil {
			fmt.Println("Erreur de décodage JSON :", errDecode)
			return
		}

		for _, drink := range dataCocktail.Drinks {
			CocktailAlcohol = append(CocktailAlcohol, Drink{
				Name: drink.StrDrink,
				Img:  drink.StrDrinkThumb,
				Id:   drink.ID_Drink,
			})
		}
	} else {
		fmt.Printf("Erreur HTTP - Code : %d, Message : %s\n", res.StatusCode, res.Status)
	}
}

//==========================================================================================================================================

//================================================================FILTER NON ALCOOL=========================================================

func Filter_NonAlcohol() {
	UrlReq := "https://www.thecocktaildb.com/api/json/v1/1/filter.php?a=Non_Alcoholic"

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodGet, UrlReq, nil)
	if errReq != nil {
		fmt.Println("Erreur lors de la création de la requête :", errReq)
		return
	}

	req.Header.Set("User-Agent", "CocktailMan")

	res, errRes := httpClient.Do(req)
	if errRes != nil {
		fmt.Println("Erreur lors de l'envoi de la requête :", errRes)
		return
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		var dataCocktail Cocktail
		errDecode := json.NewDecoder(res.Body).Decode(&dataCocktail)
		if errDecode != nil {
			fmt.Println("Erreur de décodage JSON :", errDecode)
			return
		}

		for _, drink := range dataCocktail.Drinks {
			CocktailNonAlcohol = append(CocktailNonAlcohol, Drink{
				Name: drink.StrDrink,
				Img:  drink.StrDrinkThumb,
			})
		}
	} else {
		fmt.Printf("Erreur HTTP - Code : %d, Message : %s\n", res.StatusCode, res.Status)
	}
}

//=======================================================================================================================================

func SearchCocktail(name string) {
	URLReq := "https://www.thecocktaildb.com/api/json/v1/1/search.php?s=" + name

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodGet, URLReq, nil)
	if errReq != nil {
		fmt.Println("Erreur lors de la création de la requête :", errReq)
		return
	}

	req.Header.Set("User-Agent", "CocktailMan")

	res, errRes := httpClient.Do(req)
	if errRes != nil {
		fmt.Println("Erreur lors de l'envoi de la requête :", errRes)
		return
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		errDecode := json.NewDecoder(res.Body).Decode(&DataCocktail)

		if errDecode != nil {
			fmt.Println("Erreur de décodage JSON :", errDecode)
			return
		}
		fmt.Print(DataCocktail.Drinks[0].StrDrink)
		return
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
		return
	}
}
