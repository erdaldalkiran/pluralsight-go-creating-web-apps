package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"text/template"

	"fmt"

	"log"

	"github.com/erdalkiran/pluralsight-go-creating-web-apps/constants"
	"github.com/erdalkiran/pluralsight-go-creating-web-apps/viewmodels"
)

const (
	TEMPLATES_FOLDER = "templates"
)

func main() {
	templates, err := populateTemplates()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	http.HandleFunc(constants.Paths.ROOT(), func(resp http.ResponseWriter, req *http.Request) {
		requestedFile := req.URL.Path[1:]

		template := templates.Lookup(requestedFile + ".html")
		if template == nil {
			fmt.Println("unable to find template")
			resp.WriteHeader(404)
			return
		}

		err := template.Execute(resp, getContext(requestedFile))
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc(constants.Paths.CSS(), serveResource)
	http.HandleFunc(constants.Paths.IMAGE(), serveResource)

	http.ListenAndServe(":8000", nil)
}

func getContext(requestedFile string) interface{} {
	var context interface{}
	if requestedFile == "home" {
		context = viewmodels.NewHome()
	} else if requestedFile == "categories" {
		context = viewmodels.NewCategories()
	} else if requestedFile == "products" {
		context = viewmodels.NewProducts()
	} else if requestedFile == "product" {
		context = viewmodels.NewProduct()
	}
	return context
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

func populateTemplates() (*template.Template, error) {
	templatePaths, err := getTemplatePaths(TEMPLATES_FOLDER)
	if err != nil {
		return nil, err
	}
	result := template.New("mainTemplate")
	_, err = result.ParseFiles(templatePaths...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getTemplatePaths(basePath string) ([]string, error) {
	templateFolder, err := os.Open(basePath)
	if err != nil {
		return nil, err
	}

	defer templateFolder.Close()

	templatePathsRaw, err := templateFolder.Readdir(-1)
	if err != nil {
		return nil, err
	}
	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if pathInfo.IsDir() {
			dirTemplatePaths, err := getTemplatePaths(basePath + "/" + pathInfo.Name())
			if err != nil {
				return nil, err
			}
			*templatePaths = append(*templatePaths, dirTemplatePaths...)
			continue
		}

		*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
	}

	return *templatePaths, nil
}
