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

    r.HandleFunc("/{username}", middleware.AuthRequired(userGetHandler)).Methods("GET")
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
                utils.InternalServerError(w)
        }
        return
    }

    userId, err := user.GetId()
    if err != nil {
        utils.InternalServerError(w)
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
    if err == models.ErrUsernameTaken {
        utils.ExecuteTemplate(w, "register.html", "username taken")
        return
    } else if err != nil {
        utils.InternalServerError(w)
        return
    }

    http.Redirect(w, r, "/login", 302)
}


func dashboardGetHandler(w http.ResponseWriter, r *http.Request) {
    updates, err := models.GetAllUpdates()
    if err != nil {
        utils.InternalServerError(w)
        return
    }
    utils.ExecuteTemplate(w, "dashboard.html", struct {
        Title string
        Updates []*models.Update
    }{
        Title: "All Updates",
        Updates: updates,
    })
}

func dashboardPostHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := sessions.Store.Get(r, "session")
    untypedUserId := session.Values["user_id"]
    userId, ok := untypedUserId.(int64)
    if !ok {
        utils.InternalServerError(w)
        return
    }

    r.ParseForm()
    update := r.PostForm.Get("update")
    err := models.PostUpdate(userId, update)
    if err != nil {
        utils.InternalServerError(w)
        return
    }
    http.Redirect(w, r, "/dashboard", 302)
}


func userGetHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    username := vars["username"]
    user, err := models.GetUserByUsername(username)
    if err != nil {
        utils.InternalServerError(w)
        return
    }

    userId, err := user.GetId()
    if err != nil {
        utils.InternalServerError(w)
        return
    }

    updates, err := models.GetUpdates(userId)
    if err != nil {
        utils.InternalServerError(w)
        return
    }
    utils.ExecuteTemplate(w, "dashboard.html", struct {
        Title string
        Updates []*models.Update
    }{
        Title: username,
        Updates: updates,
    })
}
