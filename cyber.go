package quiz

import "fmt"

var CyberPoints int

func Cyber() {
	CyberPoints = 0
	fmt.Println("Quel type d'attaque consiste à submerger un serveur de requêtes pour le rendre indisponible ?")
	fmt.Println("A) Phishing")
	fmt.Println("B) SQL Injection")
	fmt.Println("C) DDOS")
	fmt.Println("D) Man-in-the-Middle")

	var r0 string
	fmt.Print("Votre réponse: ")
	fmt.Scanln(&r0)

	switch r0 {
	case "C", "c":
		fmt.Println("Bonne réponse !")
		CyberPoints += 1
	case "A", "a", "B", "b", "D", "d":
		fmt.Println("Mauvaise réponse")
	default:
		fmt.Println("Réponse incorrecte, veuillez entrer A, B, C ou D.")
	}

	fmt.Println("Quel est le rôle d'un pare-feu (firewall) ?")
	fmt.Println("A) Chiffrer les données")
	fmt.Println("B) Filtrer le trafic réseau selon des règles")
	fmt.Println("C) Stocker les mots de passe")
	fmt.Println("D) Scanner les virus dans les e-mails")

	var r1 string
	fmt.Print("Votre réponse: ")
	fmt.Scanln(&r1)

	switch r1 {
	case "B", "b":
		fmt.Println("Bonne réponse !")
		CyberPoints += 1
	case "A", "a", "C", "c", "D", "d":
		fmt.Println("Mauvaise réponse")
	default:
		fmt.Println("Réponse incorrecte, veuillez entrer A, B, C ou D.")
	}

	fmt.Println("Quel principe de sécurité recommande d'accorder le minimum de privilèges nécessaires ?")
	fmt.Println("A) Défense en profondeur")
	fmt.Println("B) Principe du moindre privilège")
	fmt.Println("C) Authentification forte")
	fmt.Println("D) Chiffrement de bout en bout")

	var r2 string
	fmt.Print("Votre réponse: ")
	fmt.Scanln(&r2)

	switch r2 {
	case "B", "b":
		fmt.Println("Bonne réponse !")
		CyberPoints += 1
	case "A", "a", "C", "c", "D", "d":
		fmt.Println("Mauvaise réponse")
	default:
		fmt.Println("Réponse incorrecte, veuillez entrer A, B, C ou D.")
	}

	fmt.Println("Quel protocole sécurise les communications web en chiffrant les données entre le client et le serveur ?")
	fmt.Println("A) HTTP")
	fmt.Println("B) FTP")
	fmt.Println("C) HTTPS")
	fmt.Println("D) SMTP")

	var r3 string
	fmt.Print("Votre réponse: ")
	fmt.Scanln(&r3)

	switch r3 {
	case "C", "c":
		fmt.Println("Bonne réponse !")
		CyberPoints += 1
	case "A", "a", "B", "b", "D", "d":
		fmt.Println("Mauvaise réponse")
	default:
		fmt.Println("Réponse incorrecte, veuillez entrer A, B, C ou D.")
	}

	fmt.Println("Quel type d'attaque vise à tromper l'utilisateur pour qu'il révèle ses identifiants ?")
	fmt.Println("A) Cross-Site Scripting (XSS)")
	fmt.Println("B) Brute Force")
	fmt.Println("C) Phishing")
	fmt.Println("D) Buffer Overflow")

	var r4 string
	fmt.Print("Votre réponse: ")
	fmt.Scanln(&r4)

	switch r4 {
	case "C", "c":
		fmt.Println("Bonne réponse !")
		CyberPoints += 1
	case "A", "a", "B", "b", "D", "d":
		fmt.Println("Mauvaise réponse")
	default:
		fmt.Println("Réponse incorrecte, veuillez entrer A, B, C ou D.")
	}

	_ = r0
	_ = r1
	_ = r2
	_ = r3
	_ = r4

	fmt.Println("Vous avez scoré ", CyberPoints, "/5")
	choix_quiz()
}
