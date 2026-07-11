package services

import (
	"HtmxBlog/state"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

// Init initializes the template.
// It panics when some error occurs.
func Init() {
	state.AdminTmpl = InitAdminTemplate()

	tmpl, err := TemplateLoader("templates/user")
	if err != nil {
		panic("failed to load tmpl: " + err.Error())
	}
	state.Tmpl = tmpl
}

func InitAdminTemplate() *template.Template {
	return template.Must(TemplateLoader("templates/admin"))
}

func TemplateLoader(path string) (*template.Template, error) {
	var files []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(files...)
}
