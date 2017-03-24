package main

import(
    "net/http"
    "github.com/kenan-rhoton/vineyard/site"
)

func main() {
    site.Setup()
    http.ListenAndServe(":8080", nil)
}
