package main

import (
	"fmt"
	"log"
	"net/http"

	"asciiWeb/handler"
)

func main() {
	http.HandleFunc("/assets/", handler.HandleAssets)
	http.HandleFunc("/", handler.HandleMainPage)
	http.HandleFunc("/ascii-art", handler.HandleAsciiArt)
	http.HandleFunc("/download", handler.HandleDownloads)
	fmt.Println("Server starting at http://localhost:6500")
	log.Fatal(http.ListenAndServe(":6500", nil))
}
