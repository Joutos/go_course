package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type usuario struct {
	Nome  string
	Email string
	Idade uint
}

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("*.html"))
}

func main() {
	http.HandleFunc("/", home(usuario{"João Paulo", "joaodemeto@gmail.com", 22}))
	// Listar templates
	for _, tmpl := range templates.Templates() {
		fmt.Println(tmpl.Name())
	}
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func home(user usuario) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index.html", user)
	}
}
