package admin

import (
    "net/http"
    "net/url"
    "github.com/kenan-rhoton/vineyard/models"
)

func adminDeleteChurchHandler(w http.ResponseWriter, r *http.Request) {
    m := lastString.FindStringSubmatch(r.URL.Path)
    err := models.Delete(&models.Church{}, m[1])
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    http.Redirect(w, r, "/admin/", http.StatusFound)
}

func createChurch(form url.Values, _ string) error {
    return models.Create(&models.Church{}, form)
}

func updateChurch(form url.Values, target string) error {
    return models.Update(&models.Church{}, form, target)
}

func ChurchSetup() {
    http.HandleFunc("/admin/church/new/", postHandler(createChurch, "/admin/"))
    http.HandleFunc("/admin/church/delete/", adminDeleteChurchHandler)
    http.HandleFunc("/admin/church/update/", postHandler(updateChurch, "/admin/"))
}
