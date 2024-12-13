package controller

type Profil struct {
	Name        string     `json:"name"`
	Mail        string     `json:"mail"`
	Password    string     `json:"password"`
	LikeProduit []Cocktail `json:"like_produit"`
}

type User struct {
	Name    string
	Mail    string
	Connect bool
}

type Cocktail struct {
	Drinks []DrinkDetails `json:"drinks"`
}

type DrinkDetails struct {
	StrDrink      string `json:"strDrink"`
	StrDrinkThumb string `json:"strDrinkThumb"`
	IDDrink       string `json:"idDrink"`
}

type Drink struct {
	Name string
	Img  string
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
