package routes

import (
	"GroupieTracker/controllers"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/home", controllers.HomePage)
	http.HandleFunc("/like", controllers.Like)
	http.HandleFunc("/unlike", controllers.UnLike)
	http.HandleFunc("/search", controllers.Search)
	http.HandleFunc("/login", controllers.LoginPage)
	http.HandleFunc("/register", controllers.RegisterPage)
	http.HandleFunc("/signIn", controllers.SignIn)
	http.HandleFunc("/signUp", controllers.SignUp)
	http.HandleFunc("/profile", controllers.ProfilePage)
	http.HandleFunc("/cocktail/", controllers.CocktailPage)
}
