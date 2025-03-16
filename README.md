/home	GET	Page d'accueil avec la liste des cocktails.
/like	POST	Ajoute un cocktail aux favoris d'un utilisateur.
/unlike	POST	Supprime un cocktail des favoris d'un utilisateur.
/search	POST	Recherche un cocktail par son nom.
/login	GET	Page de connexion utilisateur.
/register	GET	Page d'inscription utilisateur.
/signIn	POST	Authentifie un utilisateur.
/signUp	POST	Crée un nouveau compte utilisateur.
/profile	GET	Affiche le profil de l'utilisateur connecté.
/cocktail/{id}	GET	Affiche la page d'un cocktail spécifique.

📌 Décomposition du projet
Le projet a été divisé en plusieurs phases clés :

Planification : Définition des fonctionnalités principales (authentification, recherche, favoris).
Mise en place du backend : Création des routes et de la logique métier en Go.
Développement du frontend : Intégration des vues HTML/CSS et gestion de l’affichage.
Connexion frontend-backend : Tests des fonctionnalités et ajustements.
Optimisation et corrections : Débogage et amélioration de l’expérience utilisateur.
⏳ Gestion du temps et priorités
Pour maximiser l’efficacité, nous avons :

Priorisé les fonctionnalités essentielles (authentification, recherche, affichage des cocktails).
Adopté une approche itérative en testant chaque partie avant d’ajouter de nouvelles fonctionnalités.
Utilisé un suivi des tâches pour organiser le développement et éviter les blocages.
📚 Stratégie de documentation
Références officielles : Documentation Go, net/http, et bases de données.
Exemples et tutoriels : Recherche sur GitHub et Stack Overflow pour résoudre des problèmes spécifiques.
Tests et expérimentations : Implémentation progressive avec des tests réguliers pour valider les choix techniques.
