package quiz

import "fmt"

var IADataPoints int

func IAData() {

	IADataPoints = 0

	fmt.Println("Dans un problème de classification, quel est l’objectif principal ?")
	fmt.Println("A) Prédire une valeur numérique")
	fmt.Println("B) Regrouper des données similaires")
	fmt.Println("C) Assigner une étiquette à une donnée")
	fmt.Println("D) Réduire la dimension des données")
	var b0 string
	fmt.Println("Votre réponse: ")
	fmt.Scanln(&b0)

	switch b0 {
	case "C", "c":
		fmt.Println("Bonne réponse !")
		IADataPoints += 1
	case "A", "a", "B", "b", "D", "d":
		fmt.Println("Mauvaise réponse")
	default:
		fmt.Println("Réponse incorrecte, veuillez entrer A, B, C ou D.")
	}

	fmt.Println("Quelle bibliothèque Python est la plus utilisée pour le machine learning ?")
	fmt.Println("A) Pandas")
	fmt.Println("B) NumPy")
	fmt.Println("C) Scikit-learn")
	fmt.Println("D) Matplotlib")
	var b1 string
	fmt.Println("Votre réponse: ")
	fmt.Scanln(&b1)

	switch b1 {
	case "C", "c":
		fmt.Println("Bonne réponse !")
		IADataPoints += 1
	case "A", "a", "B", "b", "D", "d":
		fmt.Println("Mauvaise réponse")
	default:
		fmt.Println("Réponse incorrecte, veuillez entrer A, B, C ou D.")
	}

	fmt.Println("Quel type de données est le plus adapté pour entraîner un modèle de réseau de neurones ?")
	fmt.Println("A) Données textuelles non structurées")
	fmt.Println("B) Données numériques normalisées")
	fmt.Println("C) Données catégorielles brutes")
	fmt.Println("D) Données manquantes non traitées")
	var b2 string
	fmt.Println("Votre réponse: ")
	fmt.Scanln(&b2)

	switch b2 {
	case "B", "b":
		fmt.Println("Bonne réponse !")
		IADataPoints += 1
	case "A", "a", "C", "c", "D", "d":
		fmt.Println("Mauvaise réponse")
	default:
		fmt.Println("Réponse incorrecte, veuillez entrer A, B, C ou D.")
	}

	fmt.Println("Qu’est-ce que le surapprentissage (overfitting) ?")
	fmt.Println("A) Le modèle est trop simple et ne capture pas les patterns")
	fmt.Println("B) Le modèle apprend le bruit des données d’entraînement")
	fmt.Println("C) Le modèle généralise mal sur de nouvelles données")
	fmt.Println("D) Le modèle est trop lent à s’entraîner")
	var b3 string
	fmt.Println("Votre réponse: ")
	fmt.Scanln(&b3)

	switch b3 {
	case "B", "b":
		fmt.Println("Bonne réponse !")
		IADataPoints += 1
	case "A", "a", "C", "c", "D", "d":
		fmt.Println("Mauvaise réponse")
	default:
		fmt.Println("Réponse incorrecte, veuillez entrer A, B, C ou D.")
	}

	fmt.Println("Quelle étape est essentielle avant d’entraîner un modèle ?")
	fmt.Println("A) Visualiser les résultats")
	fmt.Println("B) Nettoyer et prétraiter les données")
	fmt.Println("C) Déployer le modèle en production")
	fmt.Println("D) Choisir une couleur pour les graphiques")
	var b4 string
	fmt.Println("Votre réponse: ")
	fmt.Scanln(&b4)

	switch b4 {
	case "B", "b":
		fmt.Println("Bonne réponse !")
		IADataPoints += 1
	case "A", "a", "C", "c", "D", "d":
		fmt.Println("Mauvaise réponse")
	default:
		fmt.Println("Réponse incorrecte, veuillez entrer A, B, C ou D.")
	}

	_ = b0
	_ = b1
	_ = b2
	_ = b3
	_ = b4

	fmt.Println("Vous avez scoré ", IADataPoints, "/5")
	choix_quiz()
}
