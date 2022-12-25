package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type PageData struct {
	Content template.HTML
}

func handleIndex(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		tmplt, err := template.ParseFiles("public/template/base.html")
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		content, err := ioutil.ReadFile("public/template/index.html")
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		data := PageData{
			Content: template.HTML(string(content)),
		}
		tmplt.Execute(writer, data)
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/style.css", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "style.css")
	})
	http.HandleFunc("/public/images/profile.jpg", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "public/images/profile.jpg")
	})

	http.ListenAndServe(":8888", nil)
}
