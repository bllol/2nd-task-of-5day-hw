package main

import (
	"html/template"
	"net/http"
)

type Contact struct {
	Name    string
	Email   string
	Message string
}

var templates = template.Must(template.ParseGlob("templates/*html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "form.html", nil)
	if err != nil {
		http.Error(w, "Failed to load file", http.StatusInternalServerError)
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only post is allowed", http.StatusMethodNotAllowed)
		return
	}

	fullName := r.FormValue("fullname")
	email := r.FormValue("email")
	message := r.FormValue("message")

	user := Contact{
		Name:    fullName,
		Email:   email,
		Message: message,
	}

	err := templates.ExecuteTemplate(w, "submit.html", user)
	if err != nil {
		http.Error(w, "Template execution is failed", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/submit", submitHandler)

	http.ListenAndServe(":8080", nil)
}
