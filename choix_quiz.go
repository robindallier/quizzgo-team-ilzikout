package quiz

import "fmt"

func choix_quiz(){
	fmt.Println("Quel quiz voulez-vous faire ?")
	fmt.Println("A. Quiz Cybersécurité")
	fmt.Println("B. Quiz Informatique")
	fmt.Println("C. Quiz cyber")
	fmt.Println("D. Retour")
	var r string
	fmt.Println("Votre réponse: ")
	fmt.Scanln(&r)
	switch r {
		case "A", "a":
			Cyber()
		case "B", "b": 
			Info()
		case "C", "c": 
			IAData()
		case "D", "d":
			Menu()
		default:
			fmt.Println("Réponse incorrecte, ressaisissez votre réponse correctement.")
			choix_quiz()

	}

}

