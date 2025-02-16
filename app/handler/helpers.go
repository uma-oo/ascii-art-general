package handler

import (
	"bytes"
	"fmt"
	"html"
	"net/http"
	"regexp"
	"text/template"

	"asciiWeb/internal"
)

type Data struct {
	Text      string
	Banner    string
	FormError string
	AsciiArt  string
	Download  bool
}

// Constructor of the type
func NewData() *Data {
	return &Data{
		Text:      "",
		Banner:    "",
		FormError: "",
		AsciiArt:  "",
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}, status int) {
     var buf bytes.Buffer
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "<h1 style='color: #424242; text-align:center'>Internal server Error 500</h1>")
		return
	}
	w.WriteHeader(status)
	err = t.ExecuteTemplate(&buf, tmpl, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "<h1 style='color: #424242; text-align:center'>Internal server Error 500</h1>")
		return
	}
	w.Write(buf.Bytes())
}

func requestMethodChecker(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		handleStatusCode(w, http.StatusMethodNotAllowed, nil)
		return false
	}
	return true
}

func extractFormData(r *http.Request) (int, *Data) {
	err := r.ParseForm()
	if err != nil {
		return 400, nil
	}
	text := html.EscapeString(r.FormValue("text"))
	banner := r.FormValue("banner")
	download := r.FormValue("want_to_download")
	pagedata := NewData()
	if textReg := regexp.MustCompile(`^\r\n+`); textReg.MatchString(text) {
		pagedata.Text = "\r\n" + text
	} else {
		pagedata.Text = text
	}
	if download == "true" {
		pagedata.Download = true
	}
	pagedata.Banner = banner
	return 200, pagedata
}

func validateFormData(pageData *Data) (status int) {
	pageData.FormError = internal.UserInputChecker(pageData.Text)
	if !IsBanner(pageData.Banner) || pageData.FormError != "" {
		return 400
	}
	return 200
}

func IsBanner(banner string) bool {
	return banner == "standard" || banner == "shadow" || banner == "thinkertoy"
}

func handleStatusCode(w http.ResponseWriter, status int, pageData *Data) {
	switch status {
	case 200:
		renderTemplate(w, "index.html", pageData, status)
	case 400:
		if pageData != nil && pageData.FormError != "" {
			renderTemplate(w, "index.html", pageData, status)
		} else {
			renderTemplate(w, "errorPage.html", status, status)
		}
	default:
		renderTemplate(w, "errorPage.html", status, status)
	}
}
