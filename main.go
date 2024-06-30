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
		writer.Header().Add("Cache-Control", "public,  max-age=86400")
		writer.Header().Add("X-Content-Type-Options", "nosniff")
		writer.Header().Add("Content-Security-Policy", "frame-ancestors 'none'")
		writer.Header().Add("X-Frame-Options", "DENY")
		_, err = os.Stat(fmt.Sprintf("public/style/%s.css", templateName))
		data := pagedata.New(templateName, title, template.HTML(string(content)), !os.IsNotExist(err))
		tmplt.Execute(writer, data)
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleBookmarks(writer http.ResponseWriter, request *http.Request) {
	simpleGetHandler(writer, request, "bookmarks", "Edison Aguiar - Bookmarks")
}

func handleArticles(writer http.ResponseWriter, request *http.Request) {
	simpleGetHandler(writer, request, "articles", "Edison Aguiar - Articles")
}

func handleResume(writer http.ResponseWriter, request *http.Request) {
	simpleGetHandler(writer, request, "resume", "Edison Aguiar - Resume")
}

func handleIndex(writer http.ResponseWriter, request *http.Request) {
	// Handling cases where index is being used as a fallback for 404
	if request.URL.Path != "/" {
		writer.WriteHeader(http.StatusNotFound)
		simpleGetHandler(writer, request, "404", "Not found")
	} else {
		simpleGetHandler(writer, request, "index", "Edison Aguiar")
	}
}

func handleDatastoresDurability(writer http.ResponseWriter, request *http.Request) {
	simpleGetHandler(writer, request, "datastores-durability", "Durability Guarantees on different datastores")
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/articles", handleArticles)
	http.HandleFunc("/articles/datastores-durability", handleDatastoresDurability)
	http.HandleFunc("/bookmarks", handleBookmarks)
	http.HandleFunc("/resume", handleResume)

	http.ListenAndServe(":8888", nil)
}
