package site

import (
    "html/template"
    "net/http"
    "github.com/kenan-rhoton/vineyard/models"
)

var landingTemplate = template.Must(template.ParseFiles("site/templates/landing.html"))

func landingHandler(w http.ResponseWriter, r *http.Request) {
    c := &models.Church{}
    err := models.Grab(c,"castelldefels")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    err = landingTemplate.ExecuteTemplate(w, "landing.html", c)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func LandingSetup() {
    http.HandleFunc("/", landingHandler)
}
