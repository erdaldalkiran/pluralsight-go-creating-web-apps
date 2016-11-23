package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.Handle("/", new(myHandler))

	http.ListenAndServe(":8000", nil)
}

type myHandler struct {
	http.Handler
}

func (this *myHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	file, err := os.Open(path)
	defer func() {
		file.Close()
	}()

	if err != nil {
		resp.WriteHeader(404)
		resp.Write([]byte("404 - " + http.StatusText(404) + " " + err.Error()))
		return
	}
	resp.Header().Add("Content-Type", getContentType(path))

	bufferedReader := bufio.NewReader(file)
	bufferedReader.WriteTo(resp)
}

func getContentType(file string) string {
	var contentType string
	extension := getExtension(file, ".")

	switch extension {
	case "css":
		contentType = "text/css"
	case "html":
		contentType = "text/html"
	case "js":
		contentType = "application/javascript"
	case "png":
		contentType = "image/png"
	case "mp4":
		contentType = "video/mp4"
	default:
		contentType = "text/plaing"

	}

	return contentType
}

func getExtension(s, seperator string) string {
	ss := strings.Split(s, seperator)
	return lastStringSlice(ss)
}

func lastStringSlice(sl []string) string {
	if len(sl) == 0 {
		return ""
	}

	return sl[len(sl)-1]
}
