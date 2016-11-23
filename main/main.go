package main

import "net/http"
import "text/template"
import "fmt"

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Content-Type", "text/html")

		temp := template.New("template")
		temp.New("main").Parse(mainTemplate)
		temp.New("header").Parse(header)
		temp.New("body").Parse(body)
		temp.New("footer").Parse(footer)

		context := Context{
			Title:  "Title",
			IsTrue: true,
			Fruits: [3]string{"Lemon", "Orange", "Apple"},
		}

		err := temp.Lookup("main").Execute(resp, context)
		if err != nil {
			fmt.Println("Error")
			fmt.Println(err.Error())
			resp.Write([]byte(err.Error()))
		}
	})
	http.ListenAndServe(":8000", nil)
}

type Context struct {
	Title  string
	IsTrue bool
	Fruits [3]string
}

const mainTemplate = `
	{{template "header" .Title}}
	{{template "body" .}}
	{{template "footer"}}
`

const body = `
	<body>
		<h1>List of Fruits</h1>
		{{if eq .IsTrue}}
			<p>It's true</p>
		{{else}}
			<p>It's false</p>
		{{end}}
		<ul>
			{{range .Fruits}}
				<li>{{.}}</li>
			{{end}}
		</ul>
	</body>
`

const header = `
	<!DOCTYPE html>
	<html>
		<head>
			<title>
				{{.}}
			</title>	
		</head>
`

const footer = `
	</html>
`
