package main

import (
    "html/template"
    "net/http"
)

type PageVariables struct {
    Title       string
    TemplateName string
}

func renderTemplate(w http.ResponseWriter, tmpl string, vars PageVariables) {
    t, err := template.ParseFiles("templates/layout.html", "templates/"+tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    err = t.ExecuteTemplate(w, "layout.html", vars)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    vars := PageVariables{
        Title: "Home",
        TemplateName: "index.html",
    }
    renderTemplate(w, "index.html", vars)
}

func experienceHandler(w http.ResponseWriter, r *http.Request) {
    vars := PageVariables{
        Title: "Experience",
        TemplateName: "experience.html",
    }
    renderTemplate(w, "experience.html", vars)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        // Handle form submission logic here (e.g. sending an email)
    }
    vars := PageVariables{
        Title: "Contact",
        TemplateName: "contact.html",
    }
    renderTemplate(w, "contact.html", vars)
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/experience", experienceHandler)
    http.HandleFunc("/contact", contactHandler)

    // Serve static files
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    http.ListenAndServe(":8080", nil)
}
