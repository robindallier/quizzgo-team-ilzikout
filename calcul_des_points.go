package quiz

import "fmt"

func Points(){
    fmt.Println("Que voulez-vous voir ?")
    fmt.Println("A. Points en cybersécurité")
    fmt.Println("B. Points en informatique")
    fmt.Println("C. Points en IA et data")
    fmt.Println("D. Moyenne")
    fmt.Println("E. Retour")
    var r string
    fmt.Println("Votre réponse: ")
    fmt.Scanln(&r)
    switch r{
		case "A", "a":
			fmt.Println("Vos points sur le quiz en cybersécurité sont de ", CyberPoints, "/10.")
			Points()
		case "B", "b":
			fmt.Println("Vos points sur le quiz en informatique sont de ", InfoPoints, "/10.")
			Points()
		case "C", "c":
			fmt.Println("Vos points sur le quiz en IA et data sont de ", IADataPoints, "/10.")
			Points()
		case "D", "d":
			fmt.Println("Votre moyenne est de ", ((CyberPoints+InfoPoints+IADataPoints)/3), "/15.")
			Points()

		case "E", "e":
			Menu()

    }

}
