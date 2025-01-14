package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"asciiWeb/internal"
)

type Data struct {
	Text          string
	Banner        string
	FormError     string
	AsciiArt      string
	is_downloaded bool // just to check if the user hitted the button downlaod
}

var Pagedata = Data{}

// handler for the path "/"
func HandleMainPage(w http.ResponseWriter, r *http.Request) {
	Pagedata = Data{}
	if r.URL.Path != `/` {
		handleStatusCode(w, http.StatusNotFound)
		return
	}
	if !requestMethodChecker(w, r, http.MethodGet) {
		return
	}
	renderTemplate(w, "index.html", Pagedata, http.StatusOK)
}

// handler for the path "/ascii-art
func HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
	Pagedata = Data{}
	if !requestMethodChecker(w, r, http.MethodPost) {
		return
	}

	if status := extractFormData(r); status != 200 {
		handleStatusCode(w, status)
		return
	}

	if status := validateFormData(); status != 200 {
		handleStatusCode(w, status)
		return
	}

	asciiArt, status := internal.Ascii(Pagedata.Text, Pagedata.Banner)
	if status != 200 {
		handleStatusCode(w, status)
		return
	}

	Pagedata.AsciiArt = asciiArt
	renderTemplate(w, "index.html", Pagedata, http.StatusOK)
}

// function to serve the files and avoid the listing of directories
func HandleAssets(w http.ResponseWriter, r *http.Request) {
	if !requestMethodChecker(w, r, http.MethodGet) {
		return
	}
	if !strings.HasPrefix(r.URL.Path, "/assets") {
		handleStatusCode(w, http.StatusNotFound)
		return
	} else {
		file_info, err := os.Stat(r.URL.Path[1:])
		if err != nil {
			handleStatusCode(w, http.StatusNotFound)
			return
		} else if file_info.IsDir() {
			handleStatusCode(w, http.StatusForbidden)
			return
		} else {
			http.ServeFile(w, r, r.URL.Path[1:])
		}
	}
}

// function to handle the download process
func HandleDownloads(w http.ResponseWriter, r *http.Request) {
	// The method should be POST
	if !requestMethodChecker(w, r, http.MethodGet) {
		return
	}
	if Pagedata.AsciiArt != "" {
		w.Header().Add("Content-Type", "text/plain")
		w.Header().Add("Content-Disposition", "attachement")
		w.Header().Add("Content-Length", fmt.Sprint(len(Pagedata.AsciiArt)))
		file, err := os.Create("ascii-art.txt")
		if err != nil {
			handleStatusCode(w, http.StatusInternalServerError)
			return
		}
		err = os.WriteFile(file.Name(), []byte(Pagedata.AsciiArt), 0o644)
		if err != nil {
			handleStatusCode(w, http.StatusInternalServerError)
		}
		fmt.Println(file.Name())
		http.ServeFile(w, r, file.Name())

	} else {
		handleStatusCode(w, http.StatusBadRequest)
	}
}
