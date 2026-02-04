package template

import "html/template"

var Tmpl *template.Template

func Init() {
	Tmpl = template.Must(template.ParseGlob("templates/*.gtpl"))
}
