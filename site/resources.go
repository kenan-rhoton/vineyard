package site

import (
    "net/http"
    "github.com/kenan-rhoton/vineyard/models"
)

func ResourcesSetup() {
    http.Handle(
        "/public/",
        http.StripPrefix(
            "/public/",
            http.FileServer(
                http.Dir("public")
            )
        )
    )
}
