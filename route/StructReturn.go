package route

type ReturnMenu struct {
	ListCocktailAlcool     []CocktailReturn
	ListCocktailNonAlcool  []CocktailReturn
	ListCarrouselAlcool    []string
	ListCarrouselNonAlcool []string
}

type CocktailReturn struct {
	Cocktail_Name string
	Cocktail_img  string
	Cocktail_Id   string
	Like bool
}
