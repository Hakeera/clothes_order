package handlers

import (
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer(templateDir string) *TemplateRenderer {
	tmpl := template.New("")
	
	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(info.Name(), ".html") {
			// Usar o nome relativo como nome do template
			relPath, _ := filepath.Rel(templateDir, path)
			name := strings.ReplaceAll(relPath, string(filepath.Separator), "/")
			
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			
			_, err = tmpl.New(name).Parse(string(content))
			return err
		}
		return nil
	})
	
	if err != nil {
		log.Fatal(err)
	}
	
	return &TemplateRenderer{
		templates: tmpl,
	}
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

