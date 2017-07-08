package hello

import (
    "fmt"
    "net/http"
    "google.golang.org/appengine"
    "google.golang.org/appengine/urlfetch"
)
func init() {
    http.HandleFunc("/", handler)
}
func handler(w http.ResponseWriter, r *http.Request) {
        ctx := appengine.NewContext(r)
        client := urlfetch.Client(ctx)
        resp, err := client.Get("https://www.qzhou.com.cn/")
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        fmt.Fprintf(w, "HTTP GET returned status %v", resp.Status)
}
