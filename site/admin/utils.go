package admin

import (
    "net/http"
    "net/url"
)

func postHandler(fn func(val url.Values, which string) error, redirect_url string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            err := r.ParseForm()
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
            m := lastString.FindStringSubmatch(r.URL.Path)
            err = fn(r.Form, m[1])
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
        }
        http.Redirect(w, r, redirect_url, http.StatusFound)
    }
}
