package controllers

import (
	"GroupieTracker/models"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var ValueError models.Error

//====================================================Profil page============================================================

func ProfilePage(w http.ResponseWriter, r *http.Request) {

}

//==========================================================================================================================

//====================================================Login page============================================================

func LoginPage(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/login.html")
	if err != nil {
		fmt.Println("Il y une erreur de type : ", Red, err, Reset)
	}

	temp.Execute(w, nil)
}

//==========================================================================================================================

//===================================================Se connecter===========================================================

func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println(Red, "Erreur d'input", Reset)
		return
	}

	CheckValueName, _ := regexp.MatchString("[a-zA-Z-]{1,64}$", r.FormValue("name"))
	CheckValuePassword, _ := regexp.MatchString("[a-zA-Z-]{1,64}$", r.FormValue("password"))

	if (!CheckValueName && !CheckValuePassword) || !CheckValueName || !CheckValuePassword {
		ValueError.IfError = true
		ValueError.Error = "Your name or password are misspelled!"
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if Connect(r.FormValue("name"), r.FormValue("password")) {
		Menu.Profil.IsConnect = true
		fmt.Println("Connecter")
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
}

//==========================================================================================================================

//===================================================Register page==========================================================

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/register.html")
	if err != nil {
		fmt.Println("Il y une erreur de type : ", Red, err, Reset)
	}

	temp.Execute(w, nil)
}

//==========================================================================================================================

//===================================================S'enregistrer==========================================================

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println(Red, "Erreur d'input", Reset)
		return
	}
	CheckValueName, _ := regexp.MatchString("[a-zA-Z-]{1,64}$", r.FormValue("name"))

	if !CheckValueName {
		ValueError.IfError = true
		ValueError.Error = "Your name are misspelled!"
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	} else {
		ValueError.IfError = false

		Save(r.FormValue("name"), r.FormValue("password"))
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}

//==========================================================================================================================

//=======================================================Like===============================================================

func Like(w http.ResponseWriter, r *http.Request) {
	if Menu.Profil.IsConnect {
		Name := r.URL.Query().Get("name")
		if Name == "" {
			http.Error(w, "Missing Name parameter", http.StatusBadRequest)
			return
		}
		fmt.Println(Name, "Like")
		drink := RecupName(Name)

		for verif := range Menu.Cocktail {
			if Menu.Cocktail[verif].NameCocktail == drink.StrDrink {
				Menu.Profil.Like = append(Menu.Profil.Like, models.CocktailReturn{
					NameCocktail: Menu.Cocktail[verif].NameCocktail,
					ImgCocktail:  Menu.Cocktail[verif].ImgCocktail,
					IdCocktail:   Menu.Cocktail[verif].IdCocktail,
					Alcool:       Menu.Cocktail[verif].Alcool,
					Like:         true,
				})
				Menu.Cocktail[verif].Like = true
			}
		}
		AddCocktail()
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

//==========================================================================================================================

//====================================================Un Like===============================================================

func UnLike(w http.ResponseWriter, r *http.Request) {
	if Menu.Profil.IsConnect {
		Name := r.URL.Query().Get("name")
		if Name == "" {
			http.Error(w, "Missing Name parameter", http.StatusBadRequest)
			return
		}
		fmt.Println(Name, "UnLike")
		drink := RecupName(Name)

		for verif := range Menu.Cocktail {
			if Menu.Cocktail[verif].NameCocktail == drink.StrDrink {
				Menu.Cocktail[verif].Like = false
				RemoveCocktail(drink.StrDrink)
			}
		}
		for verif := range Menu.CocktailAlcool {
			if Menu.CocktailAlcool[verif].NameCocktail == drink.StrDrink {
				Menu.CocktailAlcool[verif].Like = false
				RemoveCocktail(drink.StrDrink)
			}
		}
		for verif := range Menu.NonCocktailAlcool {
			if Menu.NonCocktailAlcool[verif].NameCocktail == drink.StrDrink {
				Menu.NonCocktailAlcool[verif].Like = false
				RemoveCocktail(drink.StrDrink)
			}
		}
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

//==========================================================================================================================

func RemoveCocktail(nameCocktail string) {
	for drink := range Menu.Profil.Like {
		if Menu.Profil.Like[drink].NameCocktail == nameCocktail {
			Menu.Profil.Like = append(Menu.Profil.Like[:drink], Menu.Profil.Like[drink+1:]...)
		}
	}

	AddCocktail()
}

func AddCocktail() {
	newUser := models.User{
		Name:     Menu.Profil.Name,
		Password: Menu.Profil.Password,
		Like:     []models.CocktailLike{},
	}

	for drink := range Menu.Profil.Like {
		newUser.Like = append(newUser.Like, models.CocktailLike{
			Name:        Menu.Profil.Like[drink].NameCocktail,
			ImgCocktail: Menu.Profil.Like[drink].ImgCocktail,
			IdCocktail:  Menu.Profil.Like[drink].IdCocktail,
			Alcool:      Menu.Profil.Like[drink].Alcool,
		})
	}

	// Lire les utilisateurs existants à partir du fichier User.json
	var users []models.User
	file, err := os.OpenFile("./controllers/users.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Errorf("erreur lors de l'ouverture du fichier: %w", err)
	}
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Errorf("erreur lors de la lecture du fichier: %w", err)
	}

	if len(fileContent) > 0 {
		err = json.Unmarshal(fileContent, &users)
		if err != nil {
			fmt.Errorf("erreur lors de la désérialisation: %w", err)
		}
	}

	// Ajouter le nouvel utilisateur à la liste
	users = append(users, newUser)
	for L := range users {
		if users[L].Name == newUser.Name {
			users = append(users[:L], users[L+1:]...)
			break
		}
	}

	// Sérialiser les utilisateurs en JSON
	updatedContent, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Errorf("erreur lors de la sérialisation: %w", err)
	}

	// Réécrire le fichier avec les nouvelles données
	err = ioutil.WriteFile("./controllers/users.json", updatedContent, 0644)
	if err != nil {
		fmt.Errorf("erreur lors de l'écriture dans le fichier: %w", err)
	}
}
