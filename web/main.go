package web

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Question struct {
	ID      string
	Text    string
	Options []string
	Answer  string // correct option like "A", "B", "C", "D"
}

//go:embed templates/*.tmpl
var templatesFS embed.FS

var funcMap = template.FuncMap{
	"add":  func(a, b int) int { return a + b },
	"add1": func(a int) int { return a + 1 },
}

var templates = template.Must(template.New("").Funcs(funcMap).ParseFS(templatesFS, "templates/*.tmpl"))

func listTemplates() []string {
	var names []string
	for _, t := range templates.Templates() {
		if t == nil || t.Name() == "" {
			continue
		}
		names = append(names, t.Name())
	}
	return names
}

var quizzes = map[string]struct {
	Title string
	Qs    []Question
}{
	"cyber": {
		Title: "Cybersécurité",
		Qs: []Question{
			{ID: "q1", Text: "Quel type d'attaque consiste à submerger un serveur de requêtes pour le rendre indisponible ?", Options: []string{"A) Phishing", "B) SQL Injection", "C) DDoS", "D) Man-in-the-Middle"}, Answer: "C"},
			{ID: "q2", Text: "Quel est le rôle d'un pare-feu (firewall) ?", Options: []string{"A) Chiffrer les données", "B) Filtrer le trafic réseau selon des règles", "C) Stocker les mots de passe", "D) Scanner les virus dans les e-mails"}, Answer: "B"},
			{ID: "q3", Text: "Quel principe de sécurité recommande d'accorder le minimum de privilèges nécessaires ?", Options: []string{"A) Défense en profondeur", "B) Principe du moindre privilège", "C) Authentification forte", "D) Chiffrement de bout en bout"}, Answer: "B"},
			{ID: "q4", Text: "Quel protocole sécurise les communications web en chiffrant les données entre le client et le serveur ?", Options: []string{"A) HTTP", "B) FTP", "C) HTTPS", "D) SMTP"}, Answer: "C"},
			{ID: "q5", Text: "Quel type d'attaque vise à tromper l'utilisateur pour qu'il révèle ses identifiants ?", Options: []string{"A) Cross-Site Scripting (XSS)", "B) Brute Force", "C) Phishing", "D) Buffer Overflow"}, Answer: "C"},
		},
	},
	"info": {
		Title: "Informatique",
		Qs: []Question{
			{ID: "q1", Text: "Dans une architecture MVC, à quoi correspond le modèle (Model) ?", Options: []string{"A) À la logique métier et aux données", "B) À l’interface utilisateur", "C) À la gestion des requêtes HTTP", "D) À la configuration du serveur"}, Answer: "A"},
			{ID: "q2", Text: "Quelle technologie est côté client dans une application full stack classique ?", Options: []string{"A) Node.js", "B) MySQL", "C) JavaScript (dans le navigateur)", "D) Express.js"}, Answer: "C"},
			{ID: "q3", Text: "Quel est le rôle principal de Git dans le développement ?", Options: []string{"A) Héberger une base de données", "B) Gérer le versionnement du code", "C) Exécuter du code JavaScript côté serveur", "D) Styliser une page web"}, Answer: "B"},
			{ID: "q4", Text: "Dans une requête HTTP, quelle méthode est utilisée pour récupérer des données sans les modifier ?", Options: []string{"A) POST", "B) PUT", "C) GET", "D) DELETE"}, Answer: "C"},
			{ID: "q5", Text: "Quel format est le plus couramment utilisé pour échanger des données entre un serveur et un client ?", Options: []string{"A) XML", "B) CSV", "C) JSON", "D) YAML"}, Answer: "C"},
		},
	},
	"ia": {
		Title: "IA & Data",
		Qs: []Question{
			{ID: "q1", Text: "Dans un problème de classification, quel est l’objectif principal ?", Options: []string{"A) Prédire une valeur numérique", "B) Regrouper des données similaires", "C) Assigner une étiquette à une donnée", "D) Réduire la dimension des données"}, Answer: "C"},
			{ID: "q2", Text: "Quelle bibliothèque Python est la plus utilisée pour le machine learning ?", Options: []string{"A) Pandas", "B) NumPy", "C) Scikit-learn", "D) Matplotlib"}, Answer: "C"},
			{ID: "q3", Text: "Quel type de données est le plus adapté pour entraîner un modèle de réseau de neurones ?", Options: []string{"A) Données textuelles non structurées", "B) Données numériques normalisées", "C) Données catégorielles brutes", "D) Données manquantes non traitées"}, Answer: "B"},
			{ID: "q4", Text: "Qu’est-ce que le surapprentissage (overfitting) ?", Options: []string{"A) Le modèle est trop simple et ne capture pas les patterns", "B) Le modèle apprend le bruit des données d’entraînement", "C) Le modèle généralise mal sur de nouvelles données", "D) Le modèle est trop lent à s’entraîner"}, Answer: "B"},
			{ID: "q5", Text: "Quelle étape est essentielle avant d’entraîner un modèle ?", Options: []string{"A) Visualiser les résultats", "B) Nettoyer et prétraiter les données", "C) Déployer le modèle en production", "D) Choisir une couleur pour les graphiques"}, Answer: "B"},
		},
	},
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// try bare filename first, then embedded path
	if err := templates.ExecuteTemplate(w, "index.tmpl", quizzes); err != nil {
		if err2 := templates.ExecuteTemplate(w, "templates/index.tmpl", quizzes); err2 != nil {
			log.Printf("template index error (tried both): %v ; %v", err, err2)
			log.Printf("available templates: %v", listTemplates())
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}
}

func quizHandler(w http.ResponseWriter, r *http.Request) {
	// URL: /quiz/{name}
	name := r.URL.Path[len("/quiz/"):]
	q, ok := quizzes[name]
	if !ok {
		http.NotFound(w, r)
		return
	}
	data := struct {
		Name  string
		Title string
		Qs    []Question
	}{
		Name:  name,
		Title: q.Title,
		Qs:    q.Qs,
	}
	if err := templates.ExecuteTemplate(w, "quiz.tmpl", data); err != nil {
		if err2 := templates.ExecuteTemplate(w, "templates/quiz.tmpl", data); err2 != nil {
			log.Printf("template quiz error (tried both): %v ; %v", err, err2)
			log.Printf("available templates: %v", listTemplates())
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	// URL: /quiz/{name}/submit
	name := r.URL.Path[len("/quiz/") : len(r.URL.Path)-len("/submit")]
	qset, ok := quizzes[name]
	if !ok {
		http.NotFound(w, r)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "cannot parse form", http.StatusBadRequest)
		return
	}
	score := 0
	for i, ques := range qset.Qs {
		field := fmt.Sprintf("q%d", i+1)
		ans := r.FormValue(field)
		if ans == "" {
			// unanswered treated as wrong
			continue
		}
		// normalize to uppercase first char
		if ans[0] == 'a' {
			ans = "A"
		}
		if ans[0] == 'b' {
			ans = "B"
		}
		if ans[0] == 'c' {
			ans = "C"
		}
		if ans[0] == 'd' {
			ans = "D"
		}
		if ans == ques.Answer {
			score++
		}
	}
	data := struct {
		Title string
		Score int
		Total int
	}{
		Title: qset.Title,
		Score: score,
		Total: len(qset.Qs),
	}
	if err := templates.ExecuteTemplate(w, "result.tmpl", data); err != nil {
		if err2 := templates.ExecuteTemplate(w, "templates/result.tmpl", data); err2 != nil {
			log.Printf("template result error (tried both): %v ; %v", err, err2)
			log.Printf("available templates: %v", listTemplates())
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}
}

func templatesListHandler(w http.ResponseWriter, r *http.Request) {
	names := listTemplates()
	for _, n := range names {
		fmt.Fprintln(w, n)
	}
}

func Start(addr string) error {
	// register handlers with a simple logger wrapper so requests are visible in logs
	logWrap := func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s", r.Method, r.URL.Path)
			h(w, r)
		}
	}

	http.HandleFunc("/", logWrap(indexHandler))
	http.HandleFunc("/__templates", logWrap(templatesListHandler))
	http.HandleFunc("/quiz/", logWrap(func(w http.ResponseWriter, r *http.Request) {
		// dispatch /quiz/{name} and /quiz/{name}/submit
		if strings.HasSuffix(r.URL.Path, "/submit") {
			submitHandler(w, r)
			return
		}
		quizHandler(w, r)
	}))
	log.Printf("Starting server at http://%s", addr)
	return http.ListenAndServe(addr, nil)
}
