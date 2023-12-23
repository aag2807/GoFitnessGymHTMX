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
	return parseTemplatesFromRoutes(route)
}

func (pr *PartialRenderer) GetPageTemplatesToRender(templateName string) *template.Template {
	route := filepath.Join("src/templates/pages", templateName)
	return parseTemplatesFromRoutes(route)
}

func parseTemplatesFromRoutes(route string) *template.Template {
	tmpl, err := template.ParseFiles(route)
	if err != nil {
		panic(err)
	}

	return tmpl
}
