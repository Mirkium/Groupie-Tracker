package controller

type Profil struct {
	Name        string     `json:"name"`
	Mail        string     `json:"mail"`
	Password    string     `json:"password"`
	Age         int        `json:"age"`
	LikeProduit []Cocktail `json:"like_produit"`
}

type User struct {
	Name   string
	Mdp    string
	Majeur bool
}

type Cocktail struct {
	Drinks []DrinkDetails `json:"drinks"`
}

type DrinkDetails struct {
	StrDrink      string `json:"strDrink"`
	StrDrinkThumb string `json:"strDrinkThumb"`
	ID_Drink      string `json:"idDrink"`
}

type Drink struct {
	Name string
	Img  string
	Id   string
}

type Category struct {
	TypeDrink []CategoryList `json:"drinks"`
}

type CategoryList struct {
	StrCategory string `json:"strCategory"`
}

type ReturnCategory struct {
	Name string
}

type Glass struct {
	Verre []GlassDetail `json:"drinks"`
}

type GlassDetail struct {
	StrGlass string `json:"Glass"`
}

type ReturnGlass struct {
	Name string
}
