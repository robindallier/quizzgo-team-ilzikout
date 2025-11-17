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
			fmt.Println("Vos points sur le quiz en cybersécurité sont de", CyberPoints, "/5")
			Points()
		case "B", "b":
			fmt.Println("Vos points sur le quiz en informatique sont de", InfoPoints, "/5")
			Points()
		case "C", "c":
			fmt.Println("Vos points sur le quiz en IA et data sont de", IADataPoints, "/5")
			Points()
		case "D", "d":
			total := CyberPoints + InfoPoints + IADataPoints 
			avg := float64(total) / 3.0                   
			fmt.Printf("Votre moyenne est de %.2f/5 (total %d/15)\n", avg, total)
			Points()

		case "E", "e":
			Menu()

    }

}
