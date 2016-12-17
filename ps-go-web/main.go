package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/erdalkiran/pluralsight-go-creating-web-apps/constants"
	"github.com/erdalkiran/pluralsight-go-creating-web-apps/viewmodels"
)

const (
	TEMPLATES_FOLDER = "templates"
)

func main() {
	templates := populateTemplates()
	http.HandleFunc(constants.Paths.ROOT(), func(resp http.ResponseWriter, req *http.Request) {
		requestedFile := req.URL.Path[1:]

		template := templates.Lookup(requestedFile + ".html")
		if template == nil {
			resp.WriteHeader(404)
			return
		}

		var context interface{}
		if requestedFile == "home" {
			context = viewmodels.NewHome()
		}

		template.Execute(resp, context)
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
	templatePaths := getTemplatePaths(TEMPLATES_FOLDER)
	result := template.New("mainTemplate")
	result.ParseFiles(templatePaths...)

	return result
}

func getTemplatePaths(basePath string) []string {
	templateFolder, err := os.Open(basePath)
	defer templateFolder.Close()

	if err != nil {
		panic(err)
	}

	templatePathsRaw, _ := templateFolder.Readdir(-1)
	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if pathInfo.IsDir() {
			dirTemplatePaths := getTemplatePaths(basePath + "/" + pathInfo.Name())
			*templatePaths = append(*templatePaths, dirTemplatePaths...)
			continue
		}

		*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
	}

	return *templatePaths
}
