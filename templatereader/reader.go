package templatereader

import (
	"fmt"
	"io/ioutil"
	"sync"
)

type TemplateReader struct {
	Mutex     sync.Mutex
	Templates map[string][]byte
}

func New() TemplateReader {
	return TemplateReader{
		Mutex:     sync.Mutex{},
		Templates: make(map[string][]byte),
	}
}

func (reader *TemplateReader) Read(templateName string) ([]byte, error) {
	content, ok := reader.Templates[templateName]
	if !ok {
		content, err := ioutil.ReadFile(fmt.Sprintf("public/template/%s.html", templateName))
		if err != nil {
			return []byte{}, err
		}
		reader.Mutex.Lock()
		defer reader.Mutex.Unlock()
		reader.Templates[templateName] = content
		return content, nil
	}
	return content, nil
}
