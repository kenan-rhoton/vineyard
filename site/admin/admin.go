package admin

import (
    "html/template"
    "net/http"
    "github.com/kenan-rhoton/vineyard/models"
    "regexp"
)

var adminTemplate = template.Must(template.ParseGlob("site/templates/admin/*.html"))
var lastString = regexp.MustCompile("^/[a-z/]*/([a-zA-Z0-9]+)$")

func adminHandler(w http.ResponseWriter, r *http.Request) {
    c := models.GetChurches()
    m := models.GetModels()
    data := &struct{
        Churches []models.Church
        Models []*models.Model
        CurrentModel string
    }{
        Churches: c,
        Models: m,
        CurrentModel: "Iglesias",
    }
    err := adminTemplate.ExecuteTemplate(w, "admin.html", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func Setup() {
    http.HandleFunc("/admin/", adminHandler)
}
