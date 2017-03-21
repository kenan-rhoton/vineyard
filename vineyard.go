package main

import(
    "net/http"
    "github.com/kenan-rhoton/vineyard/site"
)

func main() {
    site.AdminSetup()
    site.ResourcesSetup()
    site.LandingSetup()
    http.ListenAndServe(":8080", nil)
}
