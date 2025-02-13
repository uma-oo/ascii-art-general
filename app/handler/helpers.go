package handler

import (
	"fmt"
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
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "<h1 style='color: #424242; text-align:center'>Internal server Error 500</h1>")
		return
	}
	w.WriteHeader(status)
	err = t.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "<h1 style='color: #424242; text-align:center'>Internal server Error 500</h1>")
		return
	}
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
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	pagedata := NewData()
	if textReg := regexp.MustCompile(`^\r\n+`); textReg.MatchString(text) {
		pagedata.Text = "\r\n" + text
	} else {
		pagedata.Text = text
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
		if  pageData!=nil && pageData.FormError != "" {
			renderTemplate(w, "index.html", pageData, status)
		} else {
			renderTemplate(w, "errorPage.html", status, status)
		}
	default:
		renderTemplate(w, "errorPage.html", status, status)
	}
}

