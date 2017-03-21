package site

import (
    "html/template"
    "net/http"
    "github.com/kenan-rhoton/vineyard/models"
)

var adminTemplate = template.Must(template.ParseFiles("site/templates/admin.html"))

func adminHandler(w http.ResponseWriter, r *http.Request) {
    c := models.GetDefaultChurch()
    err := adminTemplate.ExecuteTemplate(w, "admin.html", c)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func AdminSetup() {
    http.HandleFunc("/admin/", adminHandler)
}

