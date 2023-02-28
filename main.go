package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/eaneto/eanetodev/pagedata"
	"github.com/eaneto/eanetodev/templatereader"
)

var templateReader = templatereader.New()

func simpleGetHandler(writer http.ResponseWriter, request *http.Request, templateName, title string) {
	if request.Method == "GET" {
		content, err := templateReader.Read("base")
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		tmplt, err := template.New("base").Parse(string(content))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		content, err = templateReader.Read(templateName)
		// sets response as cacheable for a week
		writer.Header().Add("Cache-Control", "public,  max-age=604800")
		_, err = os.Stat(fmt.Sprintf("public/style/%s.css", templateName))
		data := pagedata.New(templateName, title, template.HTML(string(content)), !os.IsNotExist(err))
		fmt.Println(data.HasStyle)
		tmplt.Execute(writer, data)
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleLinks(writer http.ResponseWriter, request *http.Request) {
	simpleGetHandler(writer, request, "links", "Edison Aguiar - Links")
}

func handleArticles(writer http.ResponseWriter, request *http.Request) {
	simpleGetHandler(writer, request, "articles", "Edison Aguiar - Articles")
}

func handleResume(writer http.ResponseWriter, request *http.Request) {
	simpleGetHandler(writer, request, "resume", "Edison Aguiar - Resume")
}

func handleIndex(writer http.ResponseWriter, request *http.Request) {
	simpleGetHandler(writer, request, "index", "Edison Aguiar")
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/articles", handleArticles)
	http.HandleFunc("/links", handleLinks)
	http.HandleFunc("/resume", handleResume)
	http.HandleFunc("/style.css", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "style.css")
	})
	http.HandleFunc("/public/style/resume.css", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "public/style/resume.css")
	})
	http.HandleFunc("/public/images/profile.jpg", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "public/images/profile.jpg")
	})

	http.ListenAndServe(":8888", nil)
}
