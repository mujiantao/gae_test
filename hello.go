package hello

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"appengine"
	"appengine/urlfetch"
)

func init() {
	http.HandleFunc("/", serve)
}

func serve(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	url := &url.URL{Path: "http://localhost:8888"}

	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.Transport = &urlfetch.Transport{Context: c}
	proxy.ServeHTTP(w, r)
}
