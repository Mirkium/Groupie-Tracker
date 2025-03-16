package models

type Profil struct {
	Name      string
	Password  string
	Like      []CocktailReturn
	IsConnect bool
}

type User struct {
	Name     string         `json:"name"`
	Password string         `json:"password`
	Like     []CocktailLike `json:"like`
}

type CocktailLike struct {
	Name        string `json:"name_cocktail"`
	ImgCocktail string `json:"img_cocktail"`
	IdCocktail  string `json:"id_cocktail"`
	Alcool      bool   `json:"isAlcool`
}

type Menu struct {
	Profil            Profil
	Cocktail          []CocktailReturn
	CocktailAlcool    []CocktailReturn
	NonCocktailAlcool []CocktailReturn
}

type CocktailReturn struct {
	NameCocktail string
	ImgCocktail  string
	IdCocktail   string
	Alcool       bool
	Like         bool
}

type CreateCocktail struct {
	Name        string
	Img         string
	Instruction string
}
