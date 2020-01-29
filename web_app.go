package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hi", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		rw.Write([]byte(fmt.Sprintf("Hey, %s", name)))
	})
	http.ListenAndServe(":7070", nil)
}
