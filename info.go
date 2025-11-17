package quiz

import "fmt"

var InfoPoints int

func Info() {
	InfoPoints = 0
	fmt.Println("Dans une architecture MVC, à quoi correspond le modèle (Model) ?")
	fmt.Println("A) À la logique métier et aux données")
	fmt.Println("B) À l’interface utilisateur")
	fmt.Println("C) À la gestion des requêtes HTTP")
	fmt.Println("D) À la configuration du serveur")
	var a0 string
	fmt.Println("Votre réponse: ")
	fmt.Scanln(&a0)
	switch a0 {
	case "A", "a":
		fmt.Println("Bonne réponse !")
		InfoPoints += 1
	case "B", "b", "C", "c", "D", "d":
		fmt.Println("Mauvaise réponse")
	default:
		fmt.Println("Réponse incorrecte, veuillez entrer A, B, C ou D.")
	}

	fmt.Println("Quelle technologie est côté client dans une application full stack classique ?")
	fmt.Println("A) Node.js")
	fmt.Println("B) MySQL")
	fmt.Println("C) JavaScript (dans le navigateur)")
	fmt.Println("D) Express.js")
	var a1 string
	fmt.Println("Votre réponse: ")
	fmt.Scanln(&a1)

	switch a1 {
	case "C", "c":
		fmt.Println("Bonne réponse !")
		InfoPoints += 1
	case "A", "a", "B", "b", "D", "d":
		fmt.Println("Mauvaise réponse")
	default:
		fmt.Println("Réponse incorrecte, veuillez entrer A, B, C ou D.")
	}

	fmt.Println("Quel est le rôle principal de Git dans le développement ?")
	fmt.Println("A) Héberger une base de données")
	fmt.Println("B) Gérer le versionnement du code")
	fmt.Println("C) Exécuter du code JavaScript côté serveur")
	fmt.Println("D) Styliser une page web")
	var a2 string
	fmt.Println("Votre réponse: ")
	fmt.Scanln(&a2)

	switch a2 {
	case "B", "b":
		fmt.Println("Bonne réponse !")
		InfoPoints += 1
	case "A", "a", "C", "c", "D", "d":
		fmt.Println("Mauvaise réponse")
	default:
		fmt.Println("Réponse incorrecte, veuillez entrer A, B, C ou D.")
	}

	fmt.Println("Dans une requête HTTP, quelle méthode est utilisée pour récupérer des données sans les modifier ?")
	fmt.Println("A) POST")
	fmt.Println("B) PUT")
	fmt.Println("C) GET")
	fmt.Println("D) DELETE")
	var a3 string
	fmt.Println("Votre réponse: ")
	fmt.Scanln(&a3)

	switch a3 {
	case "C", "c":
		fmt.Println("Bonne réponse !")
		InfoPoints += 1
	case "A", "a", "B", "b", "D", "d":
		fmt.Println("Mauvaise réponse")
	default:
		fmt.Println("Réponse incorrecte, veuillez entrer A, B, C ou D.")
	}

	fmt.Println("Quel format est le plus couramment utilisé pour échanger des données entre un serveur et un client ?")
	fmt.Println("A) XML")
	fmt.Println("B) CSV")
	fmt.Println("C) JSON")
	fmt.Println("D) YAML")
	var a4 string
	fmt.Println("Votre réponse: ")
	fmt.Scanln(&a4)

	switch a4 {
	case "C", "c":
		fmt.Println("Bonne réponse !")
		InfoPoints += 1
	case "A", "a", "B", "b", "D", "d":
		fmt.Println("Mauvaise réponse")
	default:
		fmt.Println("Réponse incorrecte, veuillez entrer A, B, C ou D.")
	}

	_ = a0
	_ = a1
	_ = a2
	_ = a3
	_ = a4

	fmt.Println("Vous avez scoré ", InfoPoints, "/5")
	choix_quiz()
}
