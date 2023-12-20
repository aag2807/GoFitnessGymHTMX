package utils

import (
	"html/template"
	"path/filepath"
)

type PartialRenderer struct{}

func NewPartialRenderer() *PartialRenderer {
	return &PartialRenderer{}
}

func (pr *PartialRenderer) GetTemplatePartialToRender(templateName string) *template.Template {
	route := filepath.Join("src/templates/partials", templateName)
	tmpl, err := template.ParseFiles(route)
	if err != nil {
		panic(err)
	}

	return tmpl
}
