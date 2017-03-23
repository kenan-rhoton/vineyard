package site

import (
    "net/url"
    "html/template"
    "net/http"
    "github.com/kenan-rhoton/vineyard/models"
    "regexp"
)

var adminTemplate = template.Must(template.ParseFiles("site/templates/admin.html"))
var lastString = regexp.MustCompile("^/[a-z/]*/([a-zA-Z0-9]+)$")

func adminHandler(w http.ResponseWriter, r *http.Request) {
    c := models.GetChurches()
    m := models.GetModels()
    data := &struct{
        Churches []models.Church
        Models []*models.Model
    }{
        Churches: c,
        Models: m,
    }
    err := adminTemplate.ExecuteTemplate(w, "admin.html", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func postHandler(fn func(val url.Values, which string) error, redirect_url string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            err := r.ParseForm()
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
            m := lastString.FindStringSubmatch(r.URL.Path)
            err = fn(r.Form, m[1])
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
        }
        http.Redirect(w, r, redirect_url, http.StatusFound)
    }
}

func adminDeleteChurchHandler(w http.ResponseWriter, r *http.Request) {
    m := lastString.FindStringSubmatch(r.URL.Path)
    err := models.DeleteChurch(m[1])
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    http.Redirect(w, r, "/admin/", http.StatusFound)
}

func AdminSetup() {
    http.HandleFunc("/admin/church/new/", postHandler(models.SaveChurch, "/admin/"))
    http.HandleFunc("/admin/church/delete/", adminDeleteChurchHandler)
    http.HandleFunc("/admin/church/update/", postHandler(models.UpdateChurch, "/admin/"))
    http.HandleFunc("/admin/", adminHandler)
}

