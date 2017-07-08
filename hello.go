package hello

import (
    "fmt"
    "net/http"
    "google.golang.org/appengine"
    "google.golang.org/appengine/urlfetch"
    "io/ioutil"
)
func init() {
    http.HandleFunc("/", handler)
}

var headers = map[string]string{
	"User-Agent": "MQQBrowser/26 Mozilla/5.0 (Linux; U; Android 2.3.7; zh-cn; MB200 Build/GRJ22; CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
	"Referer":    "https://yun.baidu.com/share/home?uk=325913312#category/type=0",
	"Cookie":     "",
}

func handler(w http.ResponseWriter, r *http.Request) {
        ctx := appengine.NewContext(r)
        client := urlfetch.Client(ctx)
        resp, err := client.Get("https://www.qzhou.com.cn/")
        for k, v := range headers {
      		req.Header.Set(k, v)
      	}
        var resp *http.Response
      	resp, err = client.Do(req)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        var body []byte
      	body, err = ioutil.ReadAll(resp.Body)
        defer resp.Body.Close()
        fmt.Fprintf(w, string(body))
}
