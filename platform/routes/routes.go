package routes

import (
    "net/http"
    "github.com/gorilla/mux"
    "../models"
    "../sessions"
    "../middleware"
    "../utils"
)

func NewRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/", indexHandler).Methods("GET")
    r.HandleFunc("/login", loginGetHandler).Methods("GET")
    r.HandleFunc("/login", loginPostHandler).Methods("POST")
    r.HandleFunc("/register", registerGetHandler).Methods("GET")
    r.HandleFunc("/register", registerPostHandler).Methods("POST")
    r.HandleFunc("/dashboard", middleware.AuthRequired(dashboardGetHandler)).Methods("GET")
    r.HandleFunc("/dashboard", middleware.AuthRequired(dashboardPostHandler)).Methods("POST")

    fs := http.FileServer(http.Dir("./static/"))
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

    return r
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

    user, err := models.AuthenticateUser(username, password)
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

    userId, err := user.GetId()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    session, _ := sessions.Store.Get(r, "session")
    session.Values["user_id"] = userId
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
    updates, err := models.GetUpdates()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    utils.ExecuteTemplate(w, "dashboard.html", updates)
}

func dashboardPostHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := sessions.Store.Get(r, "session")
    untypedUserId := session.Values["user_id"]
    userId, ok := untypedUserId.(int64)
    if !ok {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }

    r.ParseForm()
    update := r.PostForm.Get("update")
    err := models.PostUpdate(userId, update)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    http.Redirect(w, r, "/dashboard", 302)
}
