package site

import (
    "net/http"
    "github.com/kenan-rhoton/vineyard/models"
    "html/template"
)

var loginTemplate = template.Must(template.ParseFiles("site/templates/login.html"))
var signupTemplate = template.Must(template.ParseFiles("site/templates/signup.html"))

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        //Someone is logging in
        r.ParseForm()
        key, err := models.Login(r.FormValue("user"), r.FormValue("pass"))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        c := &http.Cookie{Name: "SessionKey", Value: key}
        http.SetCookie(w, c)
        http.Redirect(w, r, "/admin/", http.StatusFound)
    } else if r.Method == "GET" {
        //Someone wants to see the login page
        err := loginTemplate.ExecuteTemplate(w, "login.html", "nothing needed")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        //Someone is creating a user
        r.ParseForm()
        err := models.Create(&models.User{}, r.Form)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        loginHandler(w,r)
    } else if r.Method == "GET" {
        //Someone wants to see the login page
        err := signupTemplate.ExecuteTemplate(w, "signup.html", "nothing needed")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}

func LoginSetup() {
    http.HandleFunc("/login/", loginHandler)
    http.HandleFunc("/signup/", signupHandler)
}
