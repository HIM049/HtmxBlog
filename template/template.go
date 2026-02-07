package template

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var Tmpl *template.Template

// Init initializes the template.
// It panics when some error occurs.
func Init() {
	var files []string

	err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	Tmpl = template.Must(template.ParseFiles(files...))
}
