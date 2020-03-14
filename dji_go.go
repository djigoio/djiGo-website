package main

import (
	"html/template"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}

type IndexPage struct {
	Title       string
	Description string
}

type GeneralTemplates struct {
	Header, Footer, Languages, Contact string
}

func Index(w http.ResponseWriter, r *http.Request) {
	generalTemplates := GeneralTemplates{"./templates/header.html", "./templates/footer.html", "./templates/languages.html", "./templates/contact.html"}
	pageData := IndexPage{"Antonio Djigo", "I'm Antonio Djigo, Web Developer from the Canary Islands"}
	t, _ := template.ParseFiles("./index.html", generalTemplates.Header, generalTemplates.Footer, generalTemplates.Languages, generalTemplates.Contact)
	t.Execute(w, pageData)
}
