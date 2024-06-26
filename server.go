//This program establishes a server and run it on my local machine
//Where the server actually loads an HTML file from my storage and
//shows it in the browser.

package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
)

func loadfile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func handler(w http.ResponseWriter, r *http.Request) {

	var html, err = loadfile(path.Base(r.URL.EscapedPath()))
	//This
	if err != nil {
		println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Didn't find your file")
		//This is to handle the case of not finding the intended page to load
	}
	fmt.Fprint(w, html)
	//fmt.Fprintf(w, "Hello %s, welcome", r.URL.Path[1:])
}

func main() {
	// http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello world"))
	// })

	http.HandleFunc("/", handler)

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":9000", nil)

}
