package admin

import (
    "html/template"
    "net/http"
    "github.com/kenan-rhoton/vineyard/models"
    "regexp"
)

var adminTemplate = template.Must(template.ParseGlob("site/templates/admin/*.html"))
var lastString = regexp.MustCompile("^/[a-z/]*/([a-zA-Z0-9]*)$")

func adminHandler(w http.ResponseWriter, r *http.Request) {
    var c []models.Church
    err := models.GrabAll(&models.Church{},&c)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    m := models.ListAll()
    data := &struct{
        Churches []models.Church
        Models []string
        CurrentModel string
    }{
        Churches: c,
        Models: m,
        CurrentModel: "Iglesias",
    }
    err = adminTemplate.ExecuteTemplate(w, "admin.html", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func Setup() {
    ChurchSetup()
    http.HandleFunc("/admin/", adminHandler)
}
