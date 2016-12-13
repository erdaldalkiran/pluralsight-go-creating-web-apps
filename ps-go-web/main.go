package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"text/template"

	"math"

	"github.com/erdalkiran/pluralsight-go-creating-web-apps/constants"
)

func main() {
	ciko := math.Pi
	templates := populateTemplates()
	http.HandleFunc(constants.Paths.ROOT(), func(resp http.ResponseWriter, req *http.Request) {
		requestedFile := req.URL.Path[1:]
		template := templates.Lookup(requestedFile + ".html")
		if template == nil {
			resp.WriteHeader(404)
			return
		}

		template.Execute(resp, nil)
	})

	http.HandleFunc(constants.Paths.CSS(), serveResource)
	http.HandleFunc(constants.Paths.IMAGE(), serveResource)

	http.ListenAndServe(":8000", nil)
}

func serveResource(resp http.ResponseWriter, req *http.Request) {
	resourcePath := "public" + req.URL.Path
	file, err := os.Open(resourcePath)
	defer file.Close()

	if err != nil {
		resp.WriteHeader(404)
		return
	}

	resp.Header().Add("content-type", getContentType(req.URL.Path))
	reader := bufio.NewReader(file)
	reader.WriteTo(resp)

}

func getContentType(path string) string {
	var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}

	return contentType
}

func populateTemplates() *template.Template {

	basePath := "templates"
	templateFolder, err := os.Open(basePath)
	defer templateFolder.Close()

	if err != nil {
		panic(err)
	}

	templatePathsRaw, _ := templateFolder.Readdir(-1)
	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if pathInfo.IsDir() {
			continue
		}

		*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
	}
	result := template.New("mainTemplate")
	result.ParseFiles(*templatePaths...)

	return result
}

type ciko struct {
}

type hede struct {
	ciko ciko
}
