package main

import (
	"fmt"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")
var redirectUrl = os.Getenv("URL")

func handle(writer http.ResponseWriter, req *http.Request) {
	http.Redirect(writer, req, redirectUrl, 302)
}

func main() {

	http.HandleFunc("/", handle)

	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil); err != nil {
		panic(err)
	}
}
