package main

import(
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "html/template"
)

var templates *template.Template

func main() {
    templates = template.Must(template.ParseGlob("templates/*.html"))
    r := mux.NewRouter()
    r.HandleFunc("/", indexHandler).Methods("GET")
    r.HandleFunc("/dashboard", dashboardHandler).Methods("GET")

    http.Handle("/", r)
    http.ListenAndServe(":9090", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    templates.ExecuteTemplate(w, "index.html", nil)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to dashboard!")
}

