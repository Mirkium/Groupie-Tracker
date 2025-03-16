package models

type Error struct {
	IfError bool
	Error   string
}

type Cocktail struct {
	IDDrink       string `json:"idDrink"`
	StrDrink      string `json:"strDrink"`
	StrDrinkThumb string `json:"strDrinkThumb"`
}

type CocktailList struct {
	Drinks []Cocktail `json:"drinks"`
}

type CocktailDetail struct {
	Drinks []CocktailDetailData `json:"drinks"`
}

type CocktailDetailData struct {
	StrDrink        string   `json:"strDrink"`
	StrInstructions string   `json:"strInstructions"`
	StrDrinkThumb   string   `json:"strDrinkThumb"`
}
