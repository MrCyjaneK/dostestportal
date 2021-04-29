package webui

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"strconv"

	"github.com/gobuffalo/packr/v2"
)

var Port = 0
var jar, _ = cookiejar.New(nil)
var httpClient = http.Client{
	Jar: jar,
}
var status = "DDoSing..."
var count = 0

// Start the webui
func Start() {
	if Port == 0 {
		Port = 2000 + rand.Intn(10000)
	}
	Port = 15932 // Sorry for hardcoding
	html := packr.New("webui", "./html")
	http.Handle("/", http.FileServer(html))
	http.HandleFunc("/api/hack", apiHack)
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html;charset=utf-8")
		fmt.Fprintln(w, "<!DOCTYPE html>\n<html>\n<head>\n<meta http-equiv=\"refresh\" content=\"1\">\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">")
		fmt.Fprintln(w, "<style>body { font-size: 150% }</style>")
		fmt.Fprintln(w, "<body><div style=\"margin: auto; max-width:650px;\">")
		fmt.Fprintln(w, status, "<br />tests:", count)
		fmt.Fprintln(w, "<hr /> Owned by Czarek Nakamoto | <a href=\"https://mrcyjanek.net\">mrcyjanek.net</a>")
		fmt.Fprintln(w, "</div></body></html>")
		//fmt.Fprintln(w, string(out))
	})
	//http.HandleFunc("/api/ping", apiPing)
	go http.ListenAndServe(":"+strconv.Itoa(Port), nil)
	fmt.Println("[webui][Start] Listening on 127.0.0.1:" + strconv.Itoa(Port))
}

func apiHack(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "text/html")
	go realHackOk(req.FormValue("url"))
	fmt.Fprint(w, `<meta http-equiv="Refresh" content="0; url='/status'" />`)
}
