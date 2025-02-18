package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var ConnectUser User

func Save(name string, mail string, mdp string) error {

	newUser := Profil{
		Name:        name,
		Mail:        mail,
		Password:    mdp,
		LikeProduit: []Cocktail{},
	}

	var users []Profil
	file, err := os.OpenFile("profil_data.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture du fichier: %w", err)
	}
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture du fichier: %w", err)
	}

	if len(fileContent) > 0 {
		err = json.Unmarshal(fileContent, &users)
		if err != nil {
			return fmt.Errorf("erreur lors de la désérialisation: %w", err)
		}
	}

	// Ajouter le nouvel utilisateur à la liste
	users = append(users, newUser)

	// Sérialiser les utilisateurs en JSON
	updatedContent, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur lors de la sérialisation: %w", err)
	}

	// Réécrire le fichier avec les nouvelles données
	err = ioutil.WriteFile("User.json", updatedContent, 0644)
	if err != nil {
		return fmt.Errorf("erreur lors de l'écriture dans le fichier: %w", err)
	}

	return nil
}

func Connect(name string, mdp string) bool {
	var users []Profil
	file, err := os.Open("User.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
		return false
	}
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return false
	}

	err = json.Unmarshal(fileContent, &users)
	if err != nil {
		return false
	}

	// Parcourir la liste des utilisateurs
	for _, user := range users {
		if user.Name == name && user.Password == mdp {
			ConnectUser.Name = name
			ConnectUser.Mdp = mdp

			return true
		}
	}

	return false
}
