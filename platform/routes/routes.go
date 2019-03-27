package routes

import (
    "net/http"
    "github.com/gorilla/mux"
    "../models"
    "../sessions"
    "../utils"
)

func NewRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/", indexHandler).Methods("GET")
    r.HandleFunc("/login", loginGetHandler).Methods("GET")
    r.HandleFunc("/login", loginPostHandler).Methods("POST")
    r.HandleFunc("/register", registerGetHandler).Methods("GET")
    r.HandleFunc("/register", registerPostHandler).Methods("POST")
    r.HandleFunc("/dashboard", AuthRequired(dashboardGetHandler)).Methods("GET")
    r.HandleFunc("/dashboard", AuthRequired(dashboardPostHandler)).Methods("POST")

    fs := http.FileServer(http.Dir("./static/"))
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

    return r
}

func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        session, _ := sessions.Store.Get(r, "session")
        _, ok := session.Values["username"]
        if !ok {
            http.Redirect(w, r, "/login", 302)
            return
        }
        handler.ServeHTTP(w, r)
    }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    utils.ExecuteTemplate(w, "index.html", nil)
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
    utils.ExecuteTemplate(w, "login.html", nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    username := r.PostForm.Get("username")
    password := r.PostForm.Get("password")

    err := models.AuthenticateUser(username, password)
    if err != nil {
        switch err {
            case models.ErrUserNotFound:
                utils.ExecuteTemplate(w, "login.html", "unknown user")
            case models.ErrInvalidLogin:
                utils.ExecuteTemplate(w, "login.html", "invalid login")
            default:
                w.WriteHeader(http.StatusInternalServerError)
                w.Write([]byte("Internal server error"))
        }
        return
    }

    session, _ := sessions.Store.Get(r, "session")
    session.Values["username"] = username
    session.Save(r, w)

    http.Redirect(w, r, "/dashboard", 302)
}

func registerGetHandler(w http.ResponseWriter, r *http.Request) {
    utils.ExecuteTemplate(w, "register.html", nil)
}

func registerPostHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    username := r.PostForm.Get("username")
    password := r.PostForm.Get("password")

    err := models.RegisterUser(username, password)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }

    http.Redirect(w, r, "/login", 302)
}

func dashboardGetHandler(w http.ResponseWriter, r *http.Request) {
    comments, err := models.GetComments()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    utils.ExecuteTemplate(w, "dashboard.html", comments)
}

func dashboardPostHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    comment := r.PostForm.Get("comment")
    err := models.PostComment(comment)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    http.Redirect(w, r, "/dashboard", 302)
}
