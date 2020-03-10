package main

import (
	"html/template"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}

type IndexPage struct {
	Title       string
	Description string
}

type GeneralTemplates struct {
	Header, Footer string
}

func getGenericTemplates() string {
	genericTemplates := []string{"./templates/index.html", "./templates/header.html", "./templates/footer.html"}
	return strings.Join(genericTemplates, ",")
}

func Index(w http.ResponseWriter, r *http.Request) {
	generalTemplates := GeneralTemplates{"./templates/header.html", "./templates/footer.html"}
	indexTemplate := "./templates/index.html"
	pageData := IndexPage{"Antonio Djigo", "A website made in go, because let's go"}
	t, _ := template.ParseFiles(generalTemplates.Header, generalTemplates.Footer, indexTemplate)
	t.Execute(w, pageData)
}
