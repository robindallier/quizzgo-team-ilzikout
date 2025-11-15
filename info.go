package quiz

import "fmt"


func Info() {
	InfoPoints := 0
	fmt.Println("Dans une architecture MVC, à quoi correspond le modèle (Model) ?")
	fmt.Println("A) À la logique métier et aux données")
	fmt.Println("B) À l’interface utilisateur")
	fmt.Println("C) À la gestion des requêtes HTTP")
	fmt.Println("D) À la configuration du serveur")
	var a0 string
	fmt.Print("Votre réponse: ")
	fmt.Scanln(&a0)
	if a0 == "A"{
		fmt.Print("Bonne réponse !")
		InfoPoints+=1
	}else{
		fmt.Println("Mauvaise réponse")
	}


	fmt.Println("Quelle technologie est côté client dans une application full stack classique ?")
	fmt.Println("A) Node.js")
	fmt.Println("B) MySQL")
	fmt.Println("C) JavaScript (dans le navigateur)")
	fmt.Println("D) Express.js")
	var a1 string
	fmt.Print("Votre réponse: ")
	fmt.Scanln(&a1)

	if a1 == "C"{
		fmt.Print("Bonne réponse !")
		InfoPoints+=1
	}else{
		fmt.Println("Mauvaise réponse")
	}

	fmt.Println("Quel est le rôle principal de Git dans le développement ?")
	fmt.Println("A) Héberger une base de données")
	fmt.Println("B) Gérer le versionnement du code")
	fmt.Println("C) Exécuter du code JavaScript côté serveur")
	fmt.Println("D) Styliser une page web")
	var a2 string
	fmt.Print("Votre réponse: ")
	fmt.Scanln(&a2)

	if a2 == "B"{
		fmt.Print("Bonne réponse !")
		InfoPoints+=1
	}else{
		fmt.Println("Mauvaise réponse")
	}

	fmt.Println("Dans une requête HTTP, quelle méthode est utilisée pour récupérer des données sans les modifier ?")
	fmt.Println("A) POST")
	fmt.Println("B) PUT")
	fmt.Println("C) GET")
	fmt.Println("D) DELETE")
	var a3 string
	fmt.Print("Votre réponse: ")
	fmt.Scanln(&a3)

	if a3 == "C"{
		fmt.Print("Bonne réponse !")
		InfoPoints+=1
	}else{
		fmt.Println("Mauvaise réponse")
	}

	fmt.Println("Quel format est le plus couramment utilisé pour échanger des données entre un serveur et un client ?")
	fmt.Println("A) XML")
	fmt.Println("B) CSV")
	fmt.Println("C) JSON")
	fmt.Println("D) YAML")
	var a4 string
	fmt.Print("Votre réponse: ")
	fmt.Scanln(&a4)

	if a4 == "C"{
		fmt.Print("Bonne réponse !")
		InfoPoints+=1
	}else{
		fmt.Println("Mauvaise réponse")
	}

	_ = a0
	_ = a1
	_ = a2
	_ = a3
	_ = a4

	fmt.Print("Vous avez scoré ", InfoPoints, "/5")
}