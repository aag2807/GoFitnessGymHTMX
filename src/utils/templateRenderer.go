package utils

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type TemplateRenderer struct {
	templatesDir string
	cache        map[string]*template.Template
	mu           sync.RWMutex // To make the cache safe for concurrent access.
}

func TemplateToString(tmpl *template.Template, data interface{}) string {
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		log.Printf("Error executing template: %v\n", err)
		return ""
	}
	return buf.String()
}

func NewTemplateRenderer(templatesDir string) *TemplateRenderer {
	return &TemplateRenderer{
		templatesDir: templatesDir,
		cache:        make(map[string]*template.Template),
	}
}

// LoadTemplates loads all templates into the cache.
func (tr *TemplateRenderer) LoadTemplates() error {
	tr.mu.Lock()
	defer tr.mu.Unlock()

	layoutPath := filepath.Join(tr.templatesDir, "layout.html")
	layout, err := template.ParseFiles(layoutPath)
	if err != nil {
		return err
	}

	// Load page templates and combine with the layout
	pages, err := filepath.Glob(filepath.Join(tr.templatesDir, "*.html"))
	if err != nil {
		return err
	}
	for _, page := range pages {
		tmpl, err := template.Must(layout.Clone()).ParseFiles(page)
		if err != nil {
			return err
		}
		tr.cache[filepath.Base(page)] = tmpl
	}

	return nil
}

// RenderHTMLTemplate renders a template with the given name and data.
func (tr *TemplateRenderer) RenderHTMLTemplate(w http.ResponseWriter, tmplName string, data interface{}) error {
	tr.mu.RLock()
	tmpl, ok := tr.cache[tmplName]
	tr.mu.RUnlock()

	if !ok {
		log.Println("Template not found in cache:", tmplName)
		return tr.loadAndExecuteTemplate(w, tmplName, data)
	}

	return tmpl.Execute(w, data)
}

// loadAndExecuteTemplate loads and executes a template that isn't in the cache.
func (tr *TemplateRenderer) loadAndExecuteTemplate(w http.ResponseWriter, tmplName string, data interface{}) error {
	tr.mu.Lock()
	defer tr.mu.Unlock()

	tmplPath := filepath.Join(tr.templatesDir, tmplName)
	layoutPath := filepath.Join(tr.templatesDir, "layout.html")
	files := []string{layoutPath, tmplPath}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("Error parsing templates:", err)
		return err
	}

	return tmpl.Execute(w, data)
}

// loadAndExecuteTemplate loads and executes a template that isn't in the cache.
func (tr *TemplateRenderer) LoadAndExecuteTemplateWithDashboardLayout(w http.ResponseWriter, tmplName string, data interface{}) error {
	tr.mu.Lock()
	defer tr.mu.Unlock()

	tmplPath := filepath.Join(tr.templatesDir, tmplName)
	layoutPath := filepath.Join(tr.templatesDir, "dashboard-layout.html")
	files := []string{layoutPath, tmplPath}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("Error parsing templates:", err)
		return err
	}

	return tmpl.Execute(w, data)
}
