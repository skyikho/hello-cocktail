package main

import (
	"fmt"
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

	fmt.Println(http.ListenAndServe(":3000", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlIndex)
}

func handleTime(w http.ResponseWriter, r *http.Request) {

	timeStr := "<h1> " + time.Now().String() + " </h1>"

	fmt.Fprintf(w, timeStr)
}

func handleVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Version 1.0 </h1>")
}
