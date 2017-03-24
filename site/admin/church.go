package admin

import (
    "net/http"
    "github.com/kenan-rhoton/vineyard/models"
)

func adminDeleteChurchHandler(w http.ResponseWriter, r *http.Request) {
    m := lastString.FindStringSubmatch(r.URL.Path)
    err := models.DeleteChurch(m[1])
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    http.Redirect(w, r, "/admin/", http.StatusFound)
}

func ChurchSetup() {
    http.HandleFunc("/admin/church/new/", postHandler(models.SaveChurch, "/admin/"))
    http.HandleFunc("/admin/church/delete/", adminDeleteChurchHandler)
    http.HandleFunc("/admin/church/update/", postHandler(models.UpdateChurch, "/admin/"))
}
