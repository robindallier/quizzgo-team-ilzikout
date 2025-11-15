package quiz

import "fmt"


func IAData() {

	IADataPoints := 0

	fmt.Println("Dans un problème de classification, quel est l’objectif principal ?")
	fmt.Println("A) Prédire une valeur numérique")
	fmt.Println("B) Regrouper des données similaires")
	fmt.Println("C) Assigner une étiquette à une donnée")
	fmt.Println("D) Réduire la dimension des données")
	var b0 string
	fmt.Print("Votre réponse: ")
	fmt.Scanln(&b0)

	if b0 == "C"{
		fmt.Print("Bonne réponse !")
		IADataPoints+=1
	}else{
		fmt.Println("Mauvaise réponse")
	}


	fmt.Println("Quelle bibliothèque Python est la plus utilisée pour le machine learning ?")
	fmt.Println("A) Pandas")
	fmt.Println("B) NumPy")
	fmt.Println("C) Scikit-learn")
	fmt.Println("D) Matplotlib")
	var b1 string
	fmt.Print("Votre réponse: ")
	fmt.Scanln(&b1)

	if b1 == "C"{
		fmt.Print("Bonne réponse !")
		IADataPoints+=1
	}else{
		fmt.Println("Mauvaise réponse")
	}


	fmt.Println("Quel type de données est le plus adapté pour entraîner un modèle de réseau de neurones ?")
	fmt.Println("A) Données textuelles non structurées")
	fmt.Println("B) Données numériques normalisées")
	fmt.Println("C) Données catégorielles brutes")
	fmt.Println("D) Données manquantes non traitées")
	var b2 string
	fmt.Print("Votre réponse: ")
	fmt.Scanln(&b2)

	if b2 == "B"{
		fmt.Print("Bonne réponse !")
		IADataPoints+=1
	}else{
		fmt.Println("Mauvaise réponse")
	}


	fmt.Println("Qu’est-ce que le surapprentissage (overfitting) ?")
	fmt.Println("A) Le modèle est trop simple et ne capture pas les patterns")
	fmt.Println("B) Le modèle apprend le bruit des données d’entraînement")
	fmt.Println("C) Le modèle généralise mal sur de nouvelles données")
	fmt.Println("D) Le modèle est trop lent à s’entraîner")
	var b3 string
	fmt.Print("Votre réponse: ")
	fmt.Scanln(&b3)

	if b3 == "B"{
		fmt.Print("Bonne réponse !")
		IADataPoints+=1
	}else{
		fmt.Println("Mauvaise réponse")
	}



	fmt.Println("Quelle étape est essentielle avant d’entraîner un modèle ?")
	fmt.Println("A) Visualiser les résultats")
	fmt.Println("B) Nettoyer et prétraiter les données")
	fmt.Println("C) Déployer le modèle en production")
	fmt.Println("D) Choisir une couleur pour les graphiques")
	var b4 string
	fmt.Print("Votre réponse: ")
	fmt.Scanln(&b4)

	if b4 == "B"{
		fmt.Print("Bonne réponse !")
		IADataPoints+=1
	}else{
		fmt.Println("Mauvaise réponse")
	}
	
  
	_ = b0
	_ = b1
	_ = b2
	_ = b3
	_ = b4

	fmt.Print("Vous avez scoré ", IADataPoints, "/5")
}