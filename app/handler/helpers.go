package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"text/template"

	"asciiWeb/internal"
)

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
		handleStatusCode(w, http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func extractFormData(r *http.Request) int {
	err := r.ParseForm()
	if err != nil {
		return 400
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	if textReg := regexp.MustCompile(`^\r\n+`); textReg.MatchString(text) {
		Pagedata.Text = "\r\n" + text
	} else {
		Pagedata.Text = text
	}
	Pagedata.Banner = banner
	return 200
}

func validateFormData() (status int) {
	Pagedata.FormError = internal.UserInputChecker(Pagedata.Text)
	if !IsBanner(Pagedata.Banner) || Pagedata.FormError != "" {
		return 400
	}
	return 200
}

func IsBanner(banner string) bool {
	return banner == "standard" || banner == "shadow" || banner == "thinkertoy"
}

func handleStatusCode(w http.ResponseWriter, status int) {
	switch status {
	case 200:
		renderTemplate(w, "index.html", Pagedata, status)
	case 400:
		if Pagedata.FormError != "" {
			renderTemplate(w, "index.html", Pagedata, status)
		} else {
			renderTemplate(w, "errorPage.html", status, status)
		}

	default:
		renderTemplate(w, "errorPage.html", status, status)
	}
}

// function written just to understand how the os.Stat works

// func CheckPaths(path string) {
// 	file_info, err := os.Stat(path)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(file_info.IsDir())
// 	fmt.Println(file_info.Name())

// }
