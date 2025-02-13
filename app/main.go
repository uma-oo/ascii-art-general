package main

import (
	"fmt"
	"log"
	"net/http"

	"asciiWeb/handler"
)

func main() {
	var data handler.Data
	http.HandleFunc("/ascii-art", data.HandleAsciiArt())
	http.HandleFunc("/download", data.HandleDownloads(&data))
	http.HandleFunc("/assets/", handler.HandleAssets)
	http.HandleFunc("/", handler.HandleMainPage)
	fmt.Println("Server starting at http://localhost:6500")
	log.Fatal(http.ListenAndServe(":6500", nil))
}
