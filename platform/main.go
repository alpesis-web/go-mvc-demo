package main

import(
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
    "github.com/go-redis/redis"
    "html/template"
)

var client *redis.Client
var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))

var templates *template.Template


func main() {
    client = redis.NewClient(&redis.Options{
        Addr: "mandelbrot-redis:6379",
    })
    templates = template.Must(template.ParseGlob("templates/*.html"))

    r := mux.NewRouter()
    r.HandleFunc("/", indexHandler).Methods("GET")
    r.HandleFunc("/dashboard", dashboardGetHandler).Methods("GET")
    r.HandleFunc("/dashboard", dashboardPostHandler).Methods("POST")

    fs := http.FileServer(http.Dir("./static/"))
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

    http.Handle("/", r)
    http.ListenAndServe(":9090", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    templates.ExecuteTemplate(w, "index.html", nil)
}

func dashboardGetHandler(w http.ResponseWriter, r *http.Request) {
    comments, err := client.LRange("comments", 0, 10).Result()
    if err != nil {
        fmt.Printf("%v\n", err)
        return
    }
    templates.ExecuteTemplate(w, "dashboard.html", comments)
}

func dashboardPostHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    comment := r.PostForm.Get("comment")
    client.LPush("comments", comment)
    http.Redirect(w, r, "/dashboard", 302)
}
