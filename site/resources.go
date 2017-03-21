package site

import (
    "net/http"
)

func ResourcesSetup() {
    http.Handle("/public/",http.StripPrefix("/public/",http.FileServer(http.Dir("site/public"))))
}
