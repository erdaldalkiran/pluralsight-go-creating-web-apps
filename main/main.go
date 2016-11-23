package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("erdal was here!"))
	})

	http.ListenAndServe(":8000", nil)
}
