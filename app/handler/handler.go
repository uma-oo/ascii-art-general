package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"asciiWeb/internal"
)

// handler for the path "/"
func HandleMainPage(w http.ResponseWriter, r *http.Request) {
	pagedata := NewData()
	if r.URL.Path != `/` {
		handleStatusCode(w, http.StatusNotFound, nil)
		return
	}
	if !requestMethodChecker(w, r, http.MethodGet) {
		return
	}
	renderTemplate(w, "index.html", pagedata, http.StatusOK)
}

// handler for the path "/ascii-art
func (D *Data) HandleAsciiArt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !requestMethodChecker(w, r, http.MethodPost) {
			return
		}
		status, D := extractFormData(r)
		if status != 200 {
			handleStatusCode(w, status, D)
			return
		}
		if status = validateFormData(D); status != 200 {
			handleStatusCode(w, status, D)
			return
		}

		asciiArt, status := internal.Ascii(D.Text, D.Banner)
		if status != 200 {
			handleStatusCode(w, status, D)
		}

		D.AsciiArt = asciiArt
        fmt.Printf("D inside the ascii art handler : %v\n", D)
		renderTemplate(w, "index.html", D, http.StatusOK)
	}
}

// function to serve the files and avoid the listing of directories
func HandleAssets(w http.ResponseWriter, r *http.Request) {
	if !requestMethodChecker(w, r, http.MethodGet) {
		return
	}
	if !strings.HasPrefix(r.URL.Path, "/assets") {
		handleStatusCode(w, http.StatusNotFound, nil)
		return
	} else {
		file_info, err := os.Stat(r.URL.Path[1:])
		if err != nil {
			handleStatusCode(w, http.StatusNotFound, nil)
			return
		} else if file_info.IsDir() {
			handleStatusCode(w, http.StatusForbidden, nil)
			return
		} else {
			http.ServeFile(w, r, r.URL.Path[1:])
		}
	}
}

// function to handle the download process
func (D *Data) HandleDownloads(Data *Data) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// The method should be GET
		fmt.Printf("d: inside downlaods %v\n", D)
		fmt.Printf("data: inside downlaods %v\n", Data)
		if !requestMethodChecker(w, r, http.MethodGet) {
			renderTemplate(w, "errorPage.html", http.StatusMethodNotAllowed, http.StatusMethodNotAllowed)
			return
		}

		if D.AsciiArt != "" && D.FormError == "" {
			w.Header().Add("Content-Type", "text/plain")
			w.Header().Add("Content-Disposition", "attachement")
			w.Header().Add("Content-Length", fmt.Sprint(len(D.AsciiArt)))
			file, err := os.Create("ascii-art.txt")
			if err != nil {
				renderTemplate(w, "errorPage.html", http.StatusInternalServerError, http.StatusInternalServerError)
				return
			}
			err = os.WriteFile(file.Name(), []byte(D.AsciiArt), 0o644)
			if err != nil {
				renderTemplate(w, "errorPage.html", http.StatusInternalServerError, http.StatusInternalServerError)
				return
			}
			http.ServeFile(w, r, file.Name())

		} else if D.FormError != "" {
			renderTemplate(w, "errorPage.html", http.StatusBadRequest, http.StatusBadRequest)
			return
		} else {
			renderTemplate(w, "errorPage.html", http.StatusBadRequest, http.StatusBadRequest)
			return
		}
	}
}
