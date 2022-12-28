package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/eaneto/eanetodev/pagedata"
)

var templates = make(map[string][]byte)

func simpleGetHandler(writer http.ResponseWriter, request *http.Request, templateName string, title string) {
	if request.Method == "GET" {
		content, err := readTemplateFromCacheOrFile("base")
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		tmplt, err := template.New("base").Parse(string(content))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		content, err = readTemplateFromCacheOrFile(templateName)
		// sets response as cacheable for a week
		writer.Header().Add("Cache-Control", "public,  max-age=604800")
		data := pagedata.New(title, template.HTML(string(content)))
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

func handleIndex(writer http.ResponseWriter, request *http.Request) {
	simpleGetHandler(writer, request, "index", "Edison Aguiar")
}

func readTemplateFromCacheOrFile(templateName string) ([]byte, error) {
	content, ok := templates[templateName]
	if !ok {
		content, err := ioutil.ReadFile(fmt.Sprintf("public/template/%s.html", templateName))
		if err != nil {
			return []byte{}, err
		}
		templates[templateName] = content
		return content, nil
	}
	return content, nil
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/articles", handleArticles)
	http.HandleFunc("/links", handleLinks)
	http.HandleFunc("/style.css", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "style.css")
	})
	http.HandleFunc("/public/images/profile.jpg", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "public/images/profile.jpg")
	})

	http.ListenAndServe(":8888", nil)
}
