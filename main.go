package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const htmlIndex = `
<html><body>
<h1> Hello Cocktail Cloud ! </h1>
</body><html>
`

func main() {

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/time", handleTime)
	http.HandleFunc("/version", handleVersion)

	log.Println("Hello Cocktail Server start")
	fmt.Println(http.ListenAndServe(":3000", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlIndex)
	log.Println("Index page was called")
}

func handleTime(w http.ResponseWriter, r *http.Request) {

	timeStr := "<h1> " + time.Now().String() + " </h1>"

	fmt.Fprintf(w, timeStr)
	log.Println("Time API was called")
}

func handleVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Version 1.7 </h1>")
	log.Println("Version API was called")
}
