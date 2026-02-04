package template

import "html/template"

var Tmpl *template.Template

// Init initializes the template.
// It panics when some error occurs.
func Init() {
	Tmpl = template.Must(template.ParseGlob("templates/*.html"))
}
