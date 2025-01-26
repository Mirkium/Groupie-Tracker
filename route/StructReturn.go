package route

type ReturnMenu struct {
	ListCocktail             []CocktailReturn
	ListCarrouselAlcool      []string
	ListCarrouselNonAlcool   []string
}

type CocktailReturn struct {
	Cocktail_Name string
	Cocktail_img  string
	Cocktail_Id   string
}
