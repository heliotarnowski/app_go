package main

import (
	"fmt"
	"net/http"
	"log"
)

const webContent = "HT Helio APP : V. 1.02"

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
