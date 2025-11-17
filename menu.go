package quiz

import "fmt"
import "os"



func Menu(){
	fmt.Println("Que voulez-vous faire ?")
	fmt.Println("A. Faire un quiz")
	fmt.Println("B. Voir mes points")
	fmt.Println("C. Quitter")
	fmt.Println("Votre réponse: ")
	var r string
	fmt.Scanln(&r)
	switch r{
		case "A", "a":
			choix_quiz()

		case "B", "b":
			Points()

		case "C","c":
    		os.Exit(0)
		default:
			fmt.Println("Erreur dans votre réponse, veuillez la ressaisir.")
			Menu()
	}

}

