package pagedata

import "html/template"

type PageData struct {
	Title   string
	Content template.HTML
}

func New(title string, content template.HTML) PageData {
	return PageData{
		Title:   title,
		Content: content,
	}
}
