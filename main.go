package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello go!")
	//fmt.Fprintf(w, "Hello go! %s\n", r.URL.Path)
    })

    http.ListenAndServe(":8080", nil)
}
