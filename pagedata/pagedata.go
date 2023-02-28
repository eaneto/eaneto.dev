package pagedata

import "html/template"

type PageData struct {
	Name     string
	Title    string
	Content  template.HTML
	HasStyle bool
}

func New(name, title string, content template.HTML, hasStyle bool) PageData {
	return PageData{
		Name:     name,
		Title:    title,
		Content:  content,
		HasStyle: hasStyle,
	}
}
